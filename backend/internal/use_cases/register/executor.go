package use_register

import "github.com/eonias189/calculationService/backend/internal/service"

type Executor struct {
	as AgentService
}

func (s *Executor) Do(maxThreads int) (int64, error) {
	return s.as.Add(service.Agent{MaxThreads: maxThreads})
}

func NewExecutor(as AgentService) *Executor {
	return &Executor{as: as}
}
