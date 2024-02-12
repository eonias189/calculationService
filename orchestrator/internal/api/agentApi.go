package api

import (
	c "backend/contract"
	"backend/utils"
	"net/http"
)

type AgentApi struct {
	cli *http.Client
}

func (aa *AgentApi) GetStatus(url string) (c.AgentStatus, error) {
	resp := c.GetAgentStatusResp{}
	body := utils.None{}
	err := utils.DoRequest(aa.cli, url, "status", "GET", body, &resp)
	if err != nil {
		return c.AgentStatus{}, err
	}
	return *resp.Status, nil
}

func NewAgentApi() *AgentApi {
	return &AgentApi{cli: &http.Client{}}
}
