package task

import (
	"log/slog"
	"net/http"
	"task-tracker-service/internal/service"
)

type TaskHandler interface {
	GetTasks() http.HandlerFunc
	GetTaskByID() http.HandlerFunc
	CreateTask() http.HandlerFunc
	UpdateTask() http.HandlerFunc
}

type taskHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) TaskHandler {
	return &taskHandler{
		service: s,
		logger:  logger,
	}
}

func (h *taskHandler) GetTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *taskHandler) GetTaskByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *taskHandler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *taskHandler) UpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
