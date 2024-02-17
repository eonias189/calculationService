package main

import (
	"agent/agent"
	"agent/internal/api"
	"agent/internal/server"
	"fmt"
	"os"
	"strconv"
)

func getArgs() (int, int, error) {
	if len(os.Args) < 2 {
		return 0, 0, fmt.Errorf("port not defined")
	}
	if len(os.Args) < 3 {
		return 0, 0, fmt.Errorf("threads not defined")
	}
	portString := os.Args[1]
	port, err := strconv.Atoi(portString)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid port")
	}
	threadsString := os.Args[2]
	threads, err := strconv.Atoi(threadsString)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid threads")
	}
	return port, threads, nil

}

func main() {
	port, threads, err := getArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	api := api.NewOrchestratorApi("http://127.0.0.1:8081")
	agent := agent.New(api, threads)
	s := server.NewServer(agent, port)
	s.Run()
}
