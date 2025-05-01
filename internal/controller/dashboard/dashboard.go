package dashboard

import (
	"log/slog"
	"net/http"
	"task-tracker-service/internal/service"
)

type DashboardHandler interface {
	GetDashboards() http.HandlerFunc
	GetDashboardByID() http.HandlerFunc
	CreateDashboard() http.HandlerFunc
	UpdateDashboard() http.HandlerFunc
	DeleteDashboard() http.HandlerFunc
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

func (h *dashboardHandler) GetDashboards() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *dashboardHandler) GetDashboardByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *dashboardHandler) CreateDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *dashboardHandler) UpdateDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *dashboardHandler) DeleteDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
