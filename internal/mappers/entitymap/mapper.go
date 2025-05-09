package entitymap

import (
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

func MapToUserModel(response *entities.User) *models.UserModel {
	return &models.UserModel{
		ID:             response.ID,
		Username:       response.Username,
		Email:          response.Email,
		HashedPassword: response.Password,
	}
}

func MapToTaskModel(response *entities.Task) *models.TaskModel {
	return &models.TaskModel{
		ID:    response.ID,
		Title: response.Title,
		Discription: func() *string {
			if response.Discription.Valid {
				return &response.Discription.String
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

func MapToCommentModel(response *entities.Comment) *models.CommentModel {
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
	return &models.DashboardModel{
		ID:    response.ID,
		Title: response.Title,
		Discription: func() *string {
			if response.Discription.Valid {
				return &response.Discription.String
			}
			return nil
		}(),
		UpdatedAt: response.UpdatedAt,
	}
}

func MapToDashboardSummaryModel(respBoard *entities.Dashboard, respTasks []*entities.Task,
	respAdmins []*entities.User) *models.DashboardSummaryModel {

	return &models.DashboardSummaryModel{
		Dashboard: MapToDashboardModel(respBoard),
		Tasks: func() []*models.TaskModel {
			res := make([]*models.TaskModel, 0, len(respTasks))
			for _, entity := range respTasks {
				res = append(res, MapToTaskModel(entity))
			}
			return res
		}(),
		Admins: func() []*models.UserModel {
			res := make([]*models.UserModel, 0, len(respAdmins))
			for _, entity := range respAdmins {
				res = append(res, MapToUserModel(entity))
			}
			return res
		}(),
	}
}
