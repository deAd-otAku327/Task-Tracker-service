package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"task-tracker-service/internal/config"
	"task-tracker-service/internal/controller"
	"task-tracker-service/internal/middleware"
	"task-tracker-service/internal/service"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/tokenizer"
	"task-tracker-service/pkg/cryptor"
	"task-tracker-service/pkg/logger"

	"github.com/gorilla/mux"
)

const AppName = "Task-Tracker"

type App struct {
	Server *http.Server
}

func New(cfg *config.Config) (*App, error) {
	cryptor := cryptor.New(cfg.AsyncHashingLimit)
	tokenizer := tokenizer.New(AppName, cfg.JWTKey, cfg.AuthExpire)

	logger, err := logger.New(os.Stdout, cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	storage, err := db.New(cfg.DBConn, logger)
	if err != nil {
		return nil, err
	}

	service := service.New(storage, logger, cryptor, tokenizer)

	controller := controller.New(service, logger, &controller.HandlersConfig{
		AuthExpire: cfg.AuthExpire,
	})
	middleware := middleware.New(tokenizer, logger, middleware.MwParams{
		RpsLimit:      cfg.RPS,
		RespTimeLimit: cfg.ResponseTime,
	})

	return &App{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: configureRouter(controller, middleware),
		},
	}, nil
}

func (s *App) Run() error {
	slog.Info("app starting on", slog.String("address", s.Server.Addr))
	return s.Server.ListenAndServe()
}

func (s *App) Shutdown() error {
	return s.Server.Shutdown(context.Background())
}

func configureRouter(controller controller.Controller, middleware middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.RpsLimit())
	router.Use(middleware.Logging())
	router.Use(middleware.ResponseTimeLimit())

	router.HandleFunc("/register", controller.Register()).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login()).Methods(http.MethodPost)

	protected := router.NewRoute().Subrouter()
	protected.Use(middleware.Auth())

	protected.HandleFunc("/users", controller.GetUsers()).Methods(http.MethodGet)

	protected.HandleFunc("/tasks", controller.GetTasks()).Methods(http.MethodGet)
	protected.HandleFunc("/tasks/{taskId:[1-9][0-9]*}", controller.GetTaskSummary()).Methods(http.MethodGet)
	protected.HandleFunc("/tasks/create", controller.CreateTask()).Methods(http.MethodPost)
	protected.HandleFunc("/tasks/update", controller.UpdateTask()).Methods(http.MethodPost)

	protected.HandleFunc("/comment", controller.Comment()).Methods(http.MethodPost)

	protected.HandleFunc("/dashboards", controller.GetDashboards()).Methods(http.MethodGet)
	protected.HandleFunc("/dashboards/{boardId:[1-9][0-9]*}", controller.GetDashboardByID()).Methods(http.MethodGet)
	protected.HandleFunc("/dashboards/create", controller.CreateDashboard()).Methods(http.MethodPost)
	protected.HandleFunc("/dashboards/update", controller.UpdateDashboard()).Methods(http.MethodPost)
	protected.HandleFunc("/dashboards/delete", controller.DeleteDashboard()).Methods(http.MethodPost)
	protected.HandleFunc("/dashboards/addBoardAdmin", controller.AddBoardAdmin()).Methods(http.MethodPost)
	protected.HandleFunc("/dashboards/deleteBoardAdmin", controller.DeleteBoardAdmin()).Methods(http.MethodPost)

	return router
}
