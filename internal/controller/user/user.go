package user

import (
	"log/slog"
	"net/http"
	"task-tracker-service/internal/service"
)

type UserHandler interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	GetUsers() http.HandlerFunc
}

type userHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) UserHandler {
	return &userHandler{
		service: s,
		logger:  logger,
	}
}

func (h *userHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *userHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *userHandler) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
