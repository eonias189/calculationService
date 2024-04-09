package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/eonias189/calculationService/agent/internal/application"
	"github.com/eonias189/calculationService/agent/internal/config"
)

func SetUpLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := SetUpLogger()

	cfg, err := config.Get()
	if err != nil {
		logger.With(slog.String("while", "starting")).Error(err.Error())
		return
	}

	app := application.New(cfg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go app.Run(ctx)

	<-sigCh
	cancel()
	app.Close()
}
