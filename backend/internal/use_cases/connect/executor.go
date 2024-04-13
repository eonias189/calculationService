package use_connect

import (
	"time"

	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type Executor struct {
	agentService    AgentService
	taskService     TaskService
	timeoutsService TimeoutsService
	distributor     Distributor
}

func (e *Executor) OnConnClose(id int64) error {
	logger.Info("closing connection with", id)
	err := e.distributor.Unsubscribe(id)
	if err != nil {
		logger.Info("error while ubsubscribing")
		return err
	}

	agent, err := e.agentService.GetById(id)
	if err != nil {
		logger.Info("error while getting agent by id")
		return err
	}

	agent.Active = false
	agent.RunningThreads = 0
	err = e.agentService.Update(agent)
	if err != nil {
		logger.Info("error while updating agent")
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
				Add: task.Timeouts.Add,
				Sub: task.Timeouts.Sub,
				Mul: task.Timeouts.Mul,
				Div: task.Timeouts.Div,
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
		logger.Info("error while sending done")
		return err
	}

	agent, err := e.agentService.GetById(id)
	if err != nil {
		logger.Info("error while getting agent by id")
		return err
	}

	agent.Ping = time.Now().UnixNano() - result.SendTime
	agent.RunningThreads--
	err = e.agentService.Update(agent)
	if err != nil {
		logger.Info("error while updating agent")
		return err
	}

	task, err := e.taskService.GetById(result.TaskId)
	if err != nil {
		logger.Info("error while getting task by id")
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
		logger.Info("error while updating task")
		return err
	}

	return nil
}

func (e *Executor) OnTask(id int64, task *pb.Task) error {
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
	agent, err := e.agentService.GetById(id)
	if err != nil {
		return err
	}

	agent.Active = true
	err = e.agentService.Update(agent)
	if err != nil {
		logger.Info("error while updating agent")
		return err
	}

	return nil
}

func (e *Executor) GetTasks(id int64) (<-chan *pb.Task, error) {
	agent, err := e.agentService.GetById(id)
	if err != nil {
		logger.Info("error while getting agent by id")
		return nil, err
	}

	return e.distributor.Subscribe(id, agent.MaxThreads), nil
}

func NewExecutor(ts TaskService, as AgentService, tms TimeoutsService, d Distributor) *Executor {
	return &Executor{agentService: as, taskService: ts, timeoutsService: tms, distributor: d}
}
