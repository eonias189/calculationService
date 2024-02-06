package useapi

import (
	"net/http"

	"github.com/eonias189/calculationService/backend/config"
	types "github.com/eonias189/calculationService/backend/interfaces"
)

type OrchestratorApi struct {
	url    string
	scheme config.OrchestratorScheme
	cli    *http.Client
}

func NewOrchestratorApi(scheme config.OrchestratorScheme, url string) *OrchestratorApi {
	return &OrchestratorApi{scheme: scheme, cli: &http.Client{}, url: url}
}

func (api *OrchestratorApi) AddTask(task types.Task) error {
	resp := &types.ErrorResponse{}
	params := types.AddTaskApi{
		Body: types.AddTaskBody{
			Task: task,
		},
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.AddTask, params)
	if err != nil {
		return err
	}
	return resp.GetError()
}
func (api *OrchestratorApi) GetResult(id string) (int, error) {
	resp := &types.GetResultResponse{}
	params := types.GetResultApi{
		RestParams: types.GetResultRestParams{
			ID: id,
		},
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.GetResult, params)
	if err != nil {
		return 0, err
	}
	return resp.Number, resp.GetError()
}

func (api *OrchestratorApi) GetTasksStatus() ([]types.TaskStatus, error) {
	resp := &types.GetTasksStatusResponse{}
	params := types.GetTasksStatusApi{
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.GetTasksStatus, params)
	if err != nil {
		return []types.TaskStatus{}, err
	}
	return resp.TasksStatus, resp.GetError()
}

func (api *OrchestratorApi) GetOperationsTimeouts() (types.OperationsTimeouts, error) {
	resp := &types.GetOperationsTimeoutsResponse{}
	params := types.GetOperationsTimeoutsApi{
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.GetOperationsTimeouts, params)
	if err != nil {
		return types.OperationsTimeouts{}, err
	}
	return resp.OperationsTimeouts, resp.GetError()
}

func (api *OrchestratorApi) SetOperationsTimeouts(timeouts types.OperationsTimeouts) error {
	resp := &types.ErrorResponse{}
	params := types.SetOperationsTimeoutsApi{
		Body: types.SetOperationsTimeoutsBody{
			OperationsTimeouts: timeouts,
		},
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.SetOperationsTimeouts, params)
	if err != nil {
		return err
	}
	return resp.GetError()
}

func (api *OrchestratorApi) GetTask() (types.Task, error) {
	resp := &types.GetTaskResponse{}
	params := types.GetTaskApi{
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.GetTask, params)
	if err != nil {
		return types.Task{}, err
	}
	return resp.Task, resp.GetError()
}

func (api *OrchestratorApi) SetResult(id string, number int) error {
	resp := &types.ErrorResponse{}
	params := types.SetResultApi{
		Body: types.SetResultBody{
			ID:     id,
			Number: number,
		},
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.SetResult, params)
	if err != nil {
		return err
	}
	return resp.GetError()
}

func (api *OrchestratorApi) Register(url string) error {
	resp := &types.ErrorResponse{}
	params := types.RegisterApi{
		Body: types.RegisterBody{
			Url: url,
		},
		Response: resp,
	}
	err := DoRequest(api.cli, api.url, api.scheme.Register, params)
	if err != nil {
		return err
	}
	return resp.GetError()
}
