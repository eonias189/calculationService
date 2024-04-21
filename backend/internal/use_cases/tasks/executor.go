package use_tasks

import (
	"errors"
	"math"
	"time"

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

func serviceTaskToTaskSource(task service.Task) TaskSource {
	if math.IsNaN(task.Result) {
		task.Result = 0
		task.Status = service.TaskStatusError
	}

	if math.IsInf(task.Result, 1) {
		task.Result = math.MaxFloat64
	}

	if math.IsInf(task.Result, -1) {
		task.Result = -math.MaxFloat64
	}

	return TaskSource{
		Id:         task.Id,
		Expression: task.Expression,
		Result:     task.Result,
		Status:     string(task.Status),
		CreateTime: task.CreateTime.Format(time.RFC3339),
	}
}

func (e *Executor) PostTask(body PostTaskBody, userId int64) (PostTaskResp, error) {
	task := service.Task{Expression: body.Expression, UserId: userId, Status: service.TaskStatusPending, CreateTime: time.Now()}
	id, err := e.taskService.Add(task)
	if err != nil {
		return PostTaskResp{}, err
	}

	task.Id = id

	timeouts, err := e.timeoutsService.GetForUser(userId)
	if errors.Is(err, errs.ErrNotFound) {
		err = nil
		timeouts = service.DefaultTimeouts
	}

	if err != nil {
		return PostTaskResp{}, err
	}

	err = e.distributer.Distribute(&pb.Task{Id: id, Expression: body.Expression, Timeouts: &pb.Timeouts{
		Add: uint64(timeouts.Add), Sub: uint64(timeouts.Sub), Mul: uint64(timeouts.Mul), Div: uint64(timeouts.Div),
	}})
	if err != nil {
		return PostTaskResp{}, err
	}

	return PostTaskResp{Task: serviceTaskToTaskSource(task)}, nil
}

func (e *Executor) GetTasks(userId int64, limit, offset int) (GetTasksResp, error) {
	tasks, err := e.taskService.GetAllForUser(userId, limit, offset)
	if err != nil {
		return GetTasksResp{}, err
	}

	return GetTasksResp{Tasks: utils.Map(tasks, serviceTaskToTaskSource)}, nil
}

func (e *Executor) GetTask(taskId int64, userId int64) (GetTaskResp, error) {
	task, err := e.taskService.GetById(taskId)
	if err != nil {
		return GetTaskResp{}, err
	}

	if task.UserId != userId {
		return GetTaskResp{}, errs.ErrNotFound
	}

	return GetTaskResp{Task: serviceTaskToTaskSource(task)}, nil
}

func NewExecutor(tasksService TaskService, timeoutsService TimeoutsService, distributer Distributer) *Executor {
	return &Executor{taskService: tasksService, timeoutsService: timeoutsService, distributer: distributer}
}
