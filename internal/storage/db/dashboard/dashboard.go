package dashboard

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
)

type DashboardDB interface {
	GetDashboardsForAdminID(ctx context.Context, adminID int) (models.DashboardListModel, error)
	GetDashboardSummaryByID(ctx context.Context, boardID int) (*models.DashboardSummaryModel, error)
	CreateDashboard(ctx context.Context, createDashboard *entities.Dashboard) (*models.DashboardModel, error)
	UpdateDashboard(ctx context.Context, updateDashboard *entities.DashboardUpdate) error
	DeleteDashboard(ctx context.Context, deleteDashboard *entities.DashboardDelete) error
	AddBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error
	DeleteBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error
}

type dashboardStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) DashboardDB {
	return &dashboardStorage{
		db:     db,
		logger: logger,
	}
}

func (s *dashboardStorage) GetDashboardsForAdminID(ctx context.Context, adminID int) (models.DashboardListModel, error) {
	joinSettings := fmt.Sprintf("%s ON (%s = %s)",
		dbconsts.TableBoardToAdmin, dbconsts.ColumnDashboardID, dbconsts.ColumnBoardToAdminBoardID)

	query, args, err := sq.Select(
		dbconsts.ColumnDashboardID, dbconsts.ColumnDashboardTitle,
		dbconsts.ColumnDashboardDescription, dbconsts.ColumnDashboardUpdatedAt,
	).
		From(dbconsts.TableDashboards).
		Join(joinSettings).
		Where(sq.Eq{dbconsts.ColumnBoardToAdminAdminID: adminID}).
		OrderBy(dbconsts.ColumnDashboardUpdatedAt).
		Suffix("DESC").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dashboards := make([]*entities.Dashboard, 0)

	for rows.Next() {
		dashboard := entities.Dashboard{}
		err := rows.Scan(
			&dashboard.ID, &dashboard.Title, &dashboard.Description, &dashboard.UpdatedAt)
		if err != nil {
			return nil, err
		}
		dashboards = append(dashboards, &dashboard)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return entitymap.MapToDashboardListModel(dashboards), nil
}

func (s *dashboardStorage) GetDashboardSummaryByID(ctx context.Context, boardID int) (*models.DashboardSummaryModel, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	dashboard, err := selectDashboardByBoardID(ctx, tx, boardID)
	if err != nil {
		helpers.RollbackTransaction(s.logger, tx)
		return nil, err
	}

	tasks, err := selectDashboardTasksByBoardID(ctx, tx, boardID)
	if err != nil {
		helpers.RollbackTransaction(s.logger, tx)
		return nil, err
	}

	admins, err := selectDashboardAdminsByBoardID(ctx, tx, boardID)
	if err != nil {
		helpers.RollbackTransaction(s.logger, tx)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return entitymap.MapToDashboardSummaryModel(dashboard, tasks, admins), nil
}

func (s *dashboardStorage) CreateDashboard(ctx context.Context, createDashboard *entities.Dashboard) (*models.DashboardModel, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	dashboard, err := insertDashboard(ctx, tx, createDashboard)
	if err != nil {
		return nil, err
	}

	err = insertBoardToAdminBound(ctx, tx, dashboard.ID, createDashboard.CreatorID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return entitymap.MapToDashboardModel(dashboard), nil
}

func (s *dashboardStorage) UpdateDashboard(ctx context.Context, updateDashboard *entities.DashboardUpdate) error {
	subq := makeSubqueryForAdminRightsCheck(updateDashboard.ID, updateDashboard.InitiatorID)

	updateQuery, args, err := buildUpdateDashboardQueryFields(updateDashboard).
		Where(sq.Eq{dbconsts.ColumnDashboardID: updateDashboard.ID}).
		Where(fmt.Sprintf("%s %s", "EXISTS", subq)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := s.db.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return err
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

func (s *dashboardStorage) DeleteDashboard(ctx context.Context, deleteDashboard *entities.DashboardDelete) error {
	subq := makeSubqueryForAdminRightsCheck(deleteDashboard.BoardID, deleteDashboard.InitiatorID)

	deleteQuery, args, err := sq.Delete(dbconsts.TableDashboards).
		Where(sq.Eq{dbconsts.ColumnDashboardID: deleteDashboard.BoardID}).
		Where(fmt.Sprintf("%s %s", "EXISTS", subq)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := s.db.ExecContext(ctx, deleteQuery, args...)
	if err != nil {
		return err
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

func (s *dashboardStorage) AddBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error {
	subq := makeSubqueryForAdminRightsCheck(boardAdminAction.BoardID, boardAdminAction.InitiatorID)
	selectForInsert := sq.Select(fmt.Sprintf("%d, %d WHERE EXISTS %s", boardAdminAction.BoardID, boardAdminAction.UserID, subq))

	insertQuery, args, err := sq.Insert(dbconsts.TableBoardToAdmin).
		Columns(dbconsts.ColumnBoardToAdminBoardID, dbconsts.ColumnBoardToAdminAdminID).
		Select(selectForInsert).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := s.db.ExecContext(ctx, insertQuery, args...)
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

func (s *dashboardStorage) DeleteBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error {
	subq := makeSubqueryForAdminRightsCheck(boardAdminAction.BoardID, boardAdminAction.InitiatorID)

	deleteQuery, args, err := sq.Delete(dbconsts.TableBoardToAdmin).
		Where(sq.Eq{dbconsts.ColumnBoardToAdminBoardID: boardAdminAction.BoardID,
			dbconsts.ColumnBoardToAdminAdminID: boardAdminAction.UserID}).
		Where(fmt.Sprintf("%s %s", "EXISTS", subq)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	result, err := s.db.ExecContext(ctx, deleteQuery, args...)
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
