package main

import (
	"orchestrator/internal/api"
	"orchestrator/internal/server"
	"orchestrator/orchestrator"
)

func main() {
	api := api.NewAgentApi()
	orchestrator := orchestrator.New(api)
	server := server.NewServer(orchestrator, ":8081")
	server.Run()
}
