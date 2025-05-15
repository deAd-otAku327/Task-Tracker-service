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
		Password: request.HashedPassword,
	}
}

func MapToTaskFilter(request *models.TaskFilterModel) *entities.TaskFilter {
	return &entities.TaskFilter{
		Status:     request.Status,
		CreatorID:  request.CreatorID,
		AssignieID: request.AssignieID,
	}
}

func MapToTask(request *models.TaskCreateModel) *entities.Task {
	return &entities.Task{
		// ID ommited cause of unregistred entity.
		// Status ommited cause of db defaults and dependencies reducing.
		Title: request.Title,
		Description: func() sql.NullString {
			if request.Description != nil {
				return sql.NullString{String: *request.Description, Valid: true}
			}
			return sql.NullString{Valid: false}
		}(),
		AuthorID: request.AuthorID,
		AssignieID: func() sql.NullInt32 {
			if request.AssignieID != nil {
				return sql.NullInt32{Int32: int32(*request.AssignieID), Valid: true}
			}
			return sql.NullInt32{Valid: false}
		}(),
		BoardID: func() sql.NullInt32 {
			if request.LinkedBoardID != nil {
				return sql.NullInt32{Int32: int32(*request.LinkedBoardID), Valid: true}
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
		Description: request.Description,
		AssignieID:  request.AssignieID,
		BoardID:     request.LinkedBoardID,
		InitiatorID: request.InitiatorID,
	}
}

func MapToComment(request *models.CommentCreateModel) *entities.Comment {
	return &entities.Comment{
		// ID ommited cause of unregistred entity.
		TaskID:   request.TaskID,
		AuthorID: request.AuthorID,
		Text:     request.Text,
	}
}

func MapToDashboard(request *models.DashboardCreateModel) *entities.Dashboard {
	return &entities.Dashboard{
		// ID ommited cause of unregistred entity.
		Title:     request.Title,
		CreatorID: request.CreatorID,
		Description: func() sql.NullString {
			if request.Description != nil {
				return sql.NullString{String: *request.Description, Valid: true}
			}
			return sql.NullString{Valid: false}
		}(),
	}
}

func MapToDashboardUpdate(request *models.DashboardUpdateModel) *entities.DashboardUpdate {
	return &entities.DashboardUpdate{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		InitiatorID: request.InitiatorID,
	}
}

func MapToDashboardDelete(request *models.DashboardDeleteModel) *entities.DashboardDelete {
	return &entities.DashboardDelete{
		BoardID:     request.BoardID,
		InitiatorID: request.InitiatorID,
	}
}

func MapToDashboardAdminAction(request *models.DashboardAdminActionModel) *entities.DashboardAdminAction {
	return &entities.DashboardAdminAction{
		BoardID:     request.BoardID,
		UserID:      request.UserID,
		InitiatorID: request.InitiatorID,
	}
}
