package dashboard

import (
	"context"
	"log"
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
	"task-tracker-service/internal/types/models"
)

type DashboardService interface {
	GetDashboards(ctx context.Context) (dto.GetDashboardsResponse, *dto.ErrorResponse)
	GetDashboardSummary(ctx context.Context, request *models.DashboardSummaryParamModel) (*dto.GetDashboardSummaryResponse, *dto.ErrorResponse)
	CreateDashboard(ctx context.Context, request *models.DashboardCreateModel) (*dto.DashboardResponse, *dto.ErrorResponse)
	UpdateDashboard(ctx context.Context, request *models.DashboardUpdateModel) *dto.ErrorResponse
	DeleteDashboard(ctx context.Context, request *models.DashboardDeleteModel) *dto.ErrorResponse
	AddBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse
	DeleteBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse
}

type dashboardService struct {
	storage db.DB
	logger  *slog.Logger
}

func New(s db.DB, logger *slog.Logger) DashboardService {
	return &dashboardService{
		storage: s,
		logger:  logger,
	}
}

func (s *dashboardService) GetDashboards(ctx context.Context) (dto.GetDashboardsResponse, *dto.ErrorResponse) {
	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	response, dberror := s.storage.GetDashboardsForAdminID(ctx, currUserID)
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetDashboardsResponse(response), nil
}

func (s *dashboardService) GetDashboardSummary(ctx context.Context, request *models.DashboardSummaryParamModel,
) (*dto.GetDashboardSummaryResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	response, dberror := s.storage.GetDashboardSummaryByID(ctx, request.BoardID)
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsReturned {
			return nil, errmap.MapToErrorResponse(serverrors.ErrNoDashboard, http.StatusBadRequest)
		}
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetDashboardByIDResponse(response), nil
}

func (s *dashboardService) CreateDashboard(ctx context.Context, request *models.DashboardCreateModel) (*dto.DashboardResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.CreatorID = currUserID

	response, dberror := s.storage.CreateDashboard(ctx, modelmap.MapToDashboard(request))
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToDashboardResponse(response), nil
}

func (s *dashboardService) UpdateDashboard(ctx context.Context, request *models.DashboardUpdateModel) *dto.ErrorResponse {
	err := request.Validate()
	if err != nil {
		return errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.InitiatorID = currUserID

	dberror := s.storage.UpdateDashboard(ctx, modelmap.MapToDashboardUpdate(request))
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsAffected {
			return errmap.MapToErrorResponse(serverrors.ErrManipulationImpossible, http.StatusBadRequest)
		}
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return nil
}

func (s *dashboardService) DeleteDashboard(ctx context.Context, request *models.DashboardDeleteModel) *dto.ErrorResponse {
	err := request.Validate()
	if err != nil {
		return errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.InitiatorID = currUserID

	dberror := s.storage.DeleteDashboard(ctx, modelmap.MapToDashboardDelete(request))
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsAffected {
			return errmap.MapToErrorResponse(serverrors.ErrManipulationImpossible, http.StatusBadRequest)
		}
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return nil
}

func (s *dashboardService) AddBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse {
	err := request.Validate()
	if err != nil {
		log.Println(err)
		return errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.InitiatorID = currUserID

	dberror := s.storage.AddBoardAdmin(ctx, modelmap.MapToDashboardAdminAction(request))
	if dberror != nil {
		if dberror == dberrors.ErrsUniqueCheckViolation[dbconsts.ConstraintBoardToAdminPrimaryKey] {
			return errmap.MapToErrorResponse(serverrors.ErrUserIsAlreadyAdmin, http.StatusBadRequest)
		}
		if dberror == dberrors.ErrsForeignKeyViolation[dbconsts.ConstraintBoardToAdminAdminIDForeignKey] {
			return errmap.MapToErrorResponse(serverrors.ErrNoUser, http.StatusBadRequest)
		}
		if dberror == dberrors.ErrNoRowsAffected {
			return errmap.MapToErrorResponse(serverrors.ErrManipulationImpossible, http.StatusBadRequest)
		}
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return nil
}

func (s *dashboardService) DeleteBoardAdmin(ctx context.Context, request *models.DashboardAdminActionModel) *dto.ErrorResponse {
	err := request.Validate()
	if err != nil {
		return errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	currUserID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.InitiatorID = currUserID

	dberror := s.storage.DeleteBoardAdmin(ctx, modelmap.MapToDashboardAdminAction(request))
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsAffected {
			return errmap.MapToErrorResponse(serverrors.ErrManipulationImpossible, http.StatusBadRequest)
		}
		return errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return nil
}
