package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"task-tracker-service/internal/config"

	"github.com/gorilla/mux"
)

const AppName = "Task-Tracker"

type App struct {
	Server *http.Server
}

func New(cfg *config.Config) (*App, error) {

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
