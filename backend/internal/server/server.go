package server

import (
	"context"
	"net/http"

	"github.com/eonias189/calculationService/backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Orchestrator interface {
	Tasks(limit, offset int) ([]service.Task, error)
	Task(int64) (service.Task, error)
	AddTask(expression string) (int64, error)
	Agents(limit, offset int) ([]service.Agent, error)
	Timeouts() (service.Timeouts, error)
	SetTimeouts(service.Timeouts) error
}

type RestServer struct {
	orchestrator Orchestrator
}

type ErrorResp struct {
	Reason string `json:"reason"`
}

func (s *RestServer) Run(ctx context.Context, addr string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{}))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, struct{ Message string }{Message: "pong"})
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", handlPostTask(s.orchestrator))
		r.Get("/", handleGetTasks(s.orchestrator))
		r.Get("/{id}", handleGetTask(s.orchestrator))
	})

	r.Route("/agents", func(r chi.Router) {
		r.Get("/", handleGetAgents(s.orchestrator))
	})

	return http.ListenAndServe(addr, r)
}

func New(orchestrator Orchestrator) *RestServer {
	return &RestServer{orchestrator: orchestrator}
}
