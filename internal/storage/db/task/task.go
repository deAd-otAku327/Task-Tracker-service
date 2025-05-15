package task

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/mappers/entitymap"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/helpers"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"

	sq "github.com/Masterminds/squirrel"
)

type TaskDB interface {
	GetTasksWithFilter(ctx context.Context, filter *entities.TaskFilter) (models.TaskListModel, error)
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

func (s *taskStorage) GetTasksWithFilter(ctx context.Context, filter *entities.TaskFilter) (models.TaskListModel, error) {
	return nil, nil
}

func (s *taskStorage) GetTaskByID(ctx context.Context, taskID int) (*models.TaskSummaryModel, error) {
	return nil, nil
}

func (s *taskStorage) CreateTask(ctx context.Context, createTask *entities.Task) (*models.TaskModel, error) {
	insertQuery, args, err := sq.Insert(dbconsts.TableTasks).
		Columns(
			dbconsts.ColumnTaskTitle,
			dbconsts.ColumnTaskDescription,
			dbconsts.ColumnTaskAuthorID,
			dbconsts.ColumnTaskAssignieID,
			dbconsts.ColumnTaskBoardID,
		).
		Values(
			createTask.Title,
			createTask.Description,
			createTask.AuthorID,
			createTask.AssignieID,
			createTask.BoardID,
		).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var task entities.Task
	row := s.db.QueryRowContext(ctx, insertQuery, args...)
	err = row.Scan(
		&task.ID, &task.Title, &task.Description, &task.Status,
		&task.AuthorID, &task.AssignieID, &task.BoardID, &task.UpdatedAt)
	if err != nil {
		return nil, helpers.CatchPQErrors(err)
	}

	return entitymap.MapToTaskModel(&task), nil
}

func (s *taskStorage) UpdateTask(ctx context.Context, taskUpdate *entities.TaskUpdate) (*models.TaskModel, error) {
	return nil, nil
}
