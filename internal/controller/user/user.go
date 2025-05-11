package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"task-tracker-service/internal/controller/_shared/apierrors"
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
	Login() http.HandlerFunc
	GetUsers() http.HandlerFunc
}

type configParams struct {
	authExpireParam time.Duration
}

type userHandler struct {
	params configParams

	service service.Service
	logger  *slog.Logger
}

func New(s service.Service, logger *slog.Logger, authExpire time.Duration) UserHandler {
	return &userHandler{
		service: s,
		logger:  logger,
		params: configParams{
			authExpireParam: authExpire,
		},
	}
}

func (h *userHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostUsersRegisterRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestBody, http.StatusBadRequest))
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

func (h *userHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostUsersLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestBody, http.StatusBadRequest))
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
			Expires:  time.Now().Add(h.params.authExpireParam),
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
