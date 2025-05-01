package service

import (
	"log/slog"
	"task-tracker-service/internal/service/comment"
	"task-tracker-service/internal/service/dashboard"
	"task-tracker-service/internal/service/task"
	"task-tracker-service/internal/service/user"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/tokenizer"
	"task-tracker-service/pkg/cryptor"
)

type Service interface {
	user.UserService
	task.TaskService
	comment.CommentService
	dashboard.DashboardService
}

type service struct {
	userService      user.UserService
	taskService      task.TaskService
	commentService   comment.CommentService
	dashboardService dashboard.DashboardService
}

func New(s db.DB, logger *slog.Logger, cryptor cryptor.Cryptor, tok tokenizer.Tokenizer) Service {
	return &service{
		userService:      user.New(s, logger, cryptor, tok),
		taskService:      task.New(s, logger),
		commentService:   comment.New(s, logger),
		dashboardService: dashboard.New(s, logger),
	}
}
