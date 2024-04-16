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
	use_auth "github.com/eonias189/calculationService/backend/internal/use_cases/auth"
	use_tasks "github.com/eonias189/calculationService/backend/internal/use_cases/tasks"
	use_timeouts "github.com/eonias189/calculationService/backend/internal/use_cases/timeouts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Application struct {
	cfg             api_config.Config
	cli             pb.OrchestratorClient
	userService     *service.UserService
	tasksService    *service.TaskService
	agentService    *service.AgentService
	timeoutsService *service.TimeoutsSerice
	server          *http.Server
	r               chi.Router
}

type Distributer struct {
	cli pb.OrchestratorClient
}

func (d *Distributer) Distribute(task *pb.Task) error {
	_, err := d.cli.Distribute(context.TODO(), task)
	return err
}

func (a *Application) MountHandlers() {
	const (
		secretKey     = "very very secret"
		signingMethod = "HS256"
		expTime       = time.Hour * 24 * 30
	)
	tokenAuth := jwtauth.New(signingMethod, []byte(secretKey), nil, jwt.WithRequiredClaim("user_id"))

	a.r.Use(middleware.Logger)
	a.r.Use(middleware.Recoverer)

	api := chi.NewRouter()

	api.Use(cors.AllowAll().Handler)
	api.Use(render.SetContentType(render.ContentTypeJSON))

	a.r.Mount("/api", api)

	api.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, render.M{"message": "pong"})
	})

	api.Mount("/auth", use_auth.MakeHandler(a.userService, tokenAuth, expTime))
	api.Mount("/tasks", use_tasks.MakeHandler(a.tasksService, a.timeoutsService, &Distributer{cli: a.cli}, tokenAuth))
	api.Mount("/timeouts", use_timeouts.MakeHandler(a.timeoutsService, tokenAuth))

	api.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Get("/inf", func(w http.ResponseWriter, r *http.Request) {
			_, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				utils.HandleError(w, r, err, http.StatusUnauthorized)
				return
			}
			render.JSON(w, r, claims)
		})
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
	logger.Info("starting at", a.cfg.Address)
	a.MountHandlers()
	return a.server.ListenAndServe()
}

func (a *Application) init() error {
	pool, err := pgxpool.New(context.TODO(), a.cfg.PostgresConn)
	if err != nil {
		return err
	}

	utils.TryUntilSuccess(context.TODO(), func() error { return pool.Ping(context.TODO()) }, time.Second*5)
	us, err := service.NewUserService(pool)
	if err != nil {
		return err
	}

	ts, err := service.NewTaskService(pool)
	if err != nil {
		return err
	}

	as, err := service.NewAgentService(pool)
	if err != nil {
		return err
	}

	tmts, err := service.NewTimeoutsService(pool)
	if err != nil {
		return err
	}

	a.userService = us
	a.tasksService = ts
	a.agentService = as
	a.timeoutsService = tmts
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
