package server

import (
	"fmt"
	"net/http"

	c "github.com/eonias189/calculationService/agent/contract"
	"github.com/eonias189/calculationService/agent/internal/config"
	"github.com/eonias189/calculationService/agent/internal/utils"
	"github.com/eonias189/calculationService/agent/pkg/agent"
)

type Server struct {
	Url      string
	scheme   config.AgentScheme
	agent    *agent.Agent
	handlers map[string]http.Handler
}

func (s *Server) handleGetAgentStatus(w http.ResponseWriter, r *http.Request) {
	resp := c.GetAgentStatusResponse{}
	defer utils.SendResponse(&resp, w)

	resp.Status, resp.Error = s.agent.GetAgentStatus()
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

	s.Handle(s.scheme.GetAgentStatus, s.handleGetAgentStatus)

	for pattern, handler := range s.handlers {
		fmt.Printf("handling %v\n", pattern)
		mux.Handle(pattern, handler)
	}

	s.agent.Run(s.Url)
	http.ListenAndServe(s.Url, utils.LogMiddleware(mux))
}

func New(agent *agent.Agent, url string, scheme config.AgentScheme) *Server {
	return &Server{agent: agent, scheme: scheme, Url: url, handlers: make(map[string]http.Handler)}
}
