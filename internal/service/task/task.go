package task

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
	"task-tracker-service/internal/types/enum"
	"task-tracker-service/internal/types/models"
)

type TaskService interface {
	GetTasks(ctx context.Context, request *models.TaskFilterModel) (dto.GetTasksResponse, *dto.ErrorResponse)
	GetTaskByID(ctx context.Context, request *models.TaskIDParamModel) (*dto.GetTaskByIDResponse, *dto.ErrorResponse)
	CreateTask(ctx context.Context, request *models.TaskCreateModel) (*dto.TaskResponse, *dto.ErrorResponse)
	UpdateTask(ctx context.Context, request *models.TaskUpdateModel) (*dto.TaskResponse, *dto.ErrorResponse)
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

func (s *taskService) GetTasks(ctx context.Context, request *models.TaskFilterModel) (dto.GetTasksResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	if request.Relation == enum.CreatedByMe.String() {
		request.CreatorID = &currUserID
	} else if request.Relation == enum.AssignedToMe.String() {
		request.AssignieID = &currUserID
	}

	response, dberror := s.storage.GetTasksWithFilter(ctx, modelmap.MapToTaskFilter(request))
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetTaskResponse(response), nil
}

func (s *taskService) GetTaskByID(ctx context.Context, request *models.TaskIDParamModel) (*dto.GetTaskByIDResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	response, dberror := s.storage.GetTaskByID(ctx, request.TaskID)
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetTaskByIDResponse(response), nil
}

func (s *taskService) CreateTask(ctx context.Context, request *models.TaskCreateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.AuthorID = currUserID

	response, dberror := s.storage.CreateTask(ctx, modelmap.MapToTask(request))
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToTaskResponse(response), nil
}

func (s *taskService) UpdateTask(ctx context.Context, request *models.TaskUpdateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.InitiatorID = currUserID

	response, dberror := s.storage.UpdateTask(ctx, modelmap.MapToTaskUpdate(request))
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToTaskResponse(response), nil
}
