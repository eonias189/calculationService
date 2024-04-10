package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/eonias189/calculationService/orchestrator/internal/application"
	"github.com/eonias189/calculationService/orchestrator/internal/config"
	"github.com/eonias189/calculationService/orchestrator/internal/lib/utils"
	"github.com/eonias189/calculationService/orchestrator/internal/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	cfg, err := config.Get()
	if err != nil {
		logger.Error(err)
		return
	}

	app, err := application.New(*cfg)
	if err != nil {
		logger.Error(err)
		return
	}

	select {
	case err := <-utils.Await(func() error {
		return app.Run(ctx)
	}):
		logger.Error(err)
		cancel()
	case <-sigCh:
		cancel()
		app.Close()
	}
}
