package comment

import (
	"database/sql"
	"log/slog"
)

type CommentDB interface {
}

type commentStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) CommentDB {
	return &commentStorage{
		db:     db,
		logger: logger,
	}
}
