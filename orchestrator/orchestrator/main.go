package orchestrator

import (
	"fmt"
	"orchestrator/internal/api"

	c "backend/contract"
)

type Orchestrator struct {
	api *api.AgentApi
}

func (o *Orchestrator) AddTask(expression string) error {
	fmt.Println("adding Task", expression)
	return nil
}
func (o *Orchestrator) GetTask() (c.Task, error) {
	fmt.Println("sending task")
	return c.Task{Id: "69", Expression: "1000 - 7", Timeouts: &c.Timeouts{Add: 4, Substract: 2, Multiply: 2, Divide: 88}}, nil
}
func (o *Orchestrator) GetTasks() ([]c.Task, error) {
	fmt.Println("sending tasks")
	return []c.Task{{Id: "69", Expression: "1000 - 7"}, {Id: "68", Expression: "993 + 7"}}, nil
}
func (o *Orchestrator) GetAgents() ([]c.AgentData, error) {
	fmt.Println("sending agents data")
	return []c.AgentData{
		{Ping: 83, Status: &c.AgentStatus{ExecutingThreads: 3, MaxThreads: 5}},
	}, nil
}
func (o *Orchestrator) GetTimeouts() (c.Timeouts, error) {
	fmt.Println("sending timeouts")
	return c.Timeouts{Add: 1, Substract: 2, Multiply: 3, Divide: 4}, nil
}
func (o *Orchestrator) SetTimeouts(timeouts c.Timeouts) error {
	fmt.Println("setting timeouts", timeouts.Add, timeouts.Substract, timeouts.Multiply, timeouts.Divide)
	return nil
}

func (o *Orchestrator) SetResult(id string, res int) error {
	fmt.Println("setting result", id, res)
	return nil
}
func (o *Orchestrator) Register(url string) error {
	fmt.Println("registring", url)
	return nil
}
func (o *Orchestrator) Run(url string) {
	fmt.Println("starting at", url)
}

func New(api *api.AgentApi) *Orchestrator {
	return &Orchestrator{api: api}
}
