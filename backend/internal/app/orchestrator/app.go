package orchestrator

import (
	"context"
	"net"
	"time"

	"github.com/eonias189/calculationService/backend/internal/config/orchestrator_config"
	"github.com/eonias189/calculationService/backend/internal/lib/distributor"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
	use_connect "github.com/eonias189/calculationService/backend/internal/use_cases/connect"
	use_distribute "github.com/eonias189/calculationService/backend/internal/use_cases/distribute"
	use_register "github.com/eonias189/calculationService/backend/internal/use_cases/register"
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

type GrpcHandler struct {
	registerer    use_register.Registerer
	connector     use_connect.Connector
	distributable use_distribute.Distributable
	pb.UnimplementedOrchestratorServer
}

func (gh *GrpcHandler) Connect(conn pb.Orchestrator_ConnectServer) error {
	return gh.connector.Connect(conn)
}

func (gh *GrpcHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return gh.registerer.Register(ctx, req)
}

func (gh *GrpcHandler) Distribute(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	return gh.distributable.Distribute(ctx, task)
}

func (a *Application) StartGrpcServer(ctx context.Context) func() error {
	return func() error {
		listener, err := net.Listen("tcp", a.cfg.Address)
		if err != nil {
			return err
		}

		serv := grpc.NewServer()
		handler := &GrpcHandler{
			registerer:    use_register.MakeHandler(a.agentService),
			connector:     use_connect.MakeHandler(a.taskService, a.agentService, a.distributor),
			distributable: use_distribute.MakeHandler(a.distributor),
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
	logger.Info("starting at", a.cfg.Address)

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
					Timeouts: &pb.Timeouts{Add: uint64(task.Timeouts.Add), Sub: uint64(task.Timeouts.Sub),
						Mul: uint64(task.Timeouts.Mul), Div: uint64(task.Timeouts.Div)}})
			}
		}
	}()

	a.distributor.StartPushing(ctx, time.Second*5)
	g, ctx := errgroup.WithContext(ctx)
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
