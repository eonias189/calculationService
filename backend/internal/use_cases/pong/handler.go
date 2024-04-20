package use_pong

import (
	"context"
	"time"

	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

func MakeHandler(agentService AgentService) Ponger {
	return &Handler{as: agentService}
}

type Handler struct {
	as AgentService
}

func (h *Handler) Pong(ctx context.Context, req *pb.PongReq) (*pb.Empty, error) {
	ping := (time.Now().UnixNano() - req.SentTime) / 1000000
	logger.Info("agent", req.Id, "ping", ping)
	agent, err := h.as.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	agent.Ping = ping
	err = h.as.Update(agent)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
