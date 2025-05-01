package dashboard

import (
	"database/sql"
	"log/slog"
)

type DashboardDB interface {
}

type dashboardStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) DashboardDB {
	return &dashboardStorage{
		db:     db,
		logger: logger,
	}
}
