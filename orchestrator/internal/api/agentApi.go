package api

import (
	"net/http"

	c "github.com/eonias189/calculationService/orchestrator/contract"
	"github.com/eonias189/calculationService/orchestrator/internal/config"
	"github.com/eonias189/calculationService/orchestrator/internal/utils"
)

type AgentApi struct {
	scheme config.AgentScheme
	cli    *http.Client
	Url    string
}

func (a *AgentApi) GetStatus() (c.GetAgentStatusResponse, error) {
	resp := &c.GetAgentStatusResponse{}
	body, respParams := utils.None{}, utils.None{}
	params := utils.NewRequestParams(body, respParams, resp)
	err := utils.DoRequest(a.cli, a.Url, a.scheme.GetAgentStatus, params)
	return *resp, err
}

func NewAgentApi(url string, sheme config.AgentScheme) *AgentApi {
	return &AgentApi{Url: url, scheme: sheme, cli: &http.Client{}}
}
