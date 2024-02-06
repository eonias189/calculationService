package handleapi

import (
	"fmt"
	"net/http"

	"github.com/eonias189/calculationService/backend/config"
	types "github.com/eonias189/calculationService/backend/interfaces"
)

type OrchestratorServer struct {
	orchestrator types.IOrchestrator
	api          config.OrchestratorScheme
	Url          string
}

func (oServ *OrchestratorServer) AddTask(w http.ResponseWriter, r *http.Request) {
	resp := types.ErrorResponse{}
	defer SendResponse(&resp, w)

	body, err := GetBody[types.AddTaskBody](r)
	if err != nil {
		resp.Error = err.Error()
		return
	}

	err = oServ.orchestrator.AddTask(body.Task)
	if err != nil {
		resp.Error = err.Error()
	}
}

func (oServ *OrchestratorServer) GetOperationsTimeouts(w http.ResponseWriter, r *http.Request) {
	resp := types.GetOperationsTimeoutsResponse{}
	defer SendResponse(&resp, w)

	res, err := oServ.orchestrator.GetOperationsTimeouts()
	if err != nil {
		resp.Error = err.Error()
		return
	}
	resp.OperationsTimeouts = res
}

func (oServ *OrchestratorServer) GetResult(w http.ResponseWriter, r *http.Request) {
	resp := types.GetResultResponse{}
	defer SendResponse(&resp, w)

	params, err := ParseRestParams[types.GetResultRestParams](r, oServ.api.GetResult)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	number, err := oServ.orchestrator.GetResult(params.ID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Number = number
}

func (oServ *OrchestratorServer) GetTask(w http.ResponseWriter, r *http.Request) {
	resp := types.GetTaskResponse{}
	defer SendResponse(&resp, w)

	task, err := oServ.orchestrator.GetTask()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Task = task
}

func (oServ *OrchestratorServer) GetTasksStatus(w http.ResponseWriter, r *http.Request) {
	resp := types.GetTasksStatusResponse{}
	defer SendResponse(&resp, w)

	res, err := oServ.orchestrator.GetTasksStatus()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.TasksStatus = res
}

func (oServ *OrchestratorServer) Register(w http.ResponseWriter, r *http.Request) {
	resp := types.ErrorResponse{}
	defer SendResponse(&resp, w)

	body, err := GetBody[types.RegisterBody](r)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	err = oServ.orchestrator.Register(body.Url)
	if err != nil {
		resp.Error = err.Error()
	}

}

func (oServ *OrchestratorServer) SetOperationsTimeouts(w http.ResponseWriter, r *http.Request) {
	resp := types.ErrorResponse{}
	defer SendResponse(&resp, w)

	body, err := GetBody[types.SetOperationsTimeoutsBody](r)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	err = oServ.orchestrator.SetOperationsTimeouts(body.OperationsTimeouts)
	if err != nil {
		resp.Error = err.Error()
	}
}

func (oServ *OrchestratorServer) SetResult(w http.ResponseWriter, r *http.Request) {
	resp := types.ErrorResponse{}
	defer SendResponse(&resp, w)

	body, err := GetBody[types.SetResultBody](r)
	if err != nil {
		resp.Error = err.Error()
		return
	}

	err = oServ.orchestrator.SetResult(body.ID, body.Number)
	if err != nil {
		resp.Error = err.Error()
	}
}

func (oServ *OrchestratorServer) Run() {
	oServ.orchestrator.Run(oServ.Url)
	mux := http.NewServeMux()
	handlers := map[string]http.HandlerFunc{
		oServ.api.AddTask.Url:               oServ.AddTask,
		oServ.api.GetOperationsTimeouts.Url: oServ.GetOperationsTimeouts,
		oServ.api.SetOperationsTimeouts.Url: oServ.SetOperationsTimeouts,
		oServ.api.GetResult.Url:             oServ.GetResult,
		oServ.api.GetTask.Url:               oServ.GetTask,
		oServ.api.GetTasksStatus.Url:        oServ.GetTasksStatus,
		oServ.api.Register.Url:              oServ.Register,
		oServ.api.SetResult.Url:             oServ.SetResult,
	}
	for endpoint, handler := range handlers {
		pattern := "/" + endpoint + "/"
		fmt.Printf("handling %v at %v\n", handler, pattern)
		mux.HandleFunc(pattern, handler)
	}
	http.ListenAndServe(oServ.Url, LogErrorMiddleware(mux))
}

func NewOrchestratorServer(orchestrator types.IOrchestrator, url string, scheme config.OrchestratorScheme) *OrchestratorServer {
	return &OrchestratorServer{orchestrator: orchestrator, Url: url, api: scheme}
}
