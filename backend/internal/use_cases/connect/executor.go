package use_connect

import (
	"context"
	"fmt"
	"time"

	"github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Executor struct {
	as  AgentService
	ts  TaskService
	tms TimeoutsService
}

func (e *Executor) Do(ctx context.Context, id int64, tasksOut chan<- *pb.Task, respsIn <-chan *pb.ResultResp) func() error {
	logger.Info("handing connect with", id)
	return func() error {
		for {
			select {
			case <-ctx.Done():
				err := e.as.Delete(id)

				if err != nil {
					return err
				}

				return ctx.Err()

			case reult, ok := <-respsIn:
				if !ok {
					return errors.ErrChanClosed
				}

				logger.Info("got result", reult)

			case <-time.After(time.Second * 10):
				select {
				case <-ctx.Done():
					fmt.Println(ctx.Err())
					return ctx.Err()
				case tasksOut <- &pb.Task{Id: 1, Expression: "2 + 2 * 2", Timeouts: &pb.Timeouts{Add: 3}}:
				}
			}
		}
	}
}

func NewExecutor(ts TaskService, as AgentService, tms TimeoutsService) *Executor {
	return &Executor{as: as, ts: ts, tms: tms}
}
