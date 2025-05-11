package dashboard

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

type DashboardDB interface {
	GetDashboardsForAdminID(ctx context.Context, userID int) ([]*models.DashboardModel, error)
	GetDashboardByID(ctx context.Context, boardID int) (*models.DashboardSummaryModel, error)
	CreateDashboard(ctx context.Context, dashboard *entities.Dashboard) (*models.DashboardModel, error)
	UpdateDashboard(ctx context.Context, dashboardUpdate *entities.DashboardUpdate) (*models.DashboardModel, error)
	DeleteDashboard(ctx context.Context, dashboardDelete *entities.DashboardDelete) error
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

func (s *dashboardStorage) GetDashboardsForAdminID(ctx context.Context, userID int) ([]*models.DashboardModel, error) {
	return nil, nil
}

func (s *dashboardStorage) GetDashboardByID(ctx context.Context, boardID int) (*models.DashboardSummaryModel, error) {
	return nil, nil
}

func (s *dashboardStorage) CreateDashboard(ctx context.Context, dashboard *entities.Dashboard) (*models.DashboardModel, error) {
	return nil, nil
}

func (s *dashboardStorage) UpdateDashboard(ctx context.Context, dashboardUpdate *entities.DashboardUpdate) (*models.DashboardModel, error) {
	return nil, nil
}

func (s *dashboardStorage) DeleteDashboard(ctx context.Context, dashboardDelete *entities.DashboardDelete) error {
	return nil
}

func (s *dashboardStorage) AddBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error {
	return nil
}

func (s *dashboardStorage) DeleteBoardAdmin(ctx context.Context, boardAdminAction *entities.DashboardAdminAction) error {
	return nil
}
