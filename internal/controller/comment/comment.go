package comment

import (
	"log/slog"
	"net/http"
	"task-tracker-service/internal/service"
)

type CommentHandler interface {
	Comment() http.HandlerFunc
}

type commentHandler struct {
	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger) CommentHandler {
	return &commentHandler{
		service: s,
		logger:  logger,
	}
}

func (h *commentHandler) Comment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
