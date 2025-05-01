package task

import (
	"log/slog"
	"task-tracker-service/internal/storage/db"
)

type TaskService interface {
}

type taskService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(s db.DB, logger *slog.Logger) TaskService {
	return &taskService{
		storage: s,
		logger:  logger,
	}
}
