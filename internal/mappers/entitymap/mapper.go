package entitymap

import (
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

func MapToUserModel(response *entities.User) *models.UserModel {
	if response == nil {
		return nil
	}
	return &models.UserModel{
		ID:             response.ID,
		Username:       response.Username,
		Email:          response.Email,
		HashedPassword: response.Password,
	}
}

func MapToUserListModel(response []*entities.User) models.UserListModel {
	if response == nil {
		return nil
	}
	result := make(models.UserListModel, 0, len(response))
	for _, u := range response {
		result = append(result, MapToUserModel(u))
	}

	return result
}

func MapToTaskModel(response *entities.Task) *models.TaskModel {
	if response == nil {
		return nil
	}
	return &models.TaskModel{
		ID:    response.ID,
		Title: response.Title,
		Description: func() *string {
			if response.Description.Valid {
				return &response.Description.String
			}
			return nil
		}(),
		Status: response.Status,
		AssignieID: func() *int {
			if response.AssignieID.Valid {
				res := int(response.AssignieID.Int32)
				return &res
			}
			return nil
		}(),
		BoardID: func() *int {
			if response.BoardID.Valid {
				res := int(response.BoardID.Int32)
				return &res
			}
			return nil
		}(),
		UpdatedAt: response.UpdatedAt,
	}
}

func MapToTaskListModel(response []*entities.Task) models.TaskListModel {
	if response == nil {
		return nil
	}
	result := make(models.TaskListModel, 0, len(response))
	for _, t := range response {
		result = append(result, MapToTaskModel(t))
	}

	return result
}

func MapToCommentModel(response *entities.Comment) *models.CommentModel {
	if response == nil {
		return nil
	}
	return &models.CommentModel{
		ID:       response.ID,
		AuthorID: response.AuthorID,
		Text:     response.Text,
		DateTime: response.DateTime,
	}
}

func MapToTaskSummaryModel(respTask *entities.Task, respComms []*entities.Comment,
	respAuthor, respAssignie *entities.User, respBoard *entities.Dashboard) *models.TaskSummaryModel {

	return &models.TaskSummaryModel{
		Task: MapToTaskModel(respTask),
		Comments: func() []*models.CommentModel {
			res := make([]*models.CommentModel, 0, len(respComms))
			for _, entity := range respComms {
				res = append(res, MapToCommentModel(entity))
			}
			return res
		}(),
		Author:      MapToUserModel(respAuthor),
		Assignie:    MapToUserModel(respAssignie),
		LinkedBoard: MapToDashboardModel(respBoard),
	}
}

func MapToDashboardModel(response *entities.Dashboard) *models.DashboardModel {
	if response == nil {
		return nil
	}
	return &models.DashboardModel{
		ID:    response.ID,
		Title: response.Title,
		Description: func() *string {
			if response.Description.Valid {
				return &response.Description.String
			}
			return nil
		}(),
		UpdatedAt: response.UpdatedAt,
	}
}

func MapToDashboardListModel(response []*entities.Dashboard) models.DashboardListModel {
	if response == nil {
		return nil
	}
	result := make(models.DashboardListModel, 0, len(response))
	for _, d := range response {
		result = append(result, MapToDashboardModel(d))
	}

	return result
}

func MapToDashboardSummaryModel(respBoard *entities.Dashboard, respTasks []*entities.Task,
	respAdmins []*entities.User) *models.DashboardSummaryModel {

	return &models.DashboardSummaryModel{
		Dashboard: MapToDashboardModel(respBoard),
		Tasks:     MapToTaskListModel(respTasks),
		Admins:    MapToUserListModel(respAdmins),
	}
}
