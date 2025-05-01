package user

import (
	"log/slog"
	"task-tracker-service/internal/service"
)

type UserHandler interface {
}

type userHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) UserHandler {
	return &userHandler{
		service: s,
		logger:  logger,
	}
}
