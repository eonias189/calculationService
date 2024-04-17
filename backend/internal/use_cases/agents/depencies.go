package use_agents

import "github.com/eonias189/calculationService/backend/internal/service"

type AgentService interface {
	GetAll(limit, offset int) ([]service.Agent, error)
}
