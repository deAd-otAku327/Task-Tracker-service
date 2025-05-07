package modelmap

import (
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
)

const DateTimeFormat = "2006-01-02 15:04:05"

func MapToUserResponse(response *models.UserModel) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       response.ID,
		Username: response.Username,
		Email:    response.Email,
	}
}

func MapToGetUsersResponse(response []*models.UserModel) *dto.GetUsersResponse {
	res := make(dto.GetUsersResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToUserResponse(model))
	}
	return &res
}

func MapToTaskResponse(response *models.TaskModel) *dto.TaskResponse {
	return &dto.TaskResponse{
		ID:          response.ID,
		Title:       response.Title,
		Discription: response.Discription,
		Status:      response.Status,
		Assignie: func() *bool {
			if response.AssignieID != nil {
				res := true
				return &res
			}
			return nil
		}(),
		Board: func() *bool {
			if response.BoardID != nil {
				res := true
				return &res
			}
			return nil
		}(),
		UpdatedAt: response.UpdatedAt.Format(DateTimeFormat),
	}
}

func MapToGetTaskResponse(response []*models.TaskModel) *dto.GetTasksResponse {
	res := make(dto.GetTasksResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToTaskResponse(model))
	}
	return &res
}

func MapToGetTaskByIDResponse(respTask *models.TaskModel, respComments []*models.CommentModel,
	respAuthor, respAssignie *models.UserModel, respDashboard *models.DashboardModel,
) *dto.GetTaskByIDResponse {
	return &dto.GetTaskByIDResponse{
		Task: MapToTaskResponse(respTask),
		Comments: func() []*dto.CommentResponse {
			res := make([]*dto.CommentResponse, 0, len(respComments))
			for _, model := range respComments {
				res = append(res, MapToCommentResponse(model))
			}
			return res
		}(),
		Author:   MapToUserResponse(respAuthor),
		Assignie: MapToUserResponse(respAssignie),
		LinkedBoard: &dto.BoardData{
			ID:    respDashboard.ID,
			Title: respDashboard.Title,
		},
	}
}

func MapToCommentResponse(response *models.CommentModel) *dto.CommentResponse {
	return &dto.CommentResponse{
		ID:         response.ID,
		AuthorName: response.AuthorName,
		Text:       response.Text,
		DateTime:   response.DateTime.Format(DateTimeFormat),
	}
}

func MapToDashboardResponse(response *models.DashboardModel) *dto.DashboardResponse {
	return &dto.DashboardResponse{
		ID:          response.ID,
		Title:       response.Title,
		Discription: response.Discription,
		UpdatedAt:   response.UpdatedAt.Format(DateTimeFormat),
	}
}

func MapToGetDashboardsResponse(response []*models.DashboardModel) *dto.GetDashboardsResponse {
	res := make(dto.GetDashboardsResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToDashboardResponse(model))
	}
	return &res
}

func MapToGetDashboardByIDResponse(respDashboard *models.DashboardModel, respTasks []*models.TaskModel, respAdmins []*models.UserModel,
) *dto.GetDashboardByIDResponse {
	return &dto.GetDashboardByIDResponse{
		Dashboard: MapToDashboardResponse(respDashboard),
		Tasks: func() []*dto.TaskResponse {
			res := make([]*dto.TaskResponse, 0, len(respTasks))
			for _, model := range respTasks {
				res = append(res, MapToTaskResponse(model))
			}
			return res
		}(),
		Admins: func() []*dto.UserResponse {
			res := make([]*dto.UserResponse, 0, len(respAdmins))
			for _, model := range respAdmins {
				res = append(res, MapToUserResponse(model))
			}
			return res
		}(),
	}
}
