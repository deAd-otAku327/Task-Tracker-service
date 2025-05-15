package user

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"task-tracker-service/internal/mappers/errmap"
	"task-tracker-service/internal/mappers/modelmap"
	"task-tracker-service/internal/service/_shared/serverrors"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"
	"task-tracker-service/internal/tokenizer"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
	"task-tracker-service/pkg/cryptor"
)

type UserService interface {
	RegistrateUser(ctx context.Context, request *models.UserRegisterModel) (*dto.UserResponse, *dto.ErrorResponse)
	LoginUser(ctx context.Context, request *models.UserLoginModel) (*dto.Token, *dto.ErrorResponse)
	GetUsers(ctx context.Context) (dto.GetUsersResponse, *dto.ErrorResponse)
}

type userService struct {
	storage   db.DB
	logger    *slog.Logger
	cryptor   cryptor.Cryptor
	tokenizer tokenizer.Tokenizer
}

func New(s db.DB, logger *slog.Logger, cryptor cryptor.Cryptor, tok tokenizer.Tokenizer) UserService {
	return &userService{
		storage:   s,
		logger:    logger,
		cryptor:   cryptor,
		tokenizer: tok,
	}
}

func (s *userService) RegistrateUser(ctx context.Context, request *models.UserRegisterModel) (*dto.UserResponse, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	hash, err := s.cryptor.EncryptPassword(request.Password)
	if err != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	request.HashedPassword = hash

	response, dberror := s.storage.CreateUser(ctx, modelmap.MapToUser(request))
	if dberror != nil {
		if dberror == dberrors.ErrsUniqueCheckViolation[dbconsts.ConstraintUserUniqueUsername] {
			return nil, errmap.MapToErrorResponse(serverrors.ErrUsernameOccupied, http.StatusBadRequest)
		}
		if dberror == dberrors.ErrsUniqueCheckViolation[dbconsts.ConstraintUserUniqueEmail] {
			return nil, errmap.MapToErrorResponse(serverrors.ErrEmailOccupied, http.StatusBadRequest)
		}
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToUserResponse(response), nil
}

func (s *userService) LoginUser(ctx context.Context, request *models.UserLoginModel) (*dto.Token, *dto.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, errmap.MapToErrorResponse(err, http.StatusBadRequest)
	}

	response, dberror := s.storage.GetUserByUsername(ctx, request.Username)
	if dberror != nil {
		if dberror == dberrors.ErrNoRowsReturned {
			return nil, errmap.MapToErrorResponse(serverrors.ErrUsernameIsNotRegistered, http.StatusUnauthorized)
		}
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	if err = s.cryptor.CompareHashAndPassword(response.HashedPassword, request.Password); err != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrInvalidPassword, http.StatusUnauthorized)
	}

	token, err := s.tokenizer.GenerateToken(strconv.Itoa(response.ID))
	if err != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return (*dto.Token)(token), nil
}

func (s *userService) GetUsers(ctx context.Context) (dto.GetUsersResponse, *dto.ErrorResponse) {
	response, dberror := s.storage.GetUsers(ctx)
	if dberror != nil {
		return nil, errmap.MapToErrorResponse(serverrors.ErrSomethingWentWrong, http.StatusInternalServerError)
	}

	return modelmap.MapToGetUsersResponse(response), nil
}
