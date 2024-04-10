package use_register

import "github.com/eonias189/calculationService/orchestrator/internal/service"

type AgentService interface {
	Add(agent service.Agent) (int64, error)
}
