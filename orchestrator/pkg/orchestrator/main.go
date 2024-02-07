package orchestrator

import (
	"fmt"
	"strconv"

	c "github.com/eonias189/calculationService/orchestrator/contract"
)

type Orchestrator struct {
}

func (o *Orchestrator) AddTask(c.Task) error {
	fmt.Println("this is orchestrator")
	return nil
}
func (o *Orchestrator) GetTasksStatus() ([]c.TaskStatus, error) {
	return []c.TaskStatus{}, nil
}
func (o *Orchestrator) GetResult(id string) (int, error) {
	res, _ := strconv.Atoi(id)
	return res, nil
}
func (o *Orchestrator) GetOperationsTimeouts() (c.OperationsTimeouts, error) {
	return c.OperationsTimeouts{}, nil
}
func (o *Orchestrator) SetOperationsTimeouts(c.OperationsTimeouts) error {
	return nil
}
func (o *Orchestrator) GetTask() (c.Task, error) {
	return c.Task{}, nil
}
func (o *Orchestrator) SetResult(string, int) error {
	return nil
}
func (o *Orchestrator) Register(string) error {
	return nil
}
func (o *Orchestrator) Run(url string) {
	fmt.Println("starting at", url)
}

func New() *Orchestrator {
	return &Orchestrator{}
}
