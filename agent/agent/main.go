package agent

import (
	"agent/internal/api"
	"fmt"
	"strconv"
	"strings"
	"time"

	c "backend/contract"
	"backend/utils"

	"github.com/Knetic/govaluate"
)

type Agent struct {
	api     *api.OrchestratorApi
	threads int
	id      int
	wp      *utils.WorkerPool
}

func (a *Agent) Calculate(task *c.Task, timeouts *c.Timeouts) {
	plusesCount := strings.Count(task.Expression, "+")
	minusesCount := strings.Count(task.Expression, "-")
	multpiplyCount := strings.Count(task.Expression, "*")
	divideCount := strings.Count(task.Expression, "/")
	evalExpr, err := govaluate.NewEvaluableExpression(task.Expression)
	defer func() {
		err := recover()
		if err != nil {
			a.api.SetResult(task.Id, 0, c.TaskStatus_executionError)
		}
	}()
	var resutl int
	var status c.TaskStatus
	defer func() {
		err := a.api.SetResult(task.Id, resutl, status)
		for err != nil {
			time.Sleep(time.Second * 5)
			err = a.api.SetResult(task.Id, resutl, status)
		}
	}()
	if err != nil {
		status = c.TaskStatus_executionError
		return
	}
	res, err := evalExpr.Evaluate(nil)
	if err != nil {
		status = c.TaskStatus_executionError
		return
	}
	resString := fmt.Sprint(res)
	resInt, err := strconv.Atoi(resString)
	if err != nil {
		status = c.TaskStatus_executionError
		return
	}
	time.Sleep(time.Second * time.Duration(plusesCount) * time.Duration(timeouts.Add))
	time.Sleep(time.Second * time.Duration(minusesCount) * time.Duration(timeouts.Substract))
	time.Sleep(time.Second * time.Duration(multpiplyCount) * time.Duration(timeouts.Multiply))
	time.Sleep(time.Second * time.Duration(divideCount) * time.Duration(timeouts.Divide))
	resutl = resInt
	status = c.TaskStatus_done

}

func (a *Agent) GetAgentStatus() (*c.AgentStatus, error) {
	fmt.Println("sending status")
	return &c.AgentStatus{MaxThreads: int64(a.threads), ExecutingThreads: int64(a.wp.ExecutingWorkers)}, nil
}

func (a *Agent) register(url string) {
	var err error
	id, err := a.api.Register(url)
	for err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
		id, err = a.api.Register(url)
	}
	a.id = id

}

func (a *Agent) getTasks() {
	for {
		if a.wp.ExecutingWorkers == a.wp.MaxWorkers {
			time.Sleep(time.Second * 5)
			continue
		}
		task, timeouts, err := a.api.GetTask(a.id)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		a.wp.AddTask(utils.NewTask(func() {
			fmt.Println("start calculating")
			a.Calculate(task, timeouts)
			fmt.Println("calculated")
		}))
		time.Sleep(time.Second * 1)
	}
}

func (a *Agent) Run(url string) {
	fmt.Println("starting at", url, "with", a.threads, "threads")
	a.wp.Start()
	a.register(url)
	go a.getTasks()
}

func New(api *api.OrchestratorApi, threads int) *Agent {
	return &Agent{api: api, threads: threads, wp: utils.NewWorkerPool(threads)}
}
