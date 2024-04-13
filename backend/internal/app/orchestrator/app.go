package orchestrator

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/eonias189/calculationService/backend/internal/config/orchestrator_config"
	"github.com/eonias189/calculationService/backend/internal/lib/distributor"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
	use_connect "github.com/eonias189/calculationService/backend/internal/use_cases/connect"
	use_register "github.com/eonias189/calculationService/backend/internal/use_cases/register"
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
	distributor     *distributor.Distributor[*pb.Task]
	cfg             orchestrator_config.Config
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

		api.Get("/distribute", func(w http.ResponseWriter, r *http.Request) {
			t := service.Task{Expression: "2 + 2 * 2", UserId: 1, Status: service.TaskStatusPending}
			id, err := a.taskService.Add(t)
			if err != nil {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, render.M{"reason": err.Error()})
				return
			}
			err = a.distributor.Distribute(&pb.Task{Id: id, Expression: t.Expression, Timeouts: &pb.Timeouts{Add: 10, Sub: 2, Mul: 15, Div: 4}})

			if err != nil {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, render.M{"reason": err.Error()})
				return
			}
			render.JSON(w, r, render.M{"message": "ok"})
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
	registerer use_register.Registerer
	connector  use_connect.Connector
	pb.UnimplementedOrchestratorServer
}

func (gh *GrpcHandler) Connect(conn pb.Orchestrator_ConnectServer) error {
	return gh.connector.Connect(conn)
}

func (gh *GrpcHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return gh.registerer.Register(ctx, req)
}

func (a *Application) StartGrpcServer(ctx context.Context) func() error {
	return func() error {
		listener, err := net.Listen("tcp", a.cfg.GRPCAddress)
		if err != nil {
			return err
		}

		serv := grpc.NewServer()
		handler := &GrpcHandler{
			registerer: use_register.MakeHandler(a.agentService),
			connector:  use_connect.MakeHandler(a.taskService, a.agentService, a.timeoutsService, a.distributor),
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

	err := a.agentService.DisactivateAll()
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(time.Second * 15)
		logger.Info("cleaning and pulling")
		a.taskService.SetPendingForDisactiveAgents()
		tasks, err := a.taskService.GetAllPending()
		if err == nil {
			for _, task := range tasks {
				a.distributor.Distribute(&pb.Task{Id: task.Task.Id, Expression: task.Task.Expression,
					Timeouts: &pb.Timeouts{Add: task.Timeouts.Add, Sub: task.Timeouts.Sub,
						Mul: task.Timeouts.Mul, Div: task.Timeouts.Div}})
			}
		}
	}()

	a.distributor.StartPushing(time.Second * 5)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(a.StartHttpServer(ctx))
	g.Go(a.StartGrpcServer(ctx))

	return g.Wait()
}

func (a *Application) Close() {
	a.distributor.Close()
}

func (a *Application) init() error {
	pool, err := pgxpool.New(context.TODO(), a.cfg.PostgresConn)
	if err != nil {
		return err
	}

	utils.TryUntilSuccess(context.TODO(), func() error { return pool.Ping(context.TODO()) }, time.Second*5)

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
	a.distributor = distributor.NewDistributor[*pb.Task](5)
	return nil
}

func New(cfg orchestrator_config.Config) (*Application, error) {
	a := &Application{cfg: cfg}
	err := a.init()

	if err != nil {
		return nil, err
	}

	return a, nil
}
