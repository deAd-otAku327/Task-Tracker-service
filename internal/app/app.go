package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"task-tracker-service/internal/config"
	"task-tracker-service/internal/controller"
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
	cryptor := cryptor.New()
	tokenizer := tokenizer.New(AppName, cfg.JWTKey)

	logger, err := logger.New(os.Stdout, cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	storage, err := db.New(cfg.DBConn, logger)
	if err != nil {
		return nil, err
	}

	service := service.New(storage, logger, cryptor, tokenizer)
	controller := controller.New(service, logger)

	router := mux.NewRouter()

	return &App{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: router,
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
