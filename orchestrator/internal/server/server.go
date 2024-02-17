package server

import (
	c "backend/contract"
	"backend/utils"
	"fmt"
	"net/http"
	"orchestrator/orchestrator"
)

type Server struct {
	Port         string
	orchestrator *orchestrator.Orchestrator
	mux          *http.ServeMux
}

func (s *Server) handleAddTask(w http.ResponseWriter, r *http.Request) {
	resp := c.AddTaskResp{Ok: true}

	body, err := utils.GetBody[c.AddTaskBody](r)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	err = s.orchestrator.AddTask(body.Expression)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	utils.SendResponse(&resp, w)
}

func (s *Server) handleGetTask(w http.ResponseWriter, r *http.Request) {
	resp := c.GetTaskResp{}
	body, err := utils.GetBody[c.GetTaskBody](r)
	if err != nil {
		utils.SendError(err, w)
	}
	task, timeouts, err := s.orchestrator.GetTask(int(body.AgentId))
	if err != nil {
		utils.SendError(err, w)
		return
	}
	resp.Task = task
	resp.Timeouts = timeouts

	utils.SendResponse(&resp, w)
}

func (s *Server) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	resp := c.GetTasksResp{}

	tasks, err := s.orchestrator.GetTasks()
	if err != nil {
		utils.SendError(err, w)
		return
	}
	resp.Tasks = tasks
	utils.SendResponse(&resp, w)
}

func (s *Server) handleGetAgents(w http.ResponseWriter, r *http.Request) {
	resp := c.GetAgentsResp{}

	agents, err := s.orchestrator.GetAgents()
	if err != nil {
		utils.SendError(err, w)
		return
	}
	resp.Agents = agents

	utils.SendResponse(&resp, w)
}

func (s *Server) handleGetTimeouts(w http.ResponseWriter, r *http.Request) {
	resp := c.GetTimeoutsResp{}

	timeouts, err := s.orchestrator.GetTimeouts()
	if err != nil {
		utils.SendError(err, w)
		return
	}

	resp.Timeouts = timeouts
	utils.SendResponse(&resp, w)
}

func (s *Server) handleSetTimeouts(w http.ResponseWriter, r *http.Request) {
	resp := c.SetTimeoutsResp{Ok: true}
	body, err := utils.GetBody[c.SetTimeoutsBody](r)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	err = s.orchestrator.SetTimeouts(body.Timeouts)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	utils.SendResponse(&resp, w)
}

func (s *Server) handleSetResult(w http.ResponseWriter, r *http.Request) {
	resp := c.SetResultResp{Ok: true}

	body, err := utils.GetBody[c.SetResultBody](r)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	err = s.orchestrator.SetResult(body.Id, int(body.Result), body.Status)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	utils.SendResponse(&resp, w)
}

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	resp := c.RegisterResp{}

	body, err := utils.GetBody[c.RegisterBody](r)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	id, err := s.orchestrator.Register(body.Url)
	if err != nil {
		utils.SendError(err, w)
		return
	}
	resp.Id = int64(id)
	utils.SendResponse(&resp, w)
}

func (s *Server) Handle(pattern, method string, handler http.HandlerFunc) {
	fmt.Println("handling", pattern)
	s.mux.Handle(pattern, utils.CheckMethodMiddleware(handler, method))
}

func setCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Run() {
	s.orchestrator.Run("http://127.0.0.1" + s.Port)

	s.Handle("/addTask", "post", s.handleAddTask)
	s.Handle("/getTask", "post", s.handleGetTask)
	s.Handle("/getTasks", "get", s.handleGetTasks)
	s.Handle("/getAgents", "get", s.handleGetAgents)
	s.Handle("/getTimeouts", "get", s.handleGetTimeouts)
	s.Handle("/setTimeouts", "post", s.handleSetTimeouts)
	s.Handle("/setResult", "post", s.handleSetResult)
	s.Handle("/register", "post", s.handleRegister)

	http.ListenAndServe(s.Port, utils.LogMiddleware(setCorsMiddleware(s.mux)))

}

func NewServer(orchestrator *orchestrator.Orchestrator, port string) *Server {
	return &Server{Port: port, orchestrator: orchestrator, mux: http.NewServeMux()}
}
