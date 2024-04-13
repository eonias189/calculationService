package use_distribute

import (
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Executor struct {
	distributor Distributor
}

func (e *Executor) Do(task *pb.Task) error {
	logger.Info("distributing", task.String())
	return e.distributor.Distribute(task)
}

func NewExecutor(d Distributor) *Executor {
	return &Executor{distributor: d}
}
