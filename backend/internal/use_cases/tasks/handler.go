package use_tasks

import (
	"errors"
	"net/http"
	"strconv"

	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func PostTaskHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.GetUserIdFromClaims(r)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		var body PostTaskBody
		err = utils.BindAndValidate(w, r, &body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		resp, err := e.PostTask(body, userId)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func GetTasksHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.GetUserIdFromClaims(r)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		limit := utils.GetIntQuery(r.URL.Query(), "limit", 10)
		offset := utils.GetIntQuery(r.URL.Query(), "offset", 0)

		resp, err := e.GetTasks(userId, limit, offset)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func GetTaskHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.GetUserIdFromClaims(r)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		resp, err := e.GetTask(int64(id), userId)
		if errors.Is(err, errs.ErrNotFound) {
			utils.HandleError(w, r, err, http.StatusNotFound)
			return
		}

		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func MakeHandler(taskService TaskService, timeoutsService TimeoutsService, distributer Distributer) http.Handler {
	r := chi.NewRouter()
	e := NewExecutor(taskService, timeoutsService, distributer)
	r.Post("/", PostTaskHandler(e))
	r.Get("/{id}", GetTaskHandler(e))
	r.Get("/", GetTasksHandler(e))

	return r
}
