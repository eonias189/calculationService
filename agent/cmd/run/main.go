package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eonias189/calculationService/agent/internal/agent"
	"github.com/eonias189/calculationService/agent/internal/config"
	"github.com/eonias189/calculationService/agent/internal/lib/pool"
	pb "github.com/eonias189/calculationService/agent/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetUpLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func GetCli(addr string, l *slog.Logger) pb.OrchestratorClient {
	var (
		err  error
		conn *grpc.ClientConn
	)
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	for err != nil {
		l.With(slog.String("while", "connecting")).Error(err.Error())
		time.Sleep(time.Second * 3)
		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return pb.NewOrchestratorClient(conn)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := SetUpLogger()

	cfg, err := config.Get()
	if err != nil {
		logger.With(slog.String("while", "starting")).Error(err.Error())
		return
	}

	cli := GetCli(cfg.OrchestratorAddr, logger)
	workerPool := pool.NewWorkerPool(cfg.MaxThreads)

	agent := agent.New(&agent.AgentOptions{
		Api:        cli,
		WorkerPool: workerPool,
		Logger:     logger,
		MaxThreads: cfg.MaxThreads,
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	err = agent.Run(ctx)
	if err != nil {
		logger.With(slog.String("while", "starting")).Error(err.Error())
		return
	}

	<-sigCh
	cancel()
	agent.Close()
}
