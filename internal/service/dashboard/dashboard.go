package dashboard

import (
	"context"
	"log/slog"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
)

type DashboardService interface {
	GetDashboards(ctx context.Context) (*dto.GetDashboardsResponse, *dto.ErrorResponse)
	GetDashboardByID(ctx context.Context, request *models.DashboardIDParamModel) (*dto.GetDashboardByIDResponse, *dto.ErrorResponse)
	CreateDashboard(ctx context.Context, request *models.DashboardCreateModel) (*dto.DashboardResponse, *dto.ErrorResponse)
	UpdateDashboard(ctx context.Context, request *models.DashboardUpdateModel) (*dto.DashboardResponse, *dto.ErrorResponse)
	DeleteDashboard(ctx context.Context, request *models.DashboardDeleteModel) *dto.ErrorResponse
	AddBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse
	DeleteBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse
}

type dashboardService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(s db.DB, logger *slog.Logger) DashboardService {
	return &dashboardService{
		storage: s,
		logger:  logger,
	}
}

func (s *dashboardService) GetDashboards(ctx context.Context) (*dto.GetDashboardsResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *dashboardService) GetDashboardByID(ctx context.Context, request *models.DashboardIDParamModel,
) (*dto.GetDashboardByIDResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *dashboardService) CreateDashboard(ctx context.Context, request *models.DashboardCreateModel) (*dto.DashboardResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *dashboardService) UpdateDashboard(ctx context.Context, request *models.DashboardUpdateModel) (*dto.DashboardResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *dashboardService) DeleteDashboard(ctx context.Context, request *models.DashboardDeleteModel) *dto.ErrorResponse {
	return nil
}

func (s *dashboardService) AddBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse {
	return nil
}

func (s *dashboardService) DeleteBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse {
	return nil
}
