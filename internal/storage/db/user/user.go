package user

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

type UserDB interface {
	CreateUser(ctx context.Context, createUser *entities.User) (*models.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error)
	GetUsers(ctx context.Context) ([]*models.UserModel, error)
	AddBoardAdminUser(ctx context.Context, boardAdminAction *entities.UserBoardAdminAction) (*models.UserModel, error)
	DeleteBoardAdminUser(ctx context.Context, boardAdminAction *entities.UserBoardAdminAction) error
}

type userStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) UserDB {
	return &userStorage{
		db:     db,
		logger: logger,
	}
}

func (s *userStorage) CreateUser(ctx context.Context, createUser *entities.User) (*models.UserModel, error) {
	return nil, nil
}

func (s *userStorage) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	return nil, nil
}

func (s *userStorage) GetUsers(ctx context.Context) ([]*models.UserModel, error) {
	return nil, nil
}

func (s *userStorage) AddBoardAdminUser(ctx context.Context, boardAdminAction *entities.UserBoardAdminAction) (*models.UserModel, error) {
	return nil, nil
}

func (s *userStorage) DeleteBoardAdminUser(ctx context.Context, boardAdminAction *entities.UserBoardAdminAction) error {
	return nil
}
