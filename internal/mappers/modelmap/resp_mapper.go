package modelmap

import (
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/models"
)

const DateTimeFormat = "2006-01-02 15:04:05"

func MapToUserResponse(response *models.UserModel) *dto.UserResponse {
	if response == nil {
		return nil
	}
	return &dto.UserResponse{
		ID:       response.ID,
		Username: response.Username,
		Email:    response.Email,
	}
}

func MapToGetUsersResponse(response models.UserListModel) dto.GetUsersResponse {
	if response == nil {
		return nil
	}
	res := make(dto.GetUsersResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToUserResponse(model))
	}
	return res
}

func MapToTaskResponse(response *models.TaskModel) *dto.TaskResponse {
	if response == nil {
		return nil
	}
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
	if response == nil {
		return nil
	}
	res := make(dto.GetTasksResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToTaskResponse(model))
	}
	return res
}

func MapToGetTaskByIDResponse(response *models.TaskSummaryModel) *dto.GetTaskSummaryResponse {
	return &dto.GetTaskSummaryResponse{
		Task: MapToTaskResponse(response.Task),
		Comments: func() []*dto.CommentResponse {
			if response.Comments == nil {
				return nil
			}
			res := make([]*dto.CommentResponse, 0, len(response.Comments))
			for _, model := range response.Comments {
				res = append(res, MapToCommentResponse(model))
			}
			return res
		}(),
		Author:   MapToUserResponse(response.Author),
		Assignie: MapToUserResponse(response.Assignie),
		LinkedBoard: func() *dto.BoardData {
			if response.LinkedBoard != nil {
				return &dto.BoardData{
					ID:    response.LinkedBoard.ID,
					Title: response.LinkedBoard.Title,
				}
			}
			return nil
		}(),
	}
}

func MapToCommentResponse(response *models.CommentModel) *dto.CommentResponse {
	if response == nil {
		return nil
	}
	return &dto.CommentResponse{
		ID:       response.ID,
		AuthorID: response.AuthorID,
		Text:     response.Text,
		DateTime: response.DateTime.Format(DateTimeFormat),
	}
}

func MapToDashboardResponse(response *models.DashboardModel) *dto.DashboardResponse {
	if response == nil {
		return nil
	}
	return &dto.DashboardResponse{
		ID:          response.ID,
		Title:       response.Title,
		Description: response.Description,
		UpdatedAt:   response.UpdatedAt.Format(DateTimeFormat),
	}
}

func MapToGetDashboardsResponse(response models.DashboardListModel) dto.GetDashboardsResponse {
	if response == nil {
		return nil
	}
	res := make(dto.GetDashboardsResponse, 0, len(response))
	for _, model := range response {
		res = append(res, MapToDashboardResponse(model))
	}
	return res
}

func MapToGetDashboardByIDResponse(response *models.DashboardSummaryModel) *dto.GetDashboardSummaryResponse {
	return &dto.GetDashboardSummaryResponse{
		Dashboard: MapToDashboardResponse(response.Dashboard),
		Tasks:     MapToGetTasksResponse(response.Tasks),
		Admins:    MapToGetUsersResponse(response.Admins),
	}
}
