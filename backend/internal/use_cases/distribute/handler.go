package use_distribute

import (
	"context"

	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Handler struct {
	e *Executor
}

func (h *Handler) Distribute(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	return &pb.Empty{}, h.e.Do(task)
}

func MakeHandler(d Distributor) Distributable {
	return &Handler{e: NewExecutor(d)}
}
