package db

import (
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/config"
	"task-tracker-service/internal/storage/db/_shared/consts"
	"task-tracker-service/internal/storage/db/comment"
	"task-tracker-service/internal/storage/db/dashboard"
	"task-tracker-service/internal/storage/db/task"
	"task-tracker-service/internal/storage/db/user"
)

type DB interface {
	user.UserDB
	task.TaskDB
	comment.CommentDB
	dashboard.DashboardDB
}

type storage struct {
	userStorage      user.UserDB
	taskStorage      task.TaskDB
	commentStorage   comment.CommentDB
	dashboardStorage dashboard.DashboardDB
}

func New(cfg config.DBConn, logger *slog.Logger) (DB, error) {
	database, err := sql.Open(consts.PGDriverName, cfg.URL)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)

	return &storage{
		userStorage:      user.New(database, logger),
		taskStorage:      task.New(database, logger),
		commentStorage:   comment.New(database, logger),
		dashboardStorage: dashboard.New(database, logger),
	}, nil
}
