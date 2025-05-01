package dashboard

import (
	"log/slog"
	"task-tracker-service/internal/service"
)

type DashboardHandler interface {
}

type dashboardHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) DashboardHandler {
	return &dashboardHandler{
		service: s,
		logger:  logger,
	}
}
