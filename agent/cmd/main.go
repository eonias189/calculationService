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
	/* api := api.NewOrchestratorApi("http://127.0.0.1:8081", asp.GetOrchestratorScheme())
	fmt.Println(api.GetTask())
	fmt.Println(api.SetResult("69", 993))
	fmt.Println(api.Register("lala")) */
}
