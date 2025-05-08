package comment

import (
	"context"
	"log/slog"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
)

type CommentService interface {
	CreateComment(ctx context.Context, request *models.CommentCreateModel) (*dto.CommentResponse, *dto.ErrorResponse)
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

func (s *commentService) CreateComment(ctx context.Context, request *models.CommentCreateModel) (*dto.CommentResponse, *dto.ErrorResponse) {
	return nil, nil
}
