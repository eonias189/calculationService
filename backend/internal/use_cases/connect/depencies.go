package use_connect

import "github.com/eonias189/calculationService/backend/internal/service"

type AgentService interface {
	GetById(id int64) (service.Agent, error)
	Update(agent service.Agent) error
	Delete(id int64) error
}

type TaskService interface {
	Update(task service.Task) error
}

type TimeoutsService interface {
	Load(userId int64) (service.Timeouts, error)
}
