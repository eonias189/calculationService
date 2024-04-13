package use_register

import (
	"context"

	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Handler struct {
	e *Executor
}

func (h *Handler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	id, err := h.e.Do(int(req.MaxThreads))

	if err != nil {
		return nil, err
	}

	return &pb.RegisterResp{Id: id}, nil
}

func MakeHandler(as AgentService) Registerer {
	return &Handler{
		e: NewExecutor(as),
	}
}
