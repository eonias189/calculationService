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

func (o *OrchestratorApi) AddTask(t c.Task) error {
	resp := &c.ErrorResponse{}
	body := c.AddTaskBody{Task: t}
	respParams := utils.None{}
	params := utils.NewRequestParams(body, respParams, resp)
	err := utils.DoRequest(o.cli, o.Url, o.scheme.AddTask, params)
	if err != nil {
		return err
	}
	return resp.GetError()
}

/* func (o *OrchestratorApi) GetTasksStatus() ([]c.TaskStatus, error)
func (o *OrchestratorApi) GetResult(string) (int, error)
func (o *OrchestratorApi) GetOperationsTimeouts() (c.OperationsTimeouts, error)
func (o *OrchestratorApi) SetOperationsTimeouts(c.OperationsTimeouts) error
func (o *OrchestratorApi) GetTask() (c.Task, error)
func (o *OrchestratorApi) SetResult(string, int) error
func (o *OrchestratorApi) Register(string) error */

func NewOrchestratorApi(url string, sheme config.OrchestratorScheme) *OrchestratorApi {
	return &OrchestratorApi{Url: url, scheme: sheme, cli: &http.Client{}}
}
