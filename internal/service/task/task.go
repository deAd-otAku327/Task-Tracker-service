package task

import (
	"context"
	"log/slog"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
)

type TaskService interface {
	GetTasks(ctx context.Context, request *models.TaskFilterModel) (*dto.GetTasksResponse, *dto.ErrorResponse)
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

func (s *taskService) GetTasks(ctx context.Context, request *models.TaskFilterModel) (*dto.GetTasksResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *taskService) GetTaskByID(ctx context.Context, request *models.TaskIDParamModel) (*dto.GetTaskByIDResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *taskService) CreateTask(ctx context.Context, request *models.TaskCreateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *taskService) UpdateTask(ctx context.Context, request *models.TaskUpdateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	return nil, nil
}
