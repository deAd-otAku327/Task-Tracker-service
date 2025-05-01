package task

import (
	"log/slog"
	"task-tracker-service/internal/service"
)

type TaskHandler interface {
}

type taskHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) TaskHandler {
	return &taskHandler{
		service: s,
		logger:  logger,
	}
}
