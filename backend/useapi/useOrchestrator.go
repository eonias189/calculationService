package useapi

import (
	"net/http"

	"github.com/eonias189/calculationService/backend/config"
	"github.com/eonias189/calculationService/backend/interfaces"
)

type OrchestratorApi struct {
	url    string
	scheme config.OrchestratorScheme
	cli    *http.Client
}

func NewOrchestratorApi(scheme config.OrchestratorScheme, url string) *OrchestratorApi {
	return &OrchestratorApi{scheme: scheme, cli: &http.Client{}, url: url}
}

func (api *OrchestratorApi) AddTask(task interfaces.Task) error {
	params := RequestParams[interfaces.Task]{Endpoint: api.scheme.AddTask, Body: task}
	_, err := DoRequest[interfaces.Task, None](api.cli, api.url, params)
	return err
}

func (api *OrchestratorApi) GetTasksStatus() ([]interfaces.TaskStatus, error) {
	params := RequestParams[None]{Endpoint: api.scheme.GetTasksStatus, Body: None{}}
	resp, err := DoRequest[None, []interfaces.TaskStatus](api.cli, api.url, params)
	if err != nil {
		return []interfaces.TaskStatus{}, err
	}
	return resp, nil
}

type GetResultResponse struct {
	Number int `json:"number"`
}

func (api *OrchestratorApi) GetResult(id string) (int, error) {
	params := RequestParams[None]{Endpoint: api.scheme.GetResult, Body: None{}}
	resp, err := DoRequest[None, GetResultResponse](api.cli, api.url, params, id)
	if err != nil {
		return 0, err
	}
	return resp.Number, nil
}

func (api *OrchestratorApi) GetOperationsTimeouts() (interfaces.OperationsTimeouts, error) {
	params := RequestParams[None]{Endpoint: api.scheme.GetOperationsTimeouts, Body: None{}}
	resp, err := DoRequest[None, interfaces.OperationsTimeouts](api.cli, api.url, params)
	return resp, err
}

func (api *OrchestratorApi) GetTask() (interfaces.Task, error) {
	params := RequestParams[None]{Endpoint: api.scheme.GetTask, Body: None{}}
	resp, err := DoRequest[None, interfaces.Task](api.cli, api.url, params)
	return resp, err
}

type SetResultBody struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
}

func (api *OrchestratorApi) SetResult(id string, number int) error {
	params := RequestParams[SetResultBody]{Endpoint: api.scheme.SetResult, Body: SetResultBody{ID: id, Number: number}}
	_, err := DoRequest[SetResultBody, None](api.cli, api.url, params)
	return err
}

type RegisterBody struct {
	Url string `json:"url"`
}

func (api *OrchestratorApi) Register(url string) {
	params := RequestParams[RegisterBody]{Endpoint: api.scheme.Register, Body: RegisterBody{Url: url}}
	DoRequest[RegisterBody, None](api.cli, api.url, params)
}
