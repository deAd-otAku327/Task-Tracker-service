package dashboard

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"task-tracker-service/internal/mappers/entitymap"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"

	sq "github.com/Masterminds/squirrel"
)

type DashboardDB interface {
	GetDashboardsForAdminID(ctx context.Context, userID int) (models.DashboardListModel, error)
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

func (s *dashboardStorage) GetDashboardsForAdminID(ctx context.Context, userID int) (models.DashboardListModel, error) {
	return nil, nil
}

func (s *dashboardStorage) GetDashboardSummaryByID(ctx context.Context, boardID int) (*models.DashboardSummaryModel, error) {
	return nil, nil
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
	subq := fmt.Sprintf(
		"(SELECT %s FROM %s WHERE %s=%d AND %s=%d)",
		dbconsts.ColumnBoardToAdminAdminID, dbconsts.TableBoardToAdmin,
		dbconsts.ColumnBoardToAdminAdminID, updateDashboard.InitiatorID,
		dbconsts.ColumnBoardToAdminBoardID, updateDashboard.ID,
	)

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
	subq := fmt.Sprintf(
		"(SELECT %s FROM %s WHERE %s=%d AND %s=%d)",
		dbconsts.ColumnBoardToAdminAdminID, dbconsts.TableBoardToAdmin,
		dbconsts.ColumnBoardToAdminAdminID, deleteDashboard.InitiatorID,
		dbconsts.ColumnBoardToAdminBoardID, deleteDashboard.BoardID,
	)

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
	return nil
}

func (s *dashboardStorage) DeleteBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error {
	return nil
}
