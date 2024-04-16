package use_tasks

import (
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type TaskService interface {
	Add(service.Task) (int64, error)
	GetById(id int64) (service.Task, error)
	GetAllForUser(userId int64, limit, offset int) ([]service.Task, error)
}

type TimeoutsService interface {
	GetForUser(id int64) (service.Timeouts, error)
}

type Distributer interface {
	Distribute(task *pb.Task) error
}
