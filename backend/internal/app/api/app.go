package api

import (
	"context"
	"net/http"
	"time"

	"github.com/eonias189/calculationService/backend/internal/config/api_config"
	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/eonias189/calculationService/backend/internal/logger"
	pb "github.com/eonias189/calculationService/backend/internal/proto"
	"github.com/eonias189/calculationService/backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Application struct {
	cfg          api_config.Config
	cli          pb.OrchestratorClient
	tasksService *service.TaskService
	server       *http.Server
	r            chi.Router
}

func (a *Application) MountHandlers() {
	a.r.Use(middleware.Logger)
	a.r.Use(middleware.Recoverer)
	a.r.Use(cors.AllowAll().Handler)

	api := chi.NewRouter()
	api.Use(render.SetContentType(render.ContentTypeJSON))
	a.r.Mount("/api", api)

	api.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, render.M{"message": "pong"})
	})

	api.Get("/distrib", func(w http.ResponseWriter, r *http.Request) {
		t := service.Task{Expression: "2 + 2 * 2", UserId: 1, Status: service.TaskStatusPending}
		id, err := a.tasksService.Add(t)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{"reason": err.Error()})
			return
		}

		task := &pb.Task{Id: id, Expression: t.Expression, Timeouts: &pb.Timeouts{
			Add: 10,
			Sub: 2,
			Mul: 15,
			Div: 4,
		}}

		_, err = a.cli.Distribute(context.TODO(), task)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{"reason": err.Error()})
			return
		}

		render.JSON(w, r, render.M{"message": "ok"})
	})
}

func (a *Application) GetCli(ctx context.Context) pb.OrchestratorClient {
	const (
		interval = time.Second * 5
	)

	var conn *grpc.ClientConn
	utils.TryUntilSuccess(ctx, func() error {
		cliConn, err := grpc.Dial(a.cfg.OrchestratorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err == nil {
			conn = cliConn
		} else {
			logger.Error(err)
		}

		return err
	}, interval)

	return pb.NewOrchestratorClient(conn)
}
func (a *Application) Run(ctx context.Context) error {
	a.MountHandlers()
	return a.server.ListenAndServe()
}

func (a *Application) init() error {
	pool, err := pgxpool.New(context.TODO(), a.cfg.PostgresConn)
	if err != nil {
		return err
	}

	utils.TryUntilSuccess(context.TODO(), func() error { return pool.Ping(context.TODO()) }, time.Second*5)
	ts, err := service.NewTaskService(pool)
	if err != nil {
		return err
	}

	a.cli = a.GetCli(context.TODO())
	a.tasksService = ts
	return nil
}

func (a *Application) Close(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func New(cfg api_config.Config) (*Application, error) {
	r := chi.NewRouter()
	app := &Application{
		cfg:    cfg,
		r:      r,
		server: &http.Server{Addr: cfg.Address, Handler: r},
	}

	err := app.init()
	if err != nil {
		return nil, err
	}

	return app, nil
}
