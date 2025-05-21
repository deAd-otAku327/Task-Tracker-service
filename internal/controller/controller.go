package controller

import (
	"log/slog"
	"net/http"
	"task-tracker-service/internal/controller/comment"
	"task-tracker-service/internal/controller/dashboard"
	"task-tracker-service/internal/controller/task"
	"task-tracker-service/internal/controller/user"
	"task-tracker-service/internal/service"
	"time"
)

type Controller interface {
	user.UserHandler
	task.TaskHandler
	comment.CommentHandler
	dashboard.DashboardHandler
}

type HandlersConfig struct {
	AuthExpire time.Duration
}

type controller struct {
	userHandler      user.UserHandler
	taskHandler      task.TaskHandler
	commentHandler   comment.CommentHandler
	dashboardHandler dashboard.DashboardHandler
}

func New(s service.Service, logger *slog.Logger, params *HandlersConfig) Controller {
	return &controller{
		userHandler:      user.New(s, logger, params.AuthExpire),
		taskHandler:      task.New(s, logger),
		commentHandler:   comment.New(s, logger),
		dashboardHandler: dashboard.New(s, logger),
	}
}

func (c *controller) Register() http.HandlerFunc {
	return c.userHandler.Register()
}

func (c *controller) Login() http.HandlerFunc {
	return c.userHandler.Login()
}

func (c *controller) GetUsers() http.HandlerFunc {
	return c.userHandler.GetUsers()
}

func (c *controller) GetTasks() http.HandlerFunc {
	return c.taskHandler.GetTasks()
}

func (c *controller) GetTaskSummary() http.HandlerFunc {
	return c.taskHandler.GetTaskSummary()
}

func (c *controller) CreateTask() http.HandlerFunc {
	return c.taskHandler.CreateTask()
}

func (c *controller) UpdateTask() http.HandlerFunc {
	return c.taskHandler.UpdateTask()
}

func (c *controller) Comment() http.HandlerFunc {
	return c.commentHandler.Comment()
}

func (c *controller) GetDashboards() http.HandlerFunc {
	return c.dashboardHandler.GetDashboards()
}

func (c *controller) GetDashboardByID() http.HandlerFunc {
	return c.dashboardHandler.GetDashboardByID()
}

func (c *controller) CreateDashboard() http.HandlerFunc {
	return c.dashboardHandler.CreateDashboard()
}

func (c *controller) UpdateDashboard() http.HandlerFunc {
	return c.dashboardHandler.UpdateDashboard()
}

func (c *controller) DeleteDashboard() http.HandlerFunc {
	return c.dashboardHandler.DeleteDashboard()
}

func (c *controller) AddBoardAdmin() http.HandlerFunc {
	return c.dashboardHandler.AddBoardAdmin()
}

func (c *controller) DeleteBoardAdmin() http.HandlerFunc {
	return c.dashboardHandler.DeleteBoardAdmin()
}
