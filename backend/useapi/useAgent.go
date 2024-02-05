package useapi

import (
	"net/http"

	"github.com/eonias189/calculationService/backend/config"
	"github.com/eonias189/calculationService/backend/interfaces"
)

type AgentApi struct {
	url    string
	scheme config.AgentScheme
	cli    *http.Client
}

func NewAgentApi(scheme config.AgentScheme, url string) *AgentApi {
	return &AgentApi{url: url, scheme: scheme, cli: &http.Client{}}
}

func (api *AgentApi) GetTaskStatus(id string) (interfaces.TaskStatus, error) {
	params := RequestParams[None]{Endpoint: api.scheme.GetTaskStatus, Body: None{}}
	resp, err := DoRequest[None, interfaces.TaskStatus](api.cli, api.url, params, id)
	return resp, err
}

func (api *AgentApi) IsWorking() bool {
	params := RequestParams[None]{Endpoint: api.scheme.IsWorking, Body: None{}}
	DoRequest[None, None](api.cli, api.url, params)
	return true
}
