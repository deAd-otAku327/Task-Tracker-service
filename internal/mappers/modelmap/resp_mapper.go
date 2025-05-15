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

func MapToGetUsersResponse(response models.UserListModel) dto.GetUsersResponse {
	res := make(dto.GetUsersResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToUserResponse(model))
	}
	return res
}

func MapToTaskResponse(response *models.TaskModel) *dto.TaskResponse {
	return &dto.TaskResponse{
		ID:          response.ID,
		Title:       response.Title,
		Description: response.Description,
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

func MapToGetTasksResponse(response models.TaskListModel) dto.GetTasksResponse {
	res := make(dto.GetTasksResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToTaskResponse(model))
	}
	return res
}

func MapToGetTaskByIDResponse(response *models.TaskSummaryModel) *dto.GetTaskByIDResponse {
	return &dto.GetTaskByIDResponse{
		Task: MapToTaskResponse(response.Task),
		Comments: func() []*dto.CommentResponse {
			res := make([]*dto.CommentResponse, 0, len(response.Comments))
			for _, model := range response.Comments {
				res = append(res, MapToCommentResponse(model))
			}
			return res
		}(),
		Author:   MapToUserResponse(response.Author),
		Assignie: MapToUserResponse(response.Assignie),
		LinkedBoard: &dto.BoardData{
			ID:    response.LinkedBoard.ID,
			Title: response.LinkedBoard.Title,
		},
	}
}

func MapToCommentResponse(response *models.CommentModel) *dto.CommentResponse {
	return &dto.CommentResponse{
		ID:       response.ID,
		AuthorID: response.AuthorID,
		Text:     response.Text,
		DateTime: response.DateTime.Format(DateTimeFormat),
	}
}

func MapToDashboardResponse(response *models.DashboardModel) *dto.DashboardResponse {
	return &dto.DashboardResponse{
		ID:          response.ID,
		Title:       response.Title,
		Description: response.Description,
		UpdatedAt:   response.UpdatedAt.Format(DateTimeFormat),
	}
}

func MapToGetDashboardsResponse(response models.DashboardListModel) dto.GetDashboardsResponse {
	res := make(dto.GetDashboardsResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToDashboardResponse(model))
	}
	return res
}

func MapToGetDashboardByIDResponse(response *models.DashboardSummaryModel) *dto.GetDashboardByIDResponse {
	return &dto.GetDashboardByIDResponse{
		Dashboard: MapToDashboardResponse(response.Dashboard),
		Tasks: func() []*dto.TaskResponse {
			res := make([]*dto.TaskResponse, 0, len(response.Tasks))
			for _, model := range response.Tasks {
				res = append(res, MapToTaskResponse(model))
			}
			return res
		}(),
		Admins: func() []*dto.UserResponse {
			res := make([]*dto.UserResponse, 0, len(response.Admins))
			for _, model := range response.Admins {
				res = append(res, MapToUserResponse(model))
			}
			return res
		}(),
	}
}
