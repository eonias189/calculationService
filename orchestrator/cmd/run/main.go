package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/eonias189/calculationService/orchestrator/internal/config"
	"github.com/eonias189/calculationService/orchestrator/internal/orchestrator"
	pb "github.com/eonias189/calculationService/orchestrator/internal/proto"
	"github.com/eonias189/calculationService/orchestrator/internal/server"
	"github.com/eonias189/calculationService/orchestrator/internal/service"
	"github.com/jackc/pgx"
	"google.golang.org/grpc"
)

func SetUpLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func ConnectToDb(cfg *config.Config) (*pgx.Conn, error) {
	return pgx.Connect(pgx.ConnConfig{
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Database: cfg.PostgresDB,
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassowrd,
	})
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := SetUpLogger()

	cfg, err := config.Get()
	if err != nil {
		logger.With(slog.String("while", "getting config")).Error(err.Error())
		return
	}

	conn, err := ConnectToDb(cfg)
	if err != nil {
		logger.With(slog.String("while", "connecting to Postgres")).Error(err.Error())
		return
	}

	taskService, err := service.NewTaskService(conn)
	if err != nil {
		logger.With(slog.String("while", "getting task service")).Error(err.Error())
		return
	}

	agentService, err := service.NewAgentService(conn)
	if err != nil {
		logger.With(slog.String("while", "getting task service")).Error(err.Error())
		return
	}

	_, err = service.NewTimeoutsService(conn)
	if err != nil {
		logger.With(slog.String("while", "getting task service")).Error(err.Error())
		return
	}

	listener, err := net.Listen("tcp", cfg.GRPCAddress)
	if err != nil {
		logger.With(slog.String("while", "starting")).Error(err.Error())
		return
	}

	grpcServ := grpc.NewServer()
	orchestrator := orchestrator.New(&orchestrator.OrchestratorConfig{
		Logger:       logger,
		TaskService:  taskService,
		AgentService: agentService,
	})
	serv := server.New(orchestrator)

	pb.RegisterOrchestratorServer(grpcServ, orchestrator)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	err = orchestrator.Start(ctx)
	if err != nil {
		logger.With(slog.String("while", "starting")).Error(err.Error())
		return
	}

	errChan := make(chan error)

	go func() {
		err := serv.Run(ctx, cfg.RestApiAddress)
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		err := grpcServ.Serve(listener)
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-sigCh:
		cancel()
		orchestrator.Close()

	case err = <-errChan:
		logger.With(slog.String("while", "starting")).Error(err.Error())
		cancel()
	}
}
