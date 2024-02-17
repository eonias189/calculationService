package server

import (
	"agent/agent"
	c "backend/contract"
	"backend/utils"
	"fmt"
	"net/http"
)

type Server struct {
	mux   *http.ServeMux
	Port  int
	agent *agent.Agent
}

func (s *Server) Handle(pattern, method string, handler http.HandlerFunc) {
	fmt.Println("handling", pattern)
	s.mux.Handle(pattern, utils.CheckMethodMiddleware(handler, method))
}

func (s *Server) handleGetStatus(w http.ResponseWriter, r *http.Request) {
	resp := c.GetAgentStatusResp{}
	status, err := s.agent.GetAgentStatus()
	if err != nil {
		utils.SendError(err, w)
		return
	}
	resp.Status = &status
	utils.SendResponse(&resp, w)
}

func (s *Server) Run() {
	s.agent.Run(fmt.Sprintf("http://127.0.0.1:%v", s.Port))

	s.Handle("/status/", "get", s.handleGetStatus)

	http.ListenAndServe(fmt.Sprintf(":%v", s.Port), utils.LogMiddleware(s.mux))
}

func NewServer(agent *agent.Agent, port int) *Server {
	return &Server{mux: http.NewServeMux(), Port: port, agent: agent}
}
