package main

import (
	"sync"

	c "github.com/eonias189/calculationService/orchestrator/internal/config"
	"github.com/eonias189/calculationService/orchestrator/internal/server"
	"github.com/eonias189/calculationService/orchestrator/pkg/orchestrator"
)

func main() {
	asp, _ := c.NewApiSchemeProvider()
	orch := orchestrator.New()
	s := server.New(orch, ":8081", asp.GetOrchestratorScheme())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.Run()
	}()

	wg.Wait()
}
