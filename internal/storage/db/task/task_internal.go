package task

import (
	"context"
	"database/sql"
	"fmt"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"
	"task-tracker-service/internal/types/entities"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func selectTaskWithAuthor(ctx context.Context, tx *sql.Tx, taskID int) (*entities.Task, *entities.User, error) {
	joinSettings := fmt.Sprintf("%s ON (%s.%s = %s.%s)",
		dbconsts.TableUsers, dbconsts.TableTasks, dbconsts.ColumnTaskAuthorID, dbconsts.TableUsers, dbconsts.ColumnUserID)

	query, args, err := sq.Select("*").
		From(dbconsts.TableTasks).
		Join(joinSettings).
		Where(sq.Eq{fmt.Sprintf("%s.%s", dbconsts.TableTasks, dbconsts.ColumnTaskID): taskID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, nil, err
	}

	task, author := entities.Task{}, entities.User{}

	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&task.ID, &task.Title, &task.Description, &task.Status,
		&task.AuthorID, &task.AssignieID, &task.BoardID, &task.UpdatedAt,
		&author.ID, &author.Username, &author.Email, &author.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, dberrors.ErrNoRowsReturned
		}
		return nil, nil, err
	}

	return &task, &author, nil
}

func selectTaskCommentsByTaskID(ctx context.Context, tx *sql.Tx, taskID int) ([]*entities.Comment, error) {
	query, args, err := sq.Select("*").
		From(dbconsts.TableComments).
		Where(sq.Eq{dbconsts.ColumnCommentTaskID: taskID}).
		OrderBy(dbconsts.ColumnCommentDateTime).
		Suffix("ASC").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]*entities.Comment, 0)

	for rows.Next() {
		comment := entities.Comment{}
		err := rows.Scan(&comment.ID, &comment.TaskID, &comment.AuthorID, &comment.Text, &comment.DateTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if rows.Err() != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, nil
	}

	return comments, nil
}

func selectTaskAssignieByID(ctx context.Context, tx *sql.Tx, userID int) (*entities.User, error) {
	query, args, err := sq.Select(
		dbconsts.ColumnUserID, dbconsts.ColumnUserName, dbconsts.ColumnUserEmail).
		From(dbconsts.TableUsers).
		Where(sq.Eq{dbconsts.ColumnUserID: userID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, query, args...)

	user := entities.User{}

	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func selectTaskBoardByID(ctx context.Context, tx *sql.Tx, boardID int) (*entities.Dashboard, error) {
	query, args, err := sq.Select("*").
		From(dbconsts.TableDashboards).
		Where(sq.Eq{dbconsts.ColumnDashboardID: boardID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, query, args...)

	dashboard := entities.Dashboard{}

	err = row.Scan(&dashboard.ID, &dashboard.Title, &dashboard.Description, &dashboard.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &dashboard, nil
}

func buildUpdateTaskQueryFields(taskUpdate *entities.TaskUpdate) sq.UpdateBuilder {
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
