package agent

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/eonias189/calculationService/agent/internal/lib/pool"
	pb "github.com/eonias189/calculationService/agent/internal/proto"
)

type Agent struct {
	api        pb.OrchestratorClient
	maxThreads int
	logger     *slog.Logger
	wp         pool.WorkerPool
	wg         sync.WaitGroup
}

var (
	ErrExec           = fmt.Errorf("ExecutionError")
	MaxSendRetries    = 5
	SendRetryInterval = time.Second * 5
)

func (a *Agent) Calculate(task *pb.Task) (float64, error) {
	addCount := strings.Count(task.Expression, "+")
	subCount := strings.Count(task.Expression, "-")
	mulCount := strings.Count(task.Expression, "*")
	divCount := strings.Count(task.Expression, "/")

	evalExpr, err := govaluate.NewEvaluableExpression(task.Expression)
	if err != nil {
		return 0, ErrExec
	}

	res, err := evalExpr.Evaluate(nil)
	if err != nil {
		return 0, ErrExec
	}

	resFloat64, ok := res.(float64)
	if !ok {
		return 0, ErrExec
	}

	time.Sleep(time.Second * time.Duration(addCount) * time.Duration(task.Timeouts.Add))
	time.Sleep(time.Second * time.Duration(subCount) * time.Duration(task.Timeouts.Sub))
	time.Sleep(time.Second * time.Duration(mulCount) * time.Duration(task.Timeouts.Mul))
	time.Sleep(time.Second * time.Duration(divCount) * time.Duration(task.Timeouts.Div))

	return resFloat64, nil

}

type ReceiveResp struct {
	Task *pb.Task
	Err  error
}

func Receive(conn pb.Orchestrator_ConnectClient) <-chan ReceiveResp {
	out := make(chan ReceiveResp)

	go func() {
		defer close(out)
		resp, err := conn.Recv()
		if err != nil {
			out <- ReceiveResp{Err: err}
			return
		}
		out <- ReceiveResp{Task: resp}
	}()

	return out
}

func (a *Agent) PutMessages(ctx context.Context, out chan *pb.Task) <-chan error {
	errCh := make(chan error)

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		defer close(errCh)

		conn, err := a.api.Connect(context.TODO(), &pb.ConnReq{MaxThreads: int64(a.maxThreads)})
		if err != nil {
			errCh <- err
			return
		}
		a.logger.With(slog.String("while", "connecting")).Info("connected")

		for {
			select {
			case <-ctx.Done():
				return

			case received := <-Receive(conn):
				if received.Err != nil {
					errCh <- received.Err
					return
				}
				a.logger.With(slog.String("while", "receiving")).Info("received" + received.Task.String())
				out <- received.Task
			}
		}
	}()

	return errCh
}

func (a *Agent) ReceiveMessages(ctx context.Context) <-chan *pb.Task {
	out := make(chan *pb.Task)

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return

			case err := <-a.PutMessages(ctx, out):
				a.logger.With(slog.String("while", "receiving")).Error(err.Error())
				time.Sleep(time.Second * 10)
			}
		}
	}()

	return out
}

func (a *Agent) CompleteTasks(ctx context.Context, tasks <-chan *pb.Task) <-chan *pb.ResultResp {
	out := make(chan *pb.ResultResp)
	a.wg.Add(1)

	go func() {
		defer a.wg.Done()
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return

			case task, ok := <-tasks:
				if !ok {
					return
				}

				complete := pool.NewTask(func() {
					res, err := a.Calculate(task)
					resp := &pb.ResultResp{TaskId: task.Id, Result: res}
					if err != nil {
						resp.Error = true
					}
					out <- resp
				})

				a.wp.AddTask(complete)
			}
		}
	}()

	return out
}

func (a *Agent) SendResult(resp *pb.ResultResp) error {
	var err error
	var retry int
	_, err = a.api.SetResult(context.TODO(), resp)

	for err != nil {
		if retry == MaxSendRetries {
			break
		}

		time.Sleep(SendRetryInterval)
		_, err = a.api.SetResult(context.TODO(), resp)
		retry++
	}

	return err
}

func (a *Agent) SendResults(ctx context.Context, resps <-chan *pb.ResultResp) {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		for {
			select {
			case <-ctx.Done():
				return

			case resp, ok := <-resps:
				if !ok {
					return
				}

				err := a.SendResult(resp)
				if err != nil {
					a.logger.With(slog.String("while", "sending result")).Error(err.Error())
				}
			}
		}
	}()
}

func (a *Agent) Run(ctx context.Context) error {
	tasks := a.ReceiveMessages(ctx)
	a.wp.Start(ctx)
	resps := a.CompleteTasks(ctx, tasks)
	a.SendResults(ctx, resps)
	return nil
}

func (a *Agent) Close() {
	a.wg.Wait()
	a.wp.Close()
}

type AgentOptions struct {
	Api        pb.OrchestratorClient
	WorkerPool pool.WorkerPool
	Logger     *slog.Logger
	MaxThreads int
}

func New(opts *AgentOptions) *Agent {
	if opts == nil {
		opts = &AgentOptions{Logger: slog.Default()}
	}

	return &Agent{api: opts.Api,
		logger:     opts.Logger,
		maxThreads: opts.MaxThreads,
		wp:         opts.WorkerPool,
	}
}
