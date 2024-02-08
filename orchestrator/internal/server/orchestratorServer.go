package server

import (
	"fmt"
	"net/http"

	c "github.com/eonias189/calculationService/orchestrator/contract"
	"github.com/eonias189/calculationService/orchestrator/internal/config"
	"github.com/eonias189/calculationService/orchestrator/internal/utils"
	"github.com/eonias189/calculationService/orchestrator/pkg/orchestrator"
)

type Server struct {
	Url          string
	scheme       config.OrchestratorScheme
	orchestrator *orchestrator.Orchestrator
	handlers     map[string]http.Handler
}

func (s *Server) handleAddTask(w http.ResponseWriter, r *http.Request) {
	resp := c.ErrorResponse{}
	defer utils.SendResponse(&resp, w)

	body, err := utils.GetBody[c.AddTaskBody](r)
	if err != nil {
		resp.Error = err
		return
	}

	err = s.orchestrator.AddTask(body.Task)
	resp.Error = err
}

func (s *Server) handleGetTask(w http.ResponseWriter, r *http.Request) {
	resp := c.GetTaskResponse{}
	defer utils.SendResponse(&resp, w)

	resp.Task, resp.Error = s.orchestrator.GetTask()
}

func (s *Server) handleGetTasksData(w http.ResponseWriter, r *http.Request) {
	resp := c.GetTasksDataResponse{}
	defer utils.SendResponse(&resp, w)

	resp.TasksData, resp.Error = s.orchestrator.GetTasksData()
}
func (s *Server) handleGetAgentsData(w http.ResponseWriter, r *http.Request) {
	resp := c.GetAgentsDataResponse{}
	defer utils.SendResponse(&resp, w)

	resp.Data, resp.Error = s.orchestrator.GetAgentsData()
}
func (s *Server) handleGetOperationsTimeouts(w http.ResponseWriter, r *http.Request) {
	resp := c.GetOperationsTimeoutsResponse{}
	defer utils.SendResponse(&resp, w)

	resp.OperationsTimeouts, resp.Error = s.orchestrator.GetOperationsTimeouts()
}
func (s *Server) handleSetOperationsTimeouts(w http.ResponseWriter, r *http.Request) {
	resp := c.ErrorResponse{}
	defer utils.SendResponse(&resp, w)

	body, err := utils.GetBody[c.SetOperationsTimeoutsBody](r)
	if err != nil {
		resp.Error = err
		return
	}
	resp.Error = s.orchestrator.SetOperationsTimeouts(body.OperationsTimeouts)
}
func (s *Server) handleSetResult(w http.ResponseWriter, r *http.Request) {
	resp := c.ErrorResponse{}
	defer utils.SendResponse(&resp, w)

	body, err := utils.GetBody[c.SetResultBody](r)
	if err != nil {
		resp.Error = err
		return
	}
	resp.Error = s.orchestrator.SetResult(body.ID, body.Number)
}
func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	resp := c.ErrorResponse{}
	defer utils.SendResponse(&resp, w)

	body, err := utils.GetBody[c.RegisterBody](r)
	if err != nil {
		resp.Error = err
		return
	}
	resp.Error = s.orchestrator.Register(body.Url)
}

func (s *Server) Handle(endpoint config.Endpoint, handler http.HandlerFunc) {
	var pattern string
	if endpoint.Method == "GET" {
		pattern = "/" + endpoint.Url + "/"
	} else {
		pattern = "/" + endpoint.Url
	}
	s.handlers[pattern] = utils.CheckMethodMiddleware(handler, endpoint.Method)
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	s.Handle(s.scheme.AddTask, s.handleAddTask)
	s.Handle(s.scheme.GetTask, s.handleGetTask)
	s.Handle(s.scheme.GetTasksData, s.handleGetTasksData)
	s.Handle(s.scheme.GetAgentsData, s.handleGetAgentsData)
	s.Handle(s.scheme.GetOperationsTimeouts, s.handleGetOperationsTimeouts)
	s.Handle(s.scheme.SetOperationsTimeouts, s.handleSetOperationsTimeouts)
	s.Handle(s.scheme.SetResult, s.handleSetResult)
	s.Handle(s.scheme.Register, s.handleRegister)

	for pattern, handler := range s.handlers {
		fmt.Printf("handling %v\n", pattern)
		mux.Handle(pattern, handler)
	}

	s.orchestrator.Run(s.Url)
	http.ListenAndServe(s.Url, utils.LogMiddleware(mux))
}

func New(orchestrator *orchestrator.Orchestrator, url string, scheme config.OrchestratorScheme) *Server {
	return &Server{orchestrator: orchestrator, scheme: scheme, Url: url, handlers: make(map[string]http.Handler)}
}
