package main

import (
	"fmt"

	"github.com/eonias189/calculationService/backend/config"
	"github.com/eonias189/calculationService/backend/interfaces"
	"github.com/eonias189/calculationService/backend/useapi"
)

// Testing useapi
func main() {
	asp := config.NewApiSchemeProvider()
	orch := useapi.NewOrchestratorApi(asp.GetOrchestratorScheme(), "http://127.0.0.1:5000")
	ag := useapi.NewAgentApi(asp.GetAgentScheme(), "http://127.0.0.1:5000")
	fmt.Println("addTask", orch.AddTask(interfaces.Task{ID: "some id"}))
	fmt.Print("GetOperationsTimeouts ")
	fmt.Println(orch.GetOperationsTimeouts())
	fmt.Print("GerResult ")
	fmt.Println(orch.GetResult("5"))
	fmt.Print("Get task ")
	fmt.Println(orch.GetTask())
	fmt.Print("get tasks status")
	fmt.Println(orch.GetTasksStatus())
	fmt.Println("register")
	orch.Register("lala")
	fmt.Println("setResult", orch.SetResult("5", 5))
	fmt.Print("get task status ")
	fmt.Println(ag.GetTaskStatus("5"))
	fmt.Println("is working", ag.IsWorking())
}
