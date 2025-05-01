package user

import (
	"database/sql"
	"log/slog"
)

type UserDB interface {
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
