package comment

import (
	"log/slog"
	"task-tracker-service/internal/storage/db"
)

type CommentService interface {
}

type commentService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(s db.DB, logger *slog.Logger) CommentService {
	return &commentService{
		storage: s,
		logger:  logger,
	}
}
