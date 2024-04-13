package use_distribute

import (
	"context"

	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Distributable interface {
	Distribute(ctx context.Context, task *pb.Task) (*pb.Empty, error)
}
