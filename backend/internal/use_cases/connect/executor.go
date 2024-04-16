package use_connect

import (
	"time"

	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type Executor struct {
	agentService AgentService
	taskService  TaskService
	distributor  Distributor
}

func (e *Executor) OnConnClose(id int64) error {
	logger.Info("closing connection with", id)
	err := e.distributor.Unsubscribe(id)
	if err != nil {
		return err
	}

	agent, err := e.agentService.GetById(id)
	if err != nil {
		return err
	}

	agent.Active = false
	agent.RunningThreads = 0
	err = e.agentService.Update(agent)
	if err != nil {
		return err
	}

	tasks, err := e.taskService.GetExecutingForAgent(id)
	if err != nil {
		return err
	}

	err = e.taskService.SetUnexecutingForAgent(id)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Status == service.TaskStatusExecuting {
			err = e.distributor.Distribute(&pb.Task{Id: task.Task.Id, Expression: task.Task.Expression, Timeouts: &pb.Timeouts{
				Add: uint64(task.Timeouts.Add),
				Sub: uint64(task.Timeouts.Sub),
				Mul: uint64(task.Timeouts.Mul),
				Div: uint64(task.Timeouts.Div),
			}})
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (e *Executor) OnResult(id int64, result *pb.ResultResp) error {
	logger.Info("got result", result.String())
	err := e.distributor.Done(id)
	if err != nil {
		return err
	}

	agent, err := e.agentService.GetById(id)
	if err != nil {
		return err
	}

	agent.RunningThreads = int(result.RunningThreads)
	agent.Ping = time.Now().UnixNano() - result.SendTime
	err = e.agentService.Update(agent)
	if err != nil {
		return err
	}

	task, err := e.taskService.GetById(result.TaskId)
	if err != nil {
		return err
	}

	task.Result = result.Result
	if result.Error {
		task.Status = service.TaskStatusError
	} else {
		task.Status = service.TaskStatusSuccess
	}

	err = e.taskService.Update(task)
	if err != nil {
		return err
	}

	return nil
}

func (e *Executor) OnTask(id int64, task *pb.Task) error {
	logger.Info("sending", task.String(), "to", id)
	t, err := e.taskService.GetById(task.Id)
	if err != nil {
		return err
	}

	t.Executor = id
	t.Status = service.TaskStatusExecuting
	err = e.taskService.Update(t)
	if err != nil {
		return err
	}

	agent, err := e.agentService.GetById(id)
	if err != nil {
		return err
	}

	agent.RunningThreads++
	err = e.agentService.Update(agent)
	if err != nil {
		return err
	}

	return nil
}

func (e *Executor) OnStart(id int64) error {
	logger.Info(id, "connected")
	agent, err := e.agentService.GetById(id)
	if err != nil {
		return err
	}

	agent.Active = true
	err = e.agentService.Update(agent)
	if err != nil {
		return err
	}

	return nil
}

func (e *Executor) GetTasks(id int64) (<-chan *pb.Task, error) {
	agent, err := e.agentService.GetById(id)
	if err != nil {
		return nil, err
	}

	return e.distributor.Subscribe(id, agent.MaxThreads), nil
}

func NewExecutor(ts TaskService, as AgentService, d Distributor) *Executor {
	return &Executor{agentService: as, taskService: ts, distributor: d}
}
