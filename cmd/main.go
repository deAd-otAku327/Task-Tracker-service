package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"task-tracker-service/internal/app"
	"task-tracker-service/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	configDir := filepath.Join("..", "configs")
	envPath := filepath.Join(configDir, ".env")

	err := godotenv.Load(envPath)
	reportOnError(err)

	configPath := filepath.Join(configDir, fmt.Sprintf("%s.yaml", os.Getenv("ENV")))

	cfg, err := config.New(configPath)
	reportOnError(err)

	app, err := app.New(cfg)
	reportOnError(err)

	go func() {
		err = app.Run()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("server error: %v", err)
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	slog.Info("starting server shutdown")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
	slog.Info("server shutdown completed")
}

func reportOnError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
