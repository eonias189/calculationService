package use_connect

import (
	"context"
	"strconv"

	"github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/metadata"
)

type Handler struct {
	e *Executor
}

type ReceiveResultResp struct {
	Result *pb.ResultResp
	Err    error
}

func ReceiveResults(ctx context.Context, conn pb.Orchestrator_ConnectServer, out chan<- *pb.ResultResp) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case resp := <-utils.Await(func() ReceiveResultResp {
				res, err := conn.Recv()
				return ReceiveResultResp{Result: res, Err: err}
			}):
				if resp.Err != nil {
					return resp.Err
				}

				select {
				case <-ctx.Done():
					return ctx.Err()

				case out <- resp.Result:
				}
			}
		}
	}
}

func SendTasks(ctx context.Context, conn pb.Orchestrator_ConnectServer, in <-chan *pb.Task) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case task, ok := <-in:
				if !ok {
					return errors.ErrChanClosed
				}

				err := conn.Send(task)
				if err != nil {
					return err
				}
			}
		}
	}
}

func ReadConnectMetadata(ctx context.Context) (int64, error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, errors.ErrMetadataInvalidOrNotProvided
	}

	ids := metadata.Get("id")
	if len(ids) == 0 {
		return 0, errors.ErrMetadataInvalidOrNotProvided
	}

	id, e := strconv.Atoi(ids[0])
	if e != nil {
		return 0, errors.ErrMetadataInvalidOrNotProvided
	}

	return int64(id), nil
}

func (h *Handler) Connect(conn pb.Orchestrator_ConnectServer) error {
	id, err := ReadConnectMetadata(conn.Context())
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(context.Background())

	tasks := make(chan *pb.Task)
	results := make(chan *pb.ResultResp)

	defer close(tasks)
	defer close(results)

	g.Go(SendTasks(ctx, conn, tasks))
	g.Go(ReceiveResults(ctx, conn, results))
	g.Go(h.e.Do(ctx, id, tasks, results))

	return g.Wait()

}

func MakeHandler(tasksService TaskService, agentsServise AgentService, timeoutsService TimeoutsService) Connector {
	return &Handler{
		e: NewExecutor(tasksService, agentsServise, timeoutsService),
	}
}
