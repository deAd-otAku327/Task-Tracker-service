package task

import (
	"context"
	"log/slog"
	"net/http"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/mappers/modelmap"
	"task-tracker-service/internal/middleware"
	"task-tracker-service/internal/service/_shared/serverrors"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/enum"
	"task-tracker-service/internal/types/models"
)

type TaskService interface {
	GetTasks(ctx context.Context, request *models.TaskFilterModel) (dto.GetTasksResponse, *dto.ErrorResponse)
	GetTaskSummary(ctx context.Context, request *models.TaskSummaryParamModel) (*dto.GetTaskSummaryResponse, *dto.ErrorResponse)
	CreateTask(ctx context.Context, request *models.TaskCreateModel) (*dto.TaskResponse, *dto.ErrorResponse)
	UpdateTask(ctx context.Context, request *models.TaskUpdateModel) *dto.ErrorResponse
}

type taskService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(s db.DB, logger *slog.Logger) TaskService {
	return &taskService{
		storage: s,
		logger:  logger,
	}
}

func (s *taskService) GetTasks(ctx context.Context, request *models.TaskFilterModel) (dto.GetTasksResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	if request.Relation == enum.CreatedByMe.String() {
		request.CreatorID = &currUserID
	} else if request.Relation == enum.AssignedToMe.String() {
		request.AssignieID = &currUserID
	}

	response, dberror := s.storage.GetTasksWithFilter(ctx, modelmap.MapToTaskFilter(request))
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetTasksResponse(response), nil
}

func (s *taskService) GetTaskSummary(ctx context.Context, request *models.TaskSummaryParamModel) (*dto.GetTaskSummaryResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	response, dberror := s.storage.GetTaskSummaryByID(ctx, request.TaskID)
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsReturned {
			return nil, errmap.MapToErrorResponse(serverrors.ErrNoTask, http.StatusBadRequest)
		}
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetTaskByIDResponse(response), nil
}

func (s *taskService) CreateTask(ctx context.Context, request *models.TaskCreateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.AuthorID = currUserID

	response, dberror := s.storage.CreateTask(ctx, modelmap.MapToTask(request))
	if dberror != nil {
		if dberror == dberrors.ErrsForeignKeyViolation[dbconsts.ConstraintTaskAssignieIDForeignKey] {
			return nil, errmap.MapToErrorResponse(serverrors.ErrNoUserToAssign, http.StatusBadRequest)
		}
		if dberror == dberrors.ErrsForeignKeyViolation[dbconsts.ConstraintTaskBoardIDForeignKey] {
			return nil, errmap.MapToErrorResponse(serverrors.ErrNoBoardToLink, http.StatusBadRequest)
		}
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToTaskResponse(response), nil
}

func (s *taskService) UpdateTask(ctx context.Context, request *models.TaskUpdateModel) *dto.ErrorResponse {
	err := request.Validate()
	if err != nil {
		return errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.InitiatorID = currUserID

	dberror := s.storage.UpdateTask(ctx, modelmap.MapToTaskUpdate(request))
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsReturned {
			return errmap.MapToErrorResponse(serverrors.ErrManipulationImpossible, http.StatusBadRequest)
		}
		if dberror == dberrors.ErrsForeignKeyViolation[dbconsts.ConstraintTaskAssignieIDForeignKey] {
			return errmap.MapToErrorResponse(serverrors.ErrNoUserToAssign, http.StatusBadRequest)
		}
		if dberror == dberrors.ErrsForeignKeyViolation[dbconsts.ConstraintTaskBoardIDForeignKey] {
			return errmap.MapToErrorResponse(serverrors.ErrNoBoardToLink, http.StatusBadRequest)
		}
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return nil
}
