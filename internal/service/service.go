package service

import (
	"context"
	"log/slog"
	"task-tracker-service/internal/service/comment"
	"task-tracker-service/internal/service/dashboard"
	"task-tracker-service/internal/service/task"
	"task-tracker-service/internal/service/user"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/tokenizer"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
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

func (s *service) RegistrateUser(ctx context.Context, request *models.UserRegisterModel) (*dto.UserResponse, *dto.ErrorResponse) {
	return s.userService.RegistrateUser(ctx, request)
}

func (s *service) LoginUser(ctx context.Context, request *models.UserLoginModel) (*dto.Token, *dto.ErrorResponse) {
	return s.userService.LoginUser(ctx, request)
}

func (s *service) GetUsers(ctx context.Context) (*dto.GetUsersResponse, *dto.ErrorResponse) {
	return s.userService.GetUsers(ctx)
}

func (s *service) AddBoardAdmin(ctx context.Context, request *models.UserBoardAdminActionModel) (*dto.UserResponse, *dto.ErrorResponse) {
	return s.userService.AddBoardAdmin(ctx, request)
}

func (s *service) DeleteBoardAdmin(ctx context.Context, request *models.UserBoardAdminActionModel) (*dto.UserResponse, *dto.ErrorResponse) {
	return s.userService.DeleteBoardAdmin(ctx, request)
}

func (s *service) GetTasks(ctx context.Context, request *models.TaskFilterModel) (*dto.GetTasksResponse, *dto.ErrorResponse) {
	return s.taskService.GetTasks(ctx, request)
}

func (s *service) GetTaskByID(ctx context.Context, request *models.TaskIDParamModel) (*dto.GetTaskByIDResponse, *dto.ErrorResponse) {
	return s.taskService.GetTaskByID(ctx, request)
}

func (s *service) CreateTask(ctx context.Context, request *models.TaskCreateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	return s.taskService.CreateTask(ctx, request)
}

func (s *service) UpdateTask(ctx context.Context, request *models.TaskUpdateModel) (*dto.TaskResponse, *dto.ErrorResponse) {
	return s.taskService.UpdateTask(ctx, request)
}

func (s *service) CreateComment(ctx context.Context, request *models.CommentCreateModel) (*dto.CommentResponse, *dto.ErrorResponse) {
	return s.commentService.CreateComment(ctx, request)
}

func (s *service) GetDashboards(ctx context.Context) (*dto.GetDashboardsResponse, *dto.ErrorResponse) {
	return s.dashboardService.GetDashboards(ctx)
}

func (s *service) GetDashboardByID(ctx context.Context, request *models.DashboardIDParamModel,
) (*dto.GetDashboardByIDResponse, *dto.ErrorResponse) {
	return s.dashboardService.GetDashboardByID(ctx, request)
}

func (s *service) CreateDashboard(ctx context.Context, request *models.DashboardCreateModel) (*dto.DashboardResponse, *dto.ErrorResponse) {
	return s.dashboardService.CreateDashboard(ctx, request)
}

func (s *service) UpdateDashboard(ctx context.Context, request *models.DashboardUpdateModel) (*dto.DashboardResponse, *dto.ErrorResponse) {
	return s.dashboardService.UpdateDashboard(ctx, request)
}

func (s *service) DeleteDashboard(ctx context.Context, request *models.DashboardDeleteModel) *dto.ErrorResponse {
	return s.dashboardService.DeleteDashboard(ctx, request)
}
