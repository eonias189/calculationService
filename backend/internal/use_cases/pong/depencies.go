package use_pong

import "github.com/eonias189/calculationService/backend/internal/service"

type AgentService interface {
	GetById(id int64) (service.Agent, error)
	Update(agent service.Agent) error
}
