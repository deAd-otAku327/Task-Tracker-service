package comment

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"task-tracker-service/internal/controller/_shared/apierrors"
	"task-tracker-service/internal/controller/_shared/responser"
	"task-tracker-service/internal/mappers/dtomap"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/service"
	"task-tracker-service/internal/types/dto"
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
		request := dto.PostCommentRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.CreateComment(r.Context(), dtomap.MapToCommentCreateModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusCreated, &response)
	}
}
