package task

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

type TaskDB interface {
	GetTasksWithFilter(ctx context.Context, filter *entities.TaskFilter) ([]*models.TaskModel, error)
	GetTaskByID(ctx context.Context, taskID int) (*models.TaskSummaryModel, error)
	CreateTask(ctx context.Context, task *entities.Task) (*models.TaskModel, error)
	UpdateTask(ctx context.Context, taskUpdate *entities.TaskUpdate) (*models.TaskModel, error)
}

type taskStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) TaskDB {
	return &taskStorage{
		db:     db,
		logger: logger,
	}
}

func (s *taskStorage) GetTasksWithFilter(ctx context.Context, filter *entities.TaskFilter) ([]*models.TaskModel, error) {
	return nil, nil
}

func (s *taskStorage) GetTaskByID(ctx context.Context, taskID int) (*models.TaskSummaryModel, error) {
	return nil, nil
}

func (s *taskStorage) CreateTask(ctx context.Context, task *entities.Task) (*models.TaskModel, error) {
	return nil, nil
}

func (s *taskStorage) UpdateTask(ctx context.Context, taskUpdate *entities.TaskUpdate) (*models.TaskModel, error) {
	return nil, nil
}
