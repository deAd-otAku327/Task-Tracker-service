package dtomap

import (
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
)

func MapToUserLoginModel(request *dto.PostUsersLoginRequest) *models.UserLoginModel {
	return &models.UserLoginModel{
		Username: request.Username,
		Password: request.Password,
	}
}

func MapToUserRegisterModel(request *dto.PostUsersRegisterRequest) *models.UserRegisterModel {
	return &models.UserRegisterModel{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}
}

func MapToUserBoardAdminModel(request *dto.PostUsersBoardAdminRequest) *models.UserBoardAdminActionModel {
	return &models.UserBoardAdminActionModel{
		BoardID: request.BoardID,
		UserID:  request.UserID,
	}
}

func MapToTaskFilterModel(request *dto.GetTasksParams) *models.TaskFilterModel {
	return &models.TaskFilterModel{
		Relation: request.Relation,
		Status:   request.Status,
	}
}

func MapToTaskIDParamModel(request *dto.GetTaskByIDParam) *models.TaskIDParamModel {
	return &models.TaskIDParamModel{
		TaskID: request.TaskID,
	}
}

func MapToTaskCreateModel(request *dto.PostTasksCreateRequest) *models.TaskCreateModel {
	return &models.TaskCreateModel{
		Title:         request.Title,
		Discription:   request.Discription,
		AssignieID:    request.AssignieID,
		LinkedBoardID: request.LinkedBoardID,
	}
}

func MapToTaskUpdateModel(request *dto.PostTasksUpdateRequest) *models.TaskUpdateModel {
	return &models.TaskUpdateModel{
		ID:            request.TaskID,
		Title:         request.Title,
		Status:        request.Status,
		Discription:   request.Discription,
		AssignieID:    request.AssignieID,
		LinkedBoardID: request.LinkedBoardID,
	}
}

func MapToCommentCreateModel(request *dto.PostCommentRequest) *models.CommentCreateModel {
	return &models.CommentCreateModel{
		TaskID: request.TaskID,
		Text:   request.Text,
	}
}

func MapToDashboardIDParamModel(request *dto.GetDashboardByIDParam) *models.DashboardIDParamModel {
	return &models.DashboardIDParamModel{
		BoardID: request.BoardID,
	}
}

func MapToDashboardCreateModel(request *dto.PostDashboardsCreateRequest) *models.DashboardCreateModel {
	return &models.DashboardCreateModel{
		Title:       request.Title,
		Discription: request.Discription,
	}
}

func MapToDashboardUpdateModel(request *dto.PostDashboardsUpdateRequest) *models.DashboardUpdateModel {
	return &models.DashboardUpdateModel{
		ID:    request.BoardID,
		Title: request.Title,
	}
}

func MapToDashboardDeleteModel(request *dto.PostDashboardsDeleteRequest) *models.DashboardDeleteModel {
	return &models.DashboardDeleteModel{
		BoardID: request.BoardID,
	}
}
