package main

import (
	"fmt"

	"github.com/eonias189/calculationService/orchestrator/internal/api"
	c "github.com/eonias189/calculationService/orchestrator/internal/config"
)

func main() {
	asp, _ := c.NewApiSchemeProvider()
	/* orch := orchestrator.New()
	s := server.New(orch, ":8081", asp.GetOrchestratorScheme())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.Run()
	}()

	wg.Wait() */
	api := api.NewAgentApi("http://127.0.0.1:8081", asp.GetAgentScheme())
	fmt.Println(api.GetStatus())
}
