package user

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/mappers/entitymap"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"

	"task-tracker-service/internal/storage/db/_shared/helpers"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"

	sq "github.com/Masterminds/squirrel"
)

type UserDB interface {
	CreateUser(ctx context.Context, createUser *entities.User) (*models.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error)
	GetUsers(ctx context.Context) (models.UserListModel, error)
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
	insertQuery, args, err := sq.Insert(dbconsts.TableUsers).
		Columns(dbconsts.ColumnUserName, dbconsts.ColumnUserEmail, dbconsts.ColumnUserPassword).
		Values(createUser.Username, createUser.Email, createUser.Password).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, insertQuery, args...)

	var user entities.User

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, helpers.CatchPQErrors(err)
	}

	return entitymap.MapToUserModel(&user), nil
}

func (s *userStorage) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	query, args, err := sq.Select("*").
		From(dbconsts.TableUsers).
		Where(sq.Eq{dbconsts.ColumnUserName: username}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var user entities.User

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dberrors.ErrNoRowsReturned
		}
		return nil, err
	}

	return entitymap.MapToUserModel(&user), nil
}

func (s *userStorage) GetUsers(ctx context.Context) (models.UserListModel, error) {
	query, args, err := sq.Select("*").
		From(dbconsts.TableUsers).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)

	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return entitymap.MapToUserListModel(users), nil
}
