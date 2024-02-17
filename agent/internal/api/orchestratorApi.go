package api

import (
	c "backend/contract"
	"backend/utils"
	"net/http"
)

type OrchestratorApi struct {
	cli *http.Client
	Url string
}

func (oa *OrchestratorApi) GetTask(id int) (*c.Task, *c.Timeouts, error) {
	resp := c.GetTaskResp{}
	body := c.GetTaskBody{AgentId: int64(id)}
	err := utils.DoRequest(oa.cli, oa.Url, "getTask", "POST", &body, &resp)
	if err != nil {
		return &c.Task{}, &c.Timeouts{}, err
	}
	return resp.GetTask(), resp.GetTimeouts(), err
}

func (oa *OrchestratorApi) SetResult(id string, result int, status c.TaskStatus) error {
	resp := c.SetResultResp{}
	body := c.SetResultBody{Id: id, Result: int64(result), Status: status}
	return utils.DoRequest(oa.cli, oa.Url, "setResult", "POST", &body, &resp)
}

func (os *OrchestratorApi) Register(url string) (int, error) {
	resp := c.RegisterResp{}
	body := c.RegisterBody{Url: url}
	err := utils.DoRequest(os.cli, os.Url, "register", "POST", &body, &resp)
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func NewOrchestratorApi(url string) *OrchestratorApi {
	return &OrchestratorApi{cli: &http.Client{}, Url: url}
}
