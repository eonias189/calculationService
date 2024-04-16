package use_connect

import (
	"context"
	"strconv"

	"github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/metadata"
)

type Handler struct {
	e *Executor
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

	err = h.e.OnStart(id)
	if err != nil {
		return err
	}

	tasks, err := h.e.GetTasks(id)
	if err != nil {
		return err
	}

	defer h.e.OnConnClose(id)

	g, ctx := errgroup.WithContext(context.Background())
	results := make(chan *pb.ResultResp)
	defer close(results)

	// send tasks
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case task, ok := <-tasks:
				if !ok {
					return errors.ErrChanClosed
				}
				err := h.e.OnTask(id, task)
				if err != nil {
					return err
				}

				err = conn.Send(task)
				if err != nil {
					return err
				}
			}
		}
	})

	// receive responses
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case resp := <-utils.Await(func() struct {
				Resp *pb.ResultResp
				Err  error
			} {
				resp, err := conn.Recv()
				return struct {
					Resp *pb.ResultResp
					Err  error
				}{Resp: resp, Err: err}
			}):
				if resp.Err != nil {
					return resp.Err
				}

				err = h.e.OnResult(id, resp.Resp)
				if err != nil {
					return err
				}
			}
		}
	})

	err = g.Wait()
	if err != nil {
		logger.Error(err)
	}
	return err

}

func MakeHandler(tasksService TaskService, agentsServise AgentService, distributor Distributor) Connector {
	return &Handler{
		e: NewExecutor(tasksService, agentsServise, distributor),
	}
}
