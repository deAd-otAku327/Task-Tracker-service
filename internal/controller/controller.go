package controller

import (
	"log/slog"
	"task-tracker-service/internal/controller/comment"
	"task-tracker-service/internal/controller/dashboard"
	"task-tracker-service/internal/controller/task"
	"task-tracker-service/internal/controller/user"
	"task-tracker-service/internal/service"
)

type Controller interface {
	user.UserHandler
	task.TaskHandler
	comment.CommentHandler
	dashboard.DashboardHandler
}

type controller struct {
	userHandler      user.UserHandler
	taskHandler      task.TaskHandler
	commentHandler   comment.CommentHandler
	dashboardHandler dashboard.DashboardHandler
}

func New(s service.Service, logger *slog.Logger) Controller {
	return &controller{
		userHandler:      user.New(s, logger),
		taskHandler:      task.New(s, logger),
		commentHandler:   comment.New(s, logger),
		dashboardHandler: dashboard.New(s, logger),
	}
}
