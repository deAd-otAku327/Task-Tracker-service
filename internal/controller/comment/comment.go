package comment

import (
	"log/slog"
	"task-tracker-service/internal/service"
)

type CommentHandler interface {
}

type commentHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) CommentHandler {
	return &commentHandler{
		service: s,
		logger:  logger,
	}
}
