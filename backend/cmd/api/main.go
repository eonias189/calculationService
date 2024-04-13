package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eonias189/calculationService/backend/internal/app/api"
	"github.com/eonias189/calculationService/backend/internal/config/api_config"
	"github.com/eonias189/calculationService/backend/internal/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := api_config.Get()
	if err != nil {
		logger.Error(err)
		return
	}

	app, err := api.New(*cfg)
	if err != nil {
		logger.Error(err)
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigCh
		shutdownCtx, _ := context.WithTimeout(ctx, time.Second*10)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Info("graceful shutdown timed out.. forcing exit.")
			}
		}()
		err := app.Close(shutdownCtx)
		if err != nil {
			logger.Error(err)
			return
		}

		cancel()
	}()

	err = app.Run(ctx)
	if err != nil {
		logger.Error(err)
	}

	<-ctx.Done()
}
