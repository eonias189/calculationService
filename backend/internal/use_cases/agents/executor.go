package use_agents

import (
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type Executor struct {
	as AgentService
}

func (e *Executor) GetAgents(limit, offset int) (GetAgentsResp, error) {
	agents, err := e.as.GetAll(limit, offset)
	if err != nil {
		return GetAgentsResp{}, err
	}

	return GetAgentsResp{Agents: utils.Map(agents, func(a service.Agent) AgentSource {
		return AgentSource{
			Id:             a.Id,
			Ping:           a.Ping,
			Active:         a.Active,
			MaxThreads:     a.MaxThreads,
			RunningThreads: a.RunningThreads,
		}
	})}, nil
}

func NewExecutor(as AgentService) *Executor {
	return &Executor{as: as}
}
