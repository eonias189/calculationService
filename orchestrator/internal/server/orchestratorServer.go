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
	handlers     map[string]http.HandlerFunc
}

func (s *Server) handleAddTask(w http.ResponseWriter, r *http.Request) {
	resp := c.ErrorResponse{}
	defer utils.SendResponse(&resp, w)

	body, err := utils.GetBody[c.AddTaskBody](r)
	if err != nil {
		resp.Error = err.Error()
		return
	}

	err = s.orchestrator.AddTask(body.Task)
	if err != nil {
		resp.Error = err.Error()
	}
}

func (s *Server) handleGetResult(w http.ResponseWriter, r *http.Request) {
	resp := c.GetResultResponse{}
	defer utils.SendResponse(&resp, w)

	params, err := utils.ParseRestParams[c.GetResultRestParams](r, s.scheme.GetResult)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	res, err := s.orchestrator.GetResult(params.ID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Number = res

}

func (s *Server) Handle(endpoint config.Endpoint, handler http.HandlerFunc) {
	var pattern string
	if endpoint.Method == "GET" {
		pattern = "/" + endpoint.Url + "/"
	} else {
		pattern = "/" + endpoint.Url
	}
	s.handlers[pattern] = handler
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	s.Handle(s.scheme.AddTask, s.handleAddTask)
	s.Handle(s.scheme.GetResult, s.handleGetResult)

	for pattern, handler := range s.handlers {
		fmt.Printf("handling %v\n", pattern)
		mux.HandleFunc(pattern, handler)
	}

	s.orchestrator.Run(s.Url)
	http.ListenAndServe(s.Url, utils.LogMiddleware(mux))
}

func New(orchestrator *orchestrator.Orchestrator, url string, scheme config.OrchestratorScheme) *Server {
	return &Server{orchestrator: orchestrator, scheme: scheme, Url: url, handlers: make(map[string]http.HandlerFunc)}
}
