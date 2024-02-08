package api

import (
	"net/http"

	c "github.com/eonias189/calculationService/agent/contract"
	"github.com/eonias189/calculationService/agent/internal/config"
	"github.com/eonias189/calculationService/agent/internal/utils"
)

type OrchestratorApi struct {
	scheme config.OrchestratorScheme
	cli    *http.Client
	Url    string
}

func (o *OrchestratorApi) GetTask() (c.Task, error) {
	resp := &c.GetTaskResponse{}
	body := utils.None{}
	respParams := utils.None{}
	params := utils.NewRequestParams(body, respParams, resp)
	err := utils.DoRequest(o.cli, o.Url, o.scheme.GetTask, params)
	return resp.Task, err
}

func (o *OrchestratorApi) SetResult(id string, res int) error {
	resp := &c.ErrorResponse{}
	body := c.SetResultBody{ID: id, Number: res}
	restParams := utils.None{}
	params := utils.NewRequestParams(body, restParams, resp)
	return utils.DoRequest(o.cli, o.Url, o.scheme.SetResult, params)
}
func (o *OrchestratorApi) Register(url string) error {
	resp := &c.ErrorResponse{}
	body := c.RegisterBody{Url: url}
	restParams := utils.None{}
	params := utils.NewRequestParams(body, restParams, resp)
	return utils.DoRequest(o.cli, o.Url, o.scheme.Register, params)
}

func NewOrchestratorApi(url string, sheme config.OrchestratorScheme) *OrchestratorApi {
	return &OrchestratorApi{Url: url, scheme: sheme, cli: &http.Client{}}
}
