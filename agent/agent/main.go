package agent

import (
	"agent/internal/api"
	"fmt"

	c "backend/contract"
)

type Agent struct {
	api *api.OrchestratorApi
}

func (o *Agent) GetAgentStatus() (c.AgentStatus, error) {
	fmt.Println("this is Agent")
	return c.AgentStatus{MaxThreads: 5, ExecutingThreads: 3}, nil
}

func (o *Agent) Run(url string) {
	fmt.Println("starting at", url)
	o.api.Register(url)
}

func New(api *api.OrchestratorApi) *Agent {
	return &Agent{api: api}
}
