package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/eonias189/calculationService/backend/config"
	"github.com/eonias189/calculationService/backend/handleapi"
	types "github.com/eonias189/calculationService/backend/interfaces"
	"github.com/eonias189/calculationService/backend/useapi"
)

type Orchestrator struct {
}

func (o *Orchestrator) AddTask(task types.Task) error {
	fmt.Println(task.ID, task.Expression)
	return nil
}
func (o *Orchestrator) GetTasksStatus() ([]types.TaskStatus, error) {
	return []types.TaskStatus{{Done: true, ID: "5"}}, nil
}

func (o *Orchestrator) GetResult(id string) (int, error) {
	fmt.Println(id)
	return 69, fmt.Errorf(id)
}
func (o *Orchestrator) GetOperationsTimeouts() (types.OperationsTimeouts, error) {
	return types.OperationsTimeouts{Add: 5}, nil
}
func (o *Orchestrator) SetOperationsTimeouts(timeouts types.OperationsTimeouts) error {
	fmt.Println(timeouts)
	return nil
}
func (o *Orchestrator) GetTask() (types.Task, error) {
	return types.Task{ID: "5", Expression: "1000 - 7"}, nil
}

func (o *Orchestrator) SetResult(id string, num int) error {
	fmt.Println(id, num)
	return nil
}
func (o *Orchestrator) Register(url string) error {
	fmt.Println(url)
	return nil
}
func (o *Orchestrator) Run(url string) {
	fmt.Println("starting at", url)
}

type Agent struct {
}

func (a *Agent) GetTaskStatus(id string) (types.TaskStatus, error) {
	fmt.Println(id)
	return types.TaskStatus{Done: true, Err: false, ID: "22"}, nil
}

func (a *Agent) Run(url string) {
	fmt.Println("starting at", url)
}

func TestServers() {
	asp := config.NewApiSchemeProvider()
	// fmt.Println(asp.GetOrchestratorScheme())
	orch := Orchestrator{}
	agent := Agent{}

	oServ := handleapi.NewOrchestratorServer(&orch, ":8081", asp.GetOrchestratorScheme())
	aServ := handleapi.NewAgentServer(&agent, ":8082", asp.GetAgentScheme())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		oServ.Run()
	}()
	go func() {
		defer wg.Done()
		aServ.Run()
	}()
	wg.Wait()
}

func TestApi() {
	asp := config.NewApiSchemeProvider()
	orch := useapi.NewOrchestratorApi(asp.GetOrchestratorScheme(), "http://127.0.0.1:8081")
	ag := useapi.NewAgentApi(asp.GetAgentScheme(), "http://127.0.0.1:8082")
	fmt.Println(orch.AddTask(types.Task{ID: "some id"}))
	fmt.Println(orch.GetOperationsTimeouts())
	fmt.Println(orch.GetResult("5"))
	fmt.Println(orch.GetTask())
	fmt.Println(orch.GetTasksStatus())
	fmt.Println(orch.Register("lala"))
	fmt.Println(orch.SetResult("5", 5))
	fmt.Println(ag.GetTaskStatus("5"))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		TestServers()
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 10)
		TestApi()
	}()
	wg.Wait()

}
