package agent

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/eonias189/calculationService/backend/internal/config/agent_config"
	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/lib/pool"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func Calculate(task *pb.Task) (float64, error) {
	addCount := strings.Count(task.Expression, "+")
	subCount := strings.Count(task.Expression, "-")
	mulCount := strings.Count(task.Expression, "*")
	divCount := strings.Count(task.Expression, "/")

	evalExpr, err := govaluate.NewEvaluableExpression(task.Expression)
	if err != nil {
		return 0, errs.ErrExecution
	}

	res, err := evalExpr.Evaluate(map[string]interface{}{})
	if err != nil {
		return 0, errs.ErrExecution
	}

	resFloat64, ok := res.(float64)
	if !ok {
		return 0, errs.ErrExecution
	}

	time.Sleep(time.Second * time.Duration(addCount) * time.Duration(task.Timeouts.Add))
	time.Sleep(time.Second * time.Duration(subCount) * time.Duration(task.Timeouts.Sub))
	time.Sleep(time.Second * time.Duration(mulCount) * time.Duration(task.Timeouts.Mul))
	time.Sleep(time.Second * time.Duration(divCount) * time.Duration(task.Timeouts.Div))

	return resFloat64, nil

}

type Application struct {
	wp         *pool.WorkerPool
	wg         *sync.WaitGroup
	maxThreads int
	address    string
}

func SetConnMetadata(ctx context.Context, id int64) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "id", fmt.Sprint(id))
}

func (a *Application) GetTasks(ctx context.Context, conn pb.Orchestrator_ConnectClient, out chan<- *pb.Task) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case received := <-utils.Await(func() struct {
				task *pb.Task
				err  error
			} {
				task, err := conn.Recv()
				return struct {
					task *pb.Task
					err  error
				}{task: task, err: err}
			}):
				if received.err != nil {
					logger.Error(received.err)
					return received.err
				}
				logger.Info("got task", received.task.String())

				select {
				case <-ctx.Done():
					return ctx.Err()
				case out <- received.task:
				}
			}
		}
	}
}

func (a *Application) SolveTasks(ctx context.Context, tasks <-chan *pb.Task, out chan<- *pb.ResultResp) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case task, ok := <-tasks:
				if !ok {
					return
				}

				a.wp.AddTask(pool.NewTask(func() {
					res, err := Calculate(task)
					resp := &pb.ResultResp{
						TaskId: task.Id,
						Result: res,
						Error:  err != nil,
					}
					logger.Info("solved task:", resp.String())
					out <- resp
				}))
			}
		}
	}()
}

func (a *Application) SendResults(ctx context.Context, conn pb.Orchestrator_ConnectClient, resps <-chan *pb.ResultResp) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case resp, ok := <-resps:
				if !ok {
					return errs.ErrChanClosed
				}

				resp.SendTime = time.Now().UnixNano()
				resp.RunningThreads = int64(a.wp.ExecutingTasks())
				err := conn.Send(resp)
				if err != nil {
					logger.Error(err)
					return err
				}
				logger.Info("sent result: ", resp.String())
			}
		}
	}

}

func (a *Application) Register(ctx context.Context, cli pb.OrchestratorClient) int64 {
	const (
		interval = time.Second * 5
	)
	req := &pb.RegisterReq{MaxThreads: int64(a.maxThreads)}
	var resp *pb.RegisterResp

	utils.TryUntilSuccess(ctx, func() error {
		r, err := cli.Register(context.TODO(), req)
		if err == nil {
			resp = r
		} else {
			logger.Error(err)
		}
		return err
	}, interval)

	return resp.GetId()
}

func (a *Application) GetCli(ctx context.Context) pb.OrchestratorClient {
	const (
		interval = time.Second * 5
	)

	var conn *grpc.ClientConn
	utils.TryUntilSuccess(ctx, func() error {
		cliConn, err := grpc.Dial(a.address, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err == nil {
			conn = cliConn
		} else {
			logger.Error(err)
		}

		return err
	}, interval)

	return pb.NewOrchestratorClient(conn)
}

func (a *Application) Run(ctx context.Context) {
	const (
		interval = time.Second * 10
	)

	tasks := make(chan *pb.Task)
	resps := make(chan *pb.ResultResp)

	a.wp.Start(ctx)
	a.SolveTasks(ctx, tasks, resps)

	cli := a.GetCli(ctx)
	logger.Info("got client")

	id := a.Register(ctx, cli)
	metadataCtx := SetConnMetadata(context.Background(), id)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := cli.Connect(metadataCtx)
			if err != nil {
				logger.Error(err)
				time.Sleep(interval)
				continue
			}

			logger.Info("got connection")

			g, errGrCtx := errgroup.WithContext(ctx)
			g.Go(a.GetTasks(errGrCtx, conn, tasks))
			g.Go(a.SendResults(errGrCtx, conn, resps))

			err = g.Wait()
			if err != nil {
				logger.Error(err)
				time.Sleep(interval)
				continue
			}
		}

	}

}

func (a *Application) Close() {
	// a.wg.Wait()
	// a.wp.Close()
}

func New(cfg agent_config.Config) *Application {
	return &Application{
		wp:         pool.NewWorkerPool(cfg.MaxThreads),
		wg:         &sync.WaitGroup{},
		maxThreads: cfg.MaxThreads,
		address:    cfg.OrchestratorAddr,
	}
}
