package use_distribute

import pb "github.com/eonias189/calculationService/backend/internal/proto"

type Distributor interface {
	Distribute(task *pb.Task) error
}
