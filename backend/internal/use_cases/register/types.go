package use_register

import (
	"context"

	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Registerer interface {
	Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error)
}
