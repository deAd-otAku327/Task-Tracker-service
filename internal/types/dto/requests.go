package dto

import (
	"strings"
)

type PostUsersLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostUsersRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetTasksParams struct {
	Relation *string     `schema:"relation,omitempty"`
	Status   *StatusPack `schema:"status,omitempty"`
}

type StatusPack struct {
	Statuses []string
}

// Implementation of gorilla/schema interface.
func (sp *StatusPack) UnmarshalText(text []byte) error {
	parsed := strings.Split(string(text), ",")
	for i := 0; i < len(parsed); i++ {
		parsed[i] = strings.TrimSpace(parsed[i])
	}

	sp.Statuses = parsed
	return nil
}

type GetTaskSummaryParam struct {
	TaskID int
}

type PostTasksCreateRequest struct {
	Title         string  `json:"title"`
	Description   *string `json:"description,omitempty"`
	AssignieID    *int    `json:"assignie_id,omitempty"`
	LinkedBoardID *int    `json:"linkedBoard_id,omitempty"`
}

type PostTasksUpdateRequest struct {
	TaskID        int     `json:"taskId"`
	Title         *string `json:"title,omitempty"`
	Status        *string `json:"status,omitempty"`
	Description   *string `json:"description,omitempty"`
	AssignieID    *int    `json:"assignie_id,omitempty"`
	LinkedBoardID *int    `json:"linkedBoard_id,omitempty"`
}

type PostCommentRequest struct {
	TaskID int    `json:"taskId"`
	Text   string `json:"text"`
}

type GetDashboardSummaryParam struct {
	BoardID int
}

type PostDashboardsCreateRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

type PostDashboardsUpdateRequest struct {
	BoardID     int     `json:"boardId"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type PostDashboardsDeleteRequest struct {
	BoardID int `json:"boardId"`
}

type PostDashboardsAdminRequest struct {
	BoardID int `json:"boardId"`
	UserID  int `json:"userId"`
}
