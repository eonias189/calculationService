package useapi

import (
	"net/http"

	"github.com/eonias189/calculationService/backend/config"
	types "github.com/eonias189/calculationService/backend/interfaces"
)

type AgentApi struct {
	url    string
	scheme config.AgentScheme
	cli    *http.Client
}

func NewAgentApi(scheme config.AgentScheme, url string) *AgentApi {
	return &AgentApi{url: url, scheme: scheme, cli: &http.Client{}}
}

func (api *AgentApi) GetTaskStatus(id string) (types.TaskStatus, error) {
	resp := &types.GetTaskStatusResponse{}
	params := types.GetTaskStatusApi{
		RestParams: types.GetTaskStatusRestParams{
			ID: id,
		},
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.GetTaskStatus, params)
	if err != nil {
		return types.TaskStatus{}, err
	}
	return resp.TaskStatus, resp.GetError()
}
