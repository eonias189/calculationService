package server

import (
	"net/http"
	"strconv"

	"github.com/eonias189/calculationService/backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TaskAdder interface {
	AddTask(expression string) (int64, error)
}

type TasksGetter interface {
	Tasks(limit, offset int) ([]service.Task, error)
}

type TaskGetter interface {
	Task(id int64) (service.Task, error)
}

type PostTaskReq struct {
	Expression string `json:"expression" validate:"required"`
}

type PostTaskResp struct {
	Id int64 `json:"id"`
}

type GetTasksResp struct {
	Tasks []service.Task `json:"tasks"`
}

type GetTaskResp struct {
	Task service.Task `json:"task"`
}

func handlPostTask(ta TaskAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ParseBody[PostTaskReq](r)
		if err != nil {
			HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		id, err := ta.AddTask(body.Expression)
		if err != nil {
			HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, PostTaskResp{Id: id})
	}
}

func handleGetTasks(tg TasksGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := GetIntFromQuery(r.URL.Query(), "limit", 5)
		offset := GetIntFromQuery(r.URL.Query(), "offset", 0)

		tasks, err := tg.Tasks(limit, offset)
		if err != nil {
			HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, GetTasksResp{Tasks: tasks})

	}
}

func handleGetTask(tg TaskGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		task, err := tg.Task(int64(id))
		if err != nil {
			HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, GetTaskResp{Task: task})

	}
}
