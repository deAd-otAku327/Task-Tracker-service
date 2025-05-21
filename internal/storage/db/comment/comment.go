package comment

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

type CommentDB interface {
	CreateComment(ctx context.Context, comment *entities.Comment) (*models.CommentModel, error)
}

type commentStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) CommentDB {
	return &commentStorage{
		db:     db,
		logger: logger,
	}
}

func (s *commentStorage) CreateComment(ctx context.Context, comment *entities.Comment) (*models.CommentModel, error) {
	return nil, nil
}
