package use_tasks

import (
	"errors"

	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type Executor struct {
	taskService     TaskService
	timeoutsService TimeoutsService
	distributer     Distributer
}

func (e *Executor) PostTask(body PostTaskBody, userId int64) (PostTaskResp, error) {
	id, err := e.taskService.Add(service.Task{Expression: body.Expression, UserId: userId, Status: service.TaskStatusPending})
	if err != nil {
		return PostTaskResp{}, err
	}

	timeouts, err := e.timeoutsService.GetForUser(userId)
	if errors.Is(err, errs.ErrNotFound) {
		err = nil
		timeouts = service.DefaultTimeouts
	}

	if err != nil {
		return PostTaskResp{}, err
	}

	err = e.distributer.Distribute(&pb.Task{Id: id, Expression: body.Expression, Timeouts: &pb.Timeouts{
		Add: timeouts.Add, Sub: timeouts.Sub, Mul: timeouts.Mul, Div: timeouts.Div,
	}})
	if err != nil {
		return PostTaskResp{}, err
	}

	return PostTaskResp{Task: TaskSource{Id: id, Expression: body.Expression, Result: 0, Status: string(service.TaskStatusPending)}}, nil
}

func (e *Executor) GetTasks(userId int64, limit, offset int) (GetTasksResp, error) {
	tasks, err := e.taskService.GetAllForUser(userId, limit, offset)
	if err != nil {
		return GetTasksResp{}, err
	}

	return GetTasksResp{Tasks: utils.Map(tasks, func(task service.Task) TaskSource {
		return TaskSource{Id: task.Id, Expression: task.Expression, Result: task.Result, Status: string(task.Status)}
	})}, nil
}

func (e *Executor) GetTask(taskId int64, userId int64) (GetTaskResp, error) {
	task, err := e.taskService.GetById(taskId)
	if err != nil {
		return GetTaskResp{}, err
	}

	if task.UserId != userId {
		return GetTaskResp{}, errs.ErrNotFound
	}

	return GetTaskResp{Task: TaskSource{Id: taskId, Expression: task.Expression, Result: task.Result, Status: string(task.Status)}}, nil
}

func NewExecutor(tasksService TaskService, timeoutsService TimeoutsService, distributer Distributer) *Executor {
	return &Executor{taskService: tasksService, timeoutsService: timeoutsService, distributer: distributer}
}
