package use_pong

import (
	"context"

	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Ponger interface {
	Pong(context.Context, *pb.PongReq) (*pb.Empty, error)
}
