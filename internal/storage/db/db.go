package db

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/config"
	"task-tracker-service/internal/storage/db/_shared/consts"
	"task-tracker-service/internal/storage/db/comment"
	"task-tracker-service/internal/storage/db/dashboard"
	"task-tracker-service/internal/storage/db/task"
	"task-tracker-service/internal/storage/db/user"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"

	_ "github.com/lib/pq"
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

func (s *storage) CreateUser(ctx context.Context, createUser *entities.User) (*models.UserModel, error) {
	return s.userStorage.CreateUser(ctx, createUser)
}

func (s *storage) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	return s.userStorage.GetUserByUsername(ctx, username)
}

func (s *storage) GetUsers(ctx context.Context) ([]*models.UserModel, error) {
	return s.userStorage.GetUsers(ctx)
}

func (s *storage) AddBoardAdminUser(ctx context.Context, boardAdminAction *entities.UserBoardAdminAction) (*models.UserModel, error) {
	return s.userStorage.AddBoardAdminUser(ctx, boardAdminAction)
}

func (s *storage) DeleteBoardAdminUser(ctx context.Context, boardAdminAction *entities.UserBoardAdminAction) error {
	return s.userStorage.DeleteBoardAdminUser(ctx, boardAdminAction)
}

func (s *storage) GetTasksWithFilter(ctx context.Context, filter *entities.TaskFilter) ([]*models.TaskModel, error) {
	return s.taskStorage.GetTasksWithFilter(ctx, filter)
}

func (s *storage) GetTaskByID(ctx context.Context, taskID int) (*models.TaskModel, []*models.CommentModel, *models.UserModel,
	*models.UserModel, *models.DashboardModel, error) {
	return s.taskStorage.GetTaskByID(ctx, taskID)
}

func (s *storage) CreateTask(ctx context.Context, task *entities.Task) (*models.TaskModel, error) {
	return s.taskStorage.CreateTask(ctx, task)
}

func (s *storage) UpdateTask(ctx context.Context, taskUpdate *entities.TaskUpdate) (*models.TaskModel, error) {
	return s.taskStorage.UpdateTask(ctx, taskUpdate)
}

func (s *storage) CreateComment(ctx context.Context, comment *entities.Comment) (*models.CommentModel, error) {
	return s.commentStorage.CreateComment(ctx, comment)
}

func (s *storage) GetDashboardsForAdminID(ctx context.Context, userID int) ([]*models.DashboardModel, error) {
	return s.dashboardStorage.GetDashboardsForAdminID(ctx, userID)
}

func (s *storage) GetDashboardByID(ctx context.Context, boardID int) (*models.DashboardModel, []*models.TaskModel, []*models.UserModel, error) {
	return s.dashboardStorage.GetDashboardByID(ctx, boardID)
}

func (s *storage) CreateDashboard(ctx context.Context, dashboard *entities.Dashboard) (*models.DashboardModel, error) {
	return s.dashboardStorage.CreateDashboard(ctx, dashboard)
}

func (s *storage) UpdateDashboard(ctx context.Context, dashboardUpdate *entities.DashboardUpdate) (*models.DashboardModel, error) {
	return s.dashboardStorage.UpdateDashboard(ctx, dashboardUpdate)
}

func (s *storage) DeleteDashboard(ctx context.Context, dashboardDelete *entities.DashboardDelete) error {
	return s.dashboardStorage.DeleteDashboard(ctx, dashboardDelete)
}
