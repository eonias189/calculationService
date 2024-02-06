package handleapi

import (
	"fmt"
	"net/http"

	"github.com/eonias189/calculationService/backend/config"
	types "github.com/eonias189/calculationService/backend/interfaces"
)

type AgentServer struct {
	agent types.IAgent
	api   config.AgentScheme
	Url   string
}

func (aServ *AgentServer) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	resp := types.GetTaskStatusResponse{}
	defer SendResponse(&resp, w)
	params, err := ParseRestParams[types.GetTaskStatusRestParams](r, aServ.api.GetTaskStatus)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	res, err := aServ.agent.GetTaskStatus(params.ID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.TaskStatus = res
}

func (aServ *AgentServer) Run() {
	aServ.agent.Run(aServ.Url)
	mux := http.NewServeMux()
	handlers := map[string]http.HandlerFunc{
		aServ.api.GetTaskStatus.Url: aServ.GetTaskStatus,
	}
	for endpoint, handler := range handlers {
		pattern := "/" + endpoint + "/"
		fmt.Printf("handling %v at %v\n", handler, pattern)
		mux.HandleFunc(pattern, handler)
	}
	http.ListenAndServe(aServ.Url, LogErrorMiddleware(mux))
}

func NewAgentServer(agent types.IAgent, url string, scheme config.AgentScheme) *AgentServer {
	return &AgentServer{agent: agent, Url: url, api: scheme}
}
