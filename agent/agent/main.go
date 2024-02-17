package agent

import (
	"agent/internal/api"
	"fmt"
	"strconv"
	"strings"
	"time"

	c "backend/contract"

	"github.com/Knetic/govaluate"
)

type Agent struct {
	api     *api.OrchestratorApi
	threads int
	wp      *WorkerPool
}

type calculateTask struct {
	f func()
}

func (ct *calculateTask) Do() {
	ct.f()
}

func newTask(f func()) *calculateTask {
	return &calculateTask{f: f}
}

func (a *Agent) Calculate(task *c.Task, timeouts *c.Timeouts) {
	plusesCount := strings.Count(task.Expression, "+")
	minusesCount := strings.Count(task.Expression, "-")
	multpiplyCount := strings.Count(task.Expression, "*")
	divideCount := strings.Count(task.Expression, "/")
	evalExpr, err := govaluate.NewEvaluableExpression(task.Expression)
	if err != nil {
		a.api.SetResult(task.Id, 0, c.TaskStatus_executionError)
	}
	res, err := evalExpr.Evaluate(nil)
	if err != nil {
		a.api.SetResult(task.Id, 0, c.TaskStatus_executionError)
	}
	resString := fmt.Sprint(res)
	resInt, err := strconv.Atoi(resString)
	if err != nil {
		a.api.SetResult(task.Id, 0, c.TaskStatus_executionError)
	}
	time.Sleep(time.Second * time.Duration(plusesCount) * time.Duration(timeouts.Add))
	time.Sleep(time.Second * time.Duration(minusesCount) * time.Duration(timeouts.Substract))
	time.Sleep(time.Second * time.Duration(multpiplyCount) * time.Duration(timeouts.Multiply))
	time.Sleep(time.Second * time.Duration(divideCount) * time.Duration(timeouts.Divide))
	a.api.SetResult(task.Id, resInt, c.TaskStatus_done)

}

func (a *Agent) GetAgentStatus() (*c.AgentStatus, error) {
	fmt.Println("sending status")
	return &c.AgentStatus{MaxThreads: int64(a.threads), ExecutingThreads: int64(a.wp.ExecutingWorkers)}, nil
}

func (a *Agent) register(url string) {
	var err error
	err = a.api.Register(url)
	for err != nil {
		time.Sleep(time.Second * 5)
		err = a.api.Register(url)
	}
}

func (a *Agent) getTasks() {
	for {
		if a.wp.ExecutingWorkers == a.wp.MaxWorkers {
			time.Sleep(time.Second * 5)
			continue
		}
		task, timeouts, err := a.api.GetTask()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		a.wp.AddTask(newTask(func() {
			fmt.Println("start calculating")
			a.Calculate(task, timeouts)
			fmt.Println("calculated")
		}))
	}
}

func (a *Agent) Run(url string) {
	fmt.Println("starting at", url, "with", a.threads, "threads")
	a.wp.Start()
	go a.register(url)
	go a.getTasks()
}

func New(api *api.OrchestratorApi, threads int) *Agent {
	return &Agent{api: api, threads: threads, wp: NewWorkerPool(threads)}
}
