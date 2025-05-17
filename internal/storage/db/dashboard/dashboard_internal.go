package dashboard

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

func selectDashboardByBoardID(ctx context.Context, tx *sql.Tx, boardID int) (*entities.Dashboard, error) {
	query, args, err := sq.Select("*").
		From(dbconsts.TableDashboards).
		Where(sq.Eq{dbconsts.ColumnDashboardID: boardID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var dashboard entities.Dashboard
	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(&dashboard.ID, &dashboard.Title, &dashboard.Description, &dashboard.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dberrors.ErrNoRowsReturned
		}
		return nil, err
	}

	return &dashboard, nil
}

func selectDashboardTasksByBoardID(ctx context.Context, tx *sql.Tx, boardID int) ([]*entities.Task, error) {
	query, args, err := sq.Select("*").
		From(dbconsts.TableTasks).
		Where(sq.Eq{dbconsts.ColumnTaskBoardID: boardID}).
		OrderBy(dbconsts.ColumnTaskUpdatedAt).
		Suffix("DESC").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
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

	return tasks, nil
}

func selectDashboardAdminsByBoardID(ctx context.Context, tx *sql.Tx, boardID int) ([]*entities.User, error) {
	joinSettings := fmt.Sprintf("%s ON (%s = %s)",
		dbconsts.TableBoardToAdmin, dbconsts.ColumnUserID, dbconsts.ColumnBoardToAdminAdminID)

	query, args, err := sq.Select(
		dbconsts.ColumnUserID, dbconsts.ColumnUserName,
		dbconsts.ColumnUserEmail).
		From(dbconsts.TableUsers).
		Join(joinSettings).
		Where(sq.Eq{dbconsts.ColumnBoardToAdminBoardID: boardID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)

	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return users, nil
}

func insertDashboard(ctx context.Context, tx *sql.Tx, createDashboard *entities.Dashboard) (*entities.Dashboard, error) {
	insertQuery, args, err := sq.Insert(dbconsts.TableDashboards).
		Columns(dbconsts.ColumnDashboardTitle, dbconsts.ColumnDashboardDescription).
		Values(createDashboard.Title, createDashboard.Description).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var dashboard entities.Dashboard
	row := tx.QueryRowContext(ctx, insertQuery, args...)
	err = row.Scan(&dashboard.ID, &dashboard.Title, &dashboard.Description, &dashboard.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &dashboard, nil
}

func insertBoardToAdminBound(ctx context.Context, tx *sql.Tx, boardID, creatorID int) error {
	insertQuery, args, err := sq.Insert(dbconsts.TableBoardToAdmin).
		Columns(dbconsts.ColumnBoardToAdminBoardID, dbconsts.ColumnBoardToAdminAdminID).
		Values(boardID, creatorID).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, insertQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func makeSubqueryForAdminRightsCheck(boardID, userID int) string {
	return fmt.Sprintf(
		"(SELECT %s FROM %s WHERE %s=%d AND %s=%d)",
		dbconsts.ColumnBoardToAdminAdminID, dbconsts.TableBoardToAdmin,
		dbconsts.ColumnBoardToAdminAdminID, userID,
		dbconsts.ColumnBoardToAdminBoardID, boardID,
	)
}

func buildUpdateDashboardQueryFields(updateDashboard *entities.DashboardUpdate) sq.UpdateBuilder {
	query := sq.Update(dbconsts.TableDashboards)

	if updateDashboard.Title != nil {
		query = query.Set(dbconsts.ColumnDashboardTitle, updateDashboard.Title)
	}
	if updateDashboard.Description != nil {
		query = query.Set(dbconsts.ColumnDashboardDescription, updateDashboard.Description)
	}

	query = query.Set(dbconsts.ColumnDashboardUpdatedAt, time.Now())

	return query
}
