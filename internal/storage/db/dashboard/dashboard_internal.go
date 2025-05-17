package dashboard

import (
	"context"
	"database/sql"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/types/entities"
	"time"

	sq "github.com/Masterminds/squirrel"
)

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
