package main

import (
	"sync"

	c "github.com/eonias189/calculationService/agent/internal/config"
	"github.com/eonias189/calculationService/agent/internal/server"
	"github.com/eonias189/calculationService/agent/pkg/agent"
)

func main() {
	asp, _ := c.NewApiSchemeProvider()
	agent := agent.New()
	s := server.New(agent, ":8081", asp.GetAgentScheme())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.Run()
	}()

	wg.Wait()
	// api := api.NewOrchestratorApi("http://127.0.0.1:5000", asp.GetOrchestratorScheme())
	// fmt.Println(api.AddTask(contract.Task{ID: "69", Expression: "1000 - 7"}))
}
