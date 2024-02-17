package agent

import (
	"agent/internal/api"
	"fmt"

	c "backend/contract"
)

type Agent struct {
	api     *api.OrchestratorApi
	threads int
}

func (o *Agent) GetAgentStatus() (c.AgentStatus, error) {
	fmt.Println("this is Agent")
	return c.AgentStatus{MaxThreads: 5, ExecutingThreads: 3}, nil
}

func (o *Agent) Run(url string) {
	fmt.Println("starting at", url, "with", o.threads, "threads")
	o.api.Register(url)
}

func New(api *api.OrchestratorApi, threads int) *Agent {
	return &Agent{api: api, threads: threads}
}
