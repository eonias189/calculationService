package use_connect

import (
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type AgentService interface {
	GetById(id int64) (service.Agent, error)
	Update(agent service.Agent) error
	Delete(id int64) error
}

type TaskService interface {
	GetById(id int64) (service.Task, error)
	Update(task service.Task) error
	SetUnexecutingForAgent(id int64) error
	GetExecutingForAgent(id int64) ([]service.TaskWithTimeouts, error)
}

type Distributor interface {
	Subscribe(id int64, maxTasks int) <-chan *pb.Task
	Unsubscribe(id int64) error
	Done(id int64) error
	Distribute(task *pb.Task) error
}
