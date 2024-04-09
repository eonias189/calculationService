package orchestrator

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"time"

	pb "github.com/eonias189/calculationService/orchestrator/internal/proto"
	"github.com/eonias189/calculationService/orchestrator/internal/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

type Service[T any] interface {
	Add(ctx context.Context, item T) (int64, error)
	GetAll(ctx context.Context, limit, offset int) ([]T, error)
	GetById(ctx context.Context, id int64) (T, error)
	Update(ctx context.Context, item T) error
	Delete(ctx context.Context, id int64) error
}

type Orchestrator struct {
	logger        *slog.Logger
	taskService   Service[service.Task]
	agentsService Service[service.Agent]
	pb.UnimplementedOrchestratorServer
}

func (o *Orchestrator) Connect(conn pb.Orchestrator_ConnectServer) error {
	metadata, ok := metadata.FromIncomingContext(conn.Context())

	if !ok {
		return errors.Errorf("metadata is invalid or not provided")
	}

	maxThreadsSlice := metadata.Get("maxThreads")
	if len(maxThreadsSlice) == 0 {
		return errors.Errorf("metadata is invalid or not provided")
	}

	maxThreads, err := strconv.Atoi(maxThreadsSlice[0])
	if err != nil {
		return errors.Errorf("metadata is invalid or not provided")
	}

	fmt.Println(maxThreads)
	cancel := make(chan struct{})

	go func() {
		for {
			msg, err := conn.Recv()
			if err != nil {
				fmt.Println(err)
				cancel <- struct{}{}
				return
			}
			fmt.Println(msg)
		}
	}()

	for {
		select {
		case <-cancel:
			return io.EOF
		case <-time.After(time.Second * 10):
			err := conn.Send(&pb.Task{Id: 1, Expression: "2 + 2 * 2", Timeouts: &pb.Timeouts{Add: 3}})
			if err != nil {
				o.logger.With(slog.String("while", "sending message")).Error(err.Error())
				break
			}
		}

	}
}

func (o *Orchestrator) Tasks(limit, offset int) ([]service.Task, error) {
	o.logger.Info("sending tasks with limit " + fmt.Sprint(limit) + " and offset " + fmt.Sprint(offset))
	return o.taskService.GetAll(context.TODO(), limit, offset)
}

func (o *Orchestrator) Task(id int64) (service.Task, error) {
	o.logger.Info("sendint task " + fmt.Sprint(id))
	return o.taskService.GetById(context.TODO(), id)
}

func (o *Orchestrator) AddTask(expression string) (int64, error) {
	o.logger.Info("adding task " + expression)
	return o.taskService.Add(context.TODO(), service.Task{Expression: expression, Status: service.TaskStatusPending})
}

func (o *Orchestrator) Agents(limit, offset int) ([]service.Agent, error) {
	o.logger.Info("sending agents with limit " + fmt.Sprint(limit) + " and offset " + fmt.Sprint(offset))
	return o.agentsService.GetAll(context.TODO(), limit, offset)
}

func (o *Orchestrator) Timeouts() (service.Timeouts, error) {
	return service.Timeouts{Add: 1, Sub: 2, Mul: 3, Div: 4}, nil
}

func (o *Orchestrator) SetTimeouts(service.Timeouts) error {
	return nil
}

func (o *Orchestrator) Start(ctx context.Context) error {
	o.logger.Info("starting")
	return nil
}

func (o *Orchestrator) Close() {
	o.logger.Info("closing orchestrator")
}

type OrchestratorConfig struct {
	Logger       *slog.Logger
	TaskService  Service[service.Task]
	AgentService Service[service.Agent]
}

func New(opts *OrchestratorConfig) *Orchestrator {
	if opts == nil {
		opts = &OrchestratorConfig{Logger: slog.Default()}
	}
	return &Orchestrator{logger: opts.Logger, taskService: opts.TaskService, agentsService: opts.AgentService}
}
