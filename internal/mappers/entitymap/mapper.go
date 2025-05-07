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

func MapToCommentModel(response *entities.Comment, authorName string) *models.CommentModel {
	return &models.CommentModel{
		ID:         response.ID,
		AuthorID:   response.AuthorID,
		AuthorName: authorName,
		Text:       response.Text,
		DateTime:   response.DateTime,
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
