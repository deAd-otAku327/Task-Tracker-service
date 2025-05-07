package modelmap

import (
	"database/sql"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"
)

func MapToUser(request *models.UserRegisterModel) *entities.User {
	return &entities.User{
		// ID ommited cause of unregistred entity.
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
}

func MapToBoardAdmin(request *models.UserBoardAdminModel) *entities.UserBoardAdmin {
	return &entities.UserBoardAdmin{
		BoardID: request.BoardID,
		UserID:  request.UserID,
	}
}

func MapToTaskFilter(request *models.TaskFilterModel) *entities.TaskFilter {
	return &entities.TaskFilter{
		Relation: request.Relation,
		Status:   request.Status,
	}
}

func MapToTask(request *models.TaskCreateModel) *entities.Task {
	return &entities.Task{
		// ID ommited cause of unregistred entity.
		// Status ommited cause of db defaults and dependencies reducing.
		Title: request.Title,
		Discription: func() sql.NullString {
			if request.Discription != nil {
				return sql.NullString{String: *request.Discription}
			}
			return sql.NullString{Valid: false}
		}(),
		AuthorID: request.AuthorID,
		AssignieID: func() sql.NullInt32 {
			if request.AssignieID != nil {
				return sql.NullInt32{Int32: int32(*request.AssignieID)}
			}
			return sql.NullInt32{Valid: false}
		}(),
		BoardID: func() sql.NullInt32 {
			if request.LinkedBoardID != nil {
				return sql.NullInt32{Int32: int32(*request.LinkedBoardID)}
			}
			return sql.NullInt32{Valid: false}
		}(),
	}
}

func MapToTaskUpdate(request *models.TaskUpdateModel) *entities.TaskUpdate {
	return &entities.TaskUpdate{
		ID:          request.ID,
		Title:       request.Title,
		Status:      request.Status,
		Discription: request.Discription,
		AssignieID:  request.AssignieID,
		BoardID:     request.LinkedBoardID,
	}
}

func MapToComment(request *models.CommentCreateModel) *entities.Comment {
	return &entities.Comment{
		// ID ommited cause of unregistred entity.
		AuthorID: request.AuthorID,
		Text:     request.Text,
	}
}

func MapToDashboard(request *models.DashboardCreateModel) *entities.Dashboard {
	return &entities.Dashboard{
		// ID ommited cause of unregistred entity.
		Title: request.Title,
		Discription: func() sql.NullString {
			if request.Discription != nil {
				return sql.NullString{String: *request.Discription}
			}
			return sql.NullString{Valid: false}
		}(),
	}
}

func MapToDashboardUpdate(request *models.DashboardUpdateModel) *entities.DashboardUpdate {
	return &entities.DashboardUpdate{
		ID:          request.ID,
		Title:       request.Title,
		Discription: request.Discription,
	}
}
