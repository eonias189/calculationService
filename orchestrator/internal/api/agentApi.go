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

func (a *AgentApi) Ping() (bool, error) {
	resp := &c.PingResponse{}
	body, respParams := utils.None{}, utils.None{}
	params := utils.NewRequestParams(body, respParams, resp)
	err := utils.DoRequest(a.cli, a.Url, a.scheme.Ping, params)
	if err != nil {
		return false, err
	}
	return resp.Ok, resp.GetError()
}

func NewAgentApi(url string, sheme config.AgentScheme) *AgentApi {
	return &AgentApi{Url: url, scheme: sheme, cli: &http.Client{}}
}
