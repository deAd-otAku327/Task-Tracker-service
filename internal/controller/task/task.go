package task

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"task-tracker-service/internal/controller/_shared/apiconsts"
	"task-tracker-service/internal/controller/_shared/apierrors"
	"task-tracker-service/internal/controller/_shared/responser"
	"task-tracker-service/internal/mappers/dtomap"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/service"
	"task-tracker-service/internal/types/dto"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type TaskHandler interface {
	GetTasks() http.HandlerFunc
	GetTaskSummary() http.HandlerFunc
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
		err := r.ParseForm()
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestParams, http.StatusBadRequest))
			return
		}

		request := dto.GetTasksParams{}
		err = schema.NewDecoder().Decode(&request, r.Form)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestParamsFormat, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.GetTasks(r.Context(), dtomap.MapToTaskFilterModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *taskHandler) GetTaskSummary() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// No check for empty map cause of path-param cannot be missing.
		param, err := strconv.Atoi(mux.Vars(r)[apiconsts.URLParamTaskID])
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestParamsFormat, http.StatusBadRequest))
			return
		}
		request := dto.GetTaskSummaryParam{
			TaskID: param,
		}

		response, servErr := h.service.GetTaskSummary(r.Context(), dtomap.MapToTaskIDParamModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}

func (h *taskHandler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostTasksCreateRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.CreateTask(r.Context(), dtomap.MapToTaskCreateModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusCreated, &response)
	}
}

func (h *taskHandler) UpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := dto.PostTasksUpdateRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrInvalidRequestBody, http.StatusBadRequest))
			return
		}

		response, servErr := h.service.UpdateTask(r.Context(), dtomap.MapToTaskUpdateModel(&request))
		if servErr != nil {
			responser.MakeErrorResponseJSON(w, servErr)
			return
		}

		responser.MakeResponseJSON(w, http.StatusOK, &response)
	}
}
