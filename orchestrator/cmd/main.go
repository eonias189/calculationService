package main

import (
	"fmt"
	"orchestrator/internal/api"
	"orchestrator/internal/db"
	"orchestrator/internal/server"
	"orchestrator/orchestrator"
)

func main() {
	api := api.NewAgentApi()
	db, err := db.NewDB("db.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	orchestrator := orchestrator.NewOrchestrator(api, db)
	server := server.NewServer(orchestrator, ":8081")
	server.Run()
}
