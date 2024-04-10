package application

import (
	"context"
	"net"
	"net/http"

	"github.com/eonias189/calculationService/orchestrator/internal/config"
	"github.com/eonias189/calculationService/orchestrator/internal/lib/utils"
	"github.com/eonias189/calculationService/orchestrator/internal/logger"
	pb "github.com/eonias189/calculationService/orchestrator/internal/proto"
	"github.com/eonias189/calculationService/orchestrator/internal/service"
	use_connect "github.com/eonias189/calculationService/orchestrator/internal/use_cases/connect"
	use_register "github.com/eonias189/calculationService/orchestrator/internal/use_cases/register"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Application struct {
	taskService     *service.TaskService
	agentService    *service.AgentService
	timeoutsService *service.TimeoutsSerice
	cfg             config.Config
}

func (a *Application) StartHttpServer(ctx context.Context) func() error {
	return func() error {
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		api := chi.NewRouter()
		api.Use(render.SetContentType(render.ContentTypeJSON))

		api.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, render.M{"message": "pong"})
		})

		r.Mount("/api", api)

		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-utils.Await(func() error {
			return http.ListenAndServe(a.cfg.HttpAddress, r)
		}):
			return err
		}
	}

}

type GrpcHandler struct {
	use_register.Registerer
	use_connect.Connector
	pb.UnimplementedOrchestratorServer
}

func (gh *GrpcHandler) Connect(conn pb.Orchestrator_ConnectServer) error {
	return gh.Connector.Connect(conn)
}

func (gh *GrpcHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return gh.Registerer.Register(ctx, req)
}

func (a *Application) StartGrpcServer(ctx context.Context) func() error {
	return func() error {
		listener, err := net.Listen("tcp", a.cfg.GRPCAddress)
		if err != nil {
			return err
		}

		serv := grpc.NewServer()
		handler := &GrpcHandler{
			Registerer: use_register.MakeHandler(a.agentService),
			Connector:  use_connect.MakeHandler(a.taskService, a.agentService, a.timeoutsService),
		}
		pb.RegisterOrchestratorServer(serv, handler)

		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-utils.Await(func() error {
			return serv.Serve(listener)
		}):
			return err
		}
	}
}

func (a *Application) Run(ctx context.Context) error {
	logger.Info("starting")
	g, ctx := errgroup.WithContext(ctx)

	g.Go(a.StartHttpServer(ctx))
	g.Go(a.StartGrpcServer(ctx))

	return g.Wait()
}

func (a *Application) Close() {

}

func (a *Application) init() error {
	pool, err := pgxpool.New(context.TODO(), a.cfg.PostgresConn)
	if err != nil {
		return err
	}

	taskService, err := service.NewTaskService(pool)
	if err != nil {
		return err
	}

	agentService, err := service.NewAgentService(pool)
	if err != nil {
		return err
	}

	timeoutsService, err := service.NewTimeoutsService(pool)
	if err != nil {
		return err
	}

	a.taskService = taskService
	a.agentService = agentService
	a.timeoutsService = timeoutsService
	return nil
}

func New(cfg config.Config) (*Application, error) {
	a := &Application{cfg: cfg}
	err := a.init()

	if err != nil {
		return nil, err
	}

	return a, nil
}
