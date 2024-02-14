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

func (oa *OrchestratorApi) GetTask() (c.Task, c.Timeouts, error) {
	resp := c.GetTaskResp{}
	body := utils.None{}
	err := utils.DoRequest(oa.cli, oa.Url, "getTask", "GET", body, &resp)
	if err != nil {
		return c.Task{}, c.Timeouts{}, err
	}
	return *resp.GetTask(), *resp.GetTimeouts(), err
}

func (oa *OrchestratorApi) SetResult(id string, result int) error {
	resp := c.SetResultResp{}
	body := c.SetResultBody{Id: id, Result: int64(result)}
	return utils.DoRequest(oa.cli, oa.Url, "setResult", "POST", body, &resp)
}

func (os *OrchestratorApi) Register(url string) error {
	resp := c.RegisterResp{}
	body := c.RegisterBody{Url: url}
	return utils.DoRequest(os.cli, os.Url, "register", "POST", body, &resp)
}

func NewOrchestratorApi(url string) *OrchestratorApi {
	return &OrchestratorApi{cli: &http.Client{}, Url: url}
}
