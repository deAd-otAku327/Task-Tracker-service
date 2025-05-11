package dashboard

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"task-tracker-service/internal/controller/_shared/cerrors"
	"task-tracker-service/internal/controller/_shared/consts"
	"task-tracker-service/internal/controller/_shared/responser"
	"task-tracker-service/internal/mappers/dtomap"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/service"
	"task-tracker-service/internal/types/dto"

	"github.com/gorilla/mux"
)

type DashboardHandler interface {
	GetDashboards() http.HandlerFunc
	GetDashboardByID() http.HandlerFunc
	CreateDashboard() http.HandlerFunc
	UpdateDashboard() http.HandlerFunc
	DeleteDashboard() http.HandlerFunc
	AddBoardAdmin() http.HandlerFunc
	DeleteBoardAdmin() http.HandlerFunc
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
		response, servErr := h.service.GetDashboards(r.Context())
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *dashboardHandler) GetDashboardByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// No check for empty map cause of path-param cannot be missing.
		param, err := strconv.Atoi(mux.Vars(r)[consts.URLParamDashboardID])
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestParamsFormat, http.StatusBadRequest))
			return
		}
		request := dto.GetDashboardByIDParam{
			BoardID: param,
		}

		response, servErr := h.service.GetDashboardByID(r.Context(), dtomap.MapToDashboardIDParamModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *dashboardHandler) CreateDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostDashboardsCreateRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.CreateDashboard(r.Context(), dtomap.MapToDashboardCreateModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusCreated, &response)
	}
}

func (h *dashboardHandler) UpdateDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostDashboardsUpdateRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.UpdateDashboard(r.Context(), dtomap.MapToDashboardUpdateModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *dashboardHandler) DeleteDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostDashboardsDeleteRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		servErr := h.service.DeleteDashboard(r.Context(), dtomap.MapToDashboardDeleteModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

func (h *dashboardHandler) AddBoardAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostDashboardsAdminRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		servErr := h.service.AddBoardAdmin(r.Context(), dtomap.MapToDashboardAdminActionModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, nil)
	}
}

func (h *dashboardHandler) DeleteBoardAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostDashboardsAdminRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(cerrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		servErr := h.service.DeleteBoardAdmin(r.Context(), dtomap.MapToDashboardAdminActionModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, nil)
	}
}
