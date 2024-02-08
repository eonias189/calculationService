package orchestrator

import (
	"fmt"

	c "github.com/eonias189/calculationService/orchestrator/contract"
)

type Orchestrator struct {
}

func (o *Orchestrator) AddTask(task c.Task) error {
	fmt.Println("adding Task", task)
	return nil
}
func (o *Orchestrator) GetTask() (c.Task, error) {
	fmt.Println("sending task")
	return c.Task{ID: "69", Expression: "1000 - 7"}, nil
}
func (o *Orchestrator) GetTasksData() ([]c.TaskData, error) {
	fmt.Println("sending task data")
	return []c.TaskData{{Task: c.Task{ID: "59", Expression: "1000 - 7"}, Status: c.TaskStatus{Done: false, Err: false}}}, nil
}
func (o *Orchestrator) GetAgentsData() ([]c.AgentData, error) {
	fmt.Println("sending agents data")
	return []c.AgentData{
		{Ping: 83, AgentStatus: c.AgentStatus{
			MaxThreadsNumber: 5, ThreadsRuning: 3,
		}},
	}, nil
}
func (o *Orchestrator) GetOperationsTimeouts() (c.OperationsTimeouts, error) {
	fmt.Println("sending timeouts")
	return c.OperationsTimeouts{Add: 1, Subtract: 2, Multiply: 3, Divide: 4}, nil
}
func (o *Orchestrator) SetOperationsTimeouts(timeouts c.OperationsTimeouts) error {
	fmt.Println("setting timeouts", timeouts)
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

func New() *Orchestrator {
	return &Orchestrator{}
}
