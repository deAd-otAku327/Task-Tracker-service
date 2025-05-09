package user

import (
	"context"
	"log/slog"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/tokenizer"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
	"task-tracker-service/pkg/cryptor"
)

type UserService interface {
	RegistrateUser(ctx context.Context, request *models.UserRegisterModel) (*dto.UserResponse, *dto.ErrorResponse)
	LoginUser(ctx context.Context, request *models.UserLoginModel) (*dto.Token, *dto.ErrorResponse)
	GetUsers(ctx context.Context) (*dto.GetUsersResponse, *dto.ErrorResponse)
	AddBoardAdmin(ctx context.Context, request *models.UserBoardAdminModel) (*dto.UserResponse, *dto.ErrorResponse)
	DeleteBoardAdmin(ctx context.Context, request *models.UserBoardAdminModel) (*dto.UserResponse, *dto.ErrorResponse)
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
	return nil, nil
}

func (s *userService) LoginUser(ctx context.Context, request *models.UserLoginModel) (*dto.Token, *dto.ErrorResponse) {
	return nil, nil
}

func (s *userService) GetUsers(ctx context.Context) (*dto.GetUsersResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *userService) AddBoardAdmin(ctx context.Context, request *models.UserBoardAdminActionModel) (*dto.UserResponse, *dto.ErrorResponse) {
	return nil, nil
}

func (s *userService) DeleteBoardAdmin(ctx context.Context, request *models.UserBoardAdminActionModel) (*dto.UserResponse, *dto.ErrorResponse) {
	return nil, nil
}
