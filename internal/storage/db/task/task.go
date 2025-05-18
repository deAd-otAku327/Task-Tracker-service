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

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type TaskDB interface {
	GetTasksWithFilter(ctx context.Context, filter *entities.TaskFilter) (models.TaskListModel, error)
	GetTaskSummaryByID(ctx context.Context, taskID int) (*models.TaskSummaryModel, error)
	CreateTask(ctx context.Context, task *entities.Task) (*models.TaskModel, error)
	UpdateTask(ctx context.Context, taskUpdate *entities.TaskUpdate) error
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

func (s *taskStorage) GetTaskSummaryByID(ctx context.Context, taskID int) (*models.TaskSummaryModel, error) {
	var task *entities.Task
	var author *entities.User
	var assignie *entities.User
	var dashboard *entities.Dashboard
	var comments []*entities.Comment

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	task, author, err = selectTaskWithAuthor(ctx, tx, taskID)
	if err != nil {
		helpers.RollbackTransaction(s.logger, tx)
		return nil, err
	}

	if task.AssignieID.Valid {
		assignie, err = selectTaskAssignieByID(ctx, tx, int(task.AssignieID.Int32))
		if err != nil {
			helpers.RollbackTransaction(s.logger, tx)
			return nil, err
		}
	}

	if task.BoardID.Valid {
		dashboard, err = selectTaskBoardByID(ctx, tx, int(task.BoardID.Int32))
		if err != nil {
			helpers.RollbackTransaction(s.logger, tx)
			return nil, err
		}
	}

	comments, err = selectTaskCommentsByTaskID(ctx, tx, taskID)
	if err != nil {
		helpers.RollbackTransaction(s.logger, tx)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return entitymap.MapToTaskSummaryModel(task, comments, author, assignie, dashboard), nil
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

func (s *taskStorage) UpdateTask(ctx context.Context, taskUpdate *entities.TaskUpdate) error {
	updateQuery, args, err := buildUpdateTaskQueryFields(taskUpdate).
		Where(sq.Eq{dbconsts.ColumnTaskID: taskUpdate.ID}).
		Where(sq.Or{
			sq.Eq{dbconsts.ColumnTaskAuthorID: taskUpdate.InitiatorID},
			sq.Eq{dbconsts.ColumnTaskAssignieID: taskUpdate.InitiatorID}}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := s.db.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return helpers.CatchPQErrors(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return dberrors.ErrNoRowsAffected
	}

	return nil
}
