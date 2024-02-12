package main

import (
	"agent/agent"
	"agent/internal/api"
	"agent/internal/server"
	"fmt"
	"os"
	"strconv"
)

func getPort() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("port not defined")
	}
	port := os.Args[1]
	if len(port) < 2 {
		return "", fmt.Errorf("invalid port")
	}
	if port[0] != ':' {
		return "", fmt.Errorf("invalid port")
	}
	_, err := strconv.Atoi(port[1:])
	if err != nil {
		return "", fmt.Errorf("invalid port")
	}
	return port, nil
}

func main() {
	port, err := getPort()
	if err != nil {
		fmt.Println(err)
		return
	}
	api := api.NewOrchestratorApi("http://127.0.0.1:8081")
	agent := agent.New(api)
	s := server.NewServer(agent, port)
	s.Run()
}
