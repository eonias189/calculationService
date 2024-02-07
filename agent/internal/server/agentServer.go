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
	handlers map[string]http.HandlerFunc
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	resp := c.PingResponse{}
	defer utils.SendResponse(&resp, w)

	ok, err := s.agent.Ping()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Ok = ok
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

	s.Handle(s.scheme.Ping, s.handlePing)

	for pattern, handler := range s.handlers {
		fmt.Printf("handling %v\n", pattern)
		mux.HandleFunc(pattern, handler)
	}

	s.agent.Run(s.Url)
	http.ListenAndServe(s.Url, utils.LogMiddleware(mux))
}

func New(agent *agent.Agent, url string, scheme config.AgentScheme) *Server {
	return &Server{agent: agent, scheme: scheme, Url: url, handlers: make(map[string]http.HandlerFunc)}
}
