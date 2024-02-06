package main

import (
	"fmt"

	"github.com/eonias189/calculationService/backend/config"
	types "github.com/eonias189/calculationService/backend/interfaces"
	"github.com/eonias189/calculationService/backend/useapi"
)

// Testing useapi
func main() {
	asp := config.NewApiSchemeProvider()
	orch := useapi.NewOrchestratorApi(asp.GetOrchestratorScheme(), "http://127.0.0.1:5000")
	ag := useapi.NewAgentApi(asp.GetAgentScheme(), "http://127.0.0.1:5000")
	fmt.Println(orch.AddTask(types.Task{ID: "some id"}))
	fmt.Println(orch.GetOperationsTimeouts())
	fmt.Println(orch.GetResult("5"))
	fmt.Println(orch.GetTask())
	fmt.Println(orch.GetTasksStatus())
	fmt.Println(orch.Register("lala"))
	fmt.Println(orch.SetResult("5", 5))
	fmt.Println(ag.GetTaskStatus("5"))
}
