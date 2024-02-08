package agent

import (
	"fmt"

	c "github.com/eonias189/calculationService/agent/contract"
)

type Agent struct {
}

func (o *Agent) GetAgentStatus() (c.AgentStatus, error) {
	fmt.Println("this is Agent")
	return c.AgentStatus{MaxThreadsNumber: 5, ThreadsRuning: 3}, nil
}

func (o *Agent) Run(url string) {
	fmt.Println("starting at", url)
}

func New() *Agent {
	return &Agent{}
}
