package comment

import (
	"context"
	"log/slog"
	"net/http"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/mappers/modelmap"
	"task-tracker-service/internal/middleware"
	"task-tracker-service/internal/service/_shared/serverrors"
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
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.AuthorID = currUserID

	response, dberror := s.storage.CreateComment(ctx, modelmap.MapToComment(request))
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToCommentResponse(response), nil
}
