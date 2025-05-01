package dashboard

import (
	"log/slog"
	"task-tracker-service/internal/storage/db"
)

type DashboardService interface {
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
