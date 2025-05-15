package task

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"task-tracker-service/internal/mappers/entitymap"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"
	"task-tracker-service/internal/storage/db/_shared/helpers"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
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
	query, args, err := sq.Select("*").
		From(dbconsts.TableTasks).
		Where(func() sq.Eq {
			if filter.CreatorID != nil {
				return sq.Eq{dbconsts.ColumnTaskAuthorID: filter.CreatorID}
			} else if filter.AssignieID != nil {
				return sq.Eq{dbconsts.ColumnTaskAssignieID: filter.AssignieID}
			}
			return sq.Eq{}
		}()).
		// Do not touch without brainstorm. Magic number - number of WHERE statements, current must be last.
		Where(fmt.Sprintf("%s = ANY($2)", dbconsts.ColumnTaskStatus)).
		OrderBy(dbconsts.ColumnTaskUpdatedAt).
		Suffix("DESC").
		PlaceholderFormat(sq.Dollar).ToSql()
	args = append(args, pq.Array(filter.Status))

	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*entities.Task, 0)

	for rows.Next() {
		task := entities.Task{}
		err := rows.Scan(
			&task.ID, &task.Title, &task.Description, &task.Status,
			&task.AuthorID, &task.AssignieID, &task.BoardID, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return entitymap.MapToTaskListModel(tasks), nil
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
	updateQuery, args, err := s.buildUpdateTaskQueryFields(taskUpdate).
		Where(sq.Eq{dbconsts.ColumnTaskID: taskUpdate.ID}).
		Where(sq.Or{
			sq.Eq{dbconsts.ColumnTaskAuthorID: taskUpdate.InitiatorID},
			sq.Eq{dbconsts.ColumnTaskAssignieID: taskUpdate.InitiatorID}}).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var task entities.Task
	row := s.db.QueryRowContext(ctx, updateQuery, args...)
	err = row.Scan(
		&task.ID, &task.Title, &task.Description, &task.Status,
		&task.AuthorID, &task.AssignieID, &task.BoardID, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dberrors.ErrNoRowsReturned
		}
		return nil, helpers.CatchPQErrors(err)
	}

	return entitymap.MapToTaskModel(&task), nil
}

func (s *taskStorage) buildUpdateTaskQueryFields(taskUpdate *entities.TaskUpdate) sq.UpdateBuilder {
	query := sq.Update(dbconsts.TableTasks)

	if taskUpdate.Title != nil {
		query = query.Set(dbconsts.ColumnTaskTitle, taskUpdate.Title)
	}
	if taskUpdate.Description != nil {
		query = query.Set(dbconsts.ColumnTaskDescription, taskUpdate.Description)
	}
	if taskUpdate.Status != nil {
		query = query.Set(dbconsts.ColumnTaskStatus, taskUpdate.Status)
	}
	if taskUpdate.AssignieID != nil {
		query = query.Set(dbconsts.ColumnTaskAssignieID, taskUpdate.AssignieID)
	}
	if taskUpdate.BoardID != nil {
		query = query.Set(dbconsts.ColumnTaskBoardID, taskUpdate.BoardID)
	}

	query = query.Set(dbconsts.ColumnTaskUpdatedAt, time.Now())

	return query
}
