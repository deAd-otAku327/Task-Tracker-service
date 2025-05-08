package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"task-tracker-service/internal/controller/_shared/cerrors"
	"task-tracker-service/internal/controller/_shared/responser"
	"task-tracker-service/internal/mappers/dtomap"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/middleware"
	"task-tracker-service/internal/service"
	"task-tracker-service/internal/types/dto"
	"time"
)

type UserHandler interface {
	Register() http.HandlerFunc
	Login(authExpire time.Duration) http.HandlerFunc
	GetUsers() http.HandlerFunc
	AddBoardAdmin() http.HandlerFunc
	DeleteBoardAdmin() http.HandlerFunc
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
		request := dto.PostUsersRegisterRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.RegistrateUser(r.Context(), dtomap.MapToUserRegisterModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusCreated, &response)
	}
}

func (h *userHandler) Login(authExpire time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostUsersLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.LoginUser(r.Context(), dtomap.MapToUserLoginModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     middleware.CookieName,
			Value:    string(*response),
			Expires:  time.Now().Add(authExpire),
			SameSite: http.SameSiteStrictMode,
		})

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *userHandler) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, servErr := h.service.GetUsers(r.Context())
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *userHandler) AddBoardAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostUsersBoardAdminRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.AddBoardAdmin(r.Context(), dtomap.MapToUserBoardAdminModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *userHandler) DeleteBoardAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostUsersBoardAdminRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.DeleteBoardAdmin(r.Context(), dtomap.MapToUserBoardAdminModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}
