package dto

import (
	"time"
)

type UserResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUsersResponse []*UserResponse

type TaskResponse struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Discription *string   `json:"discription,omitempty"`
	Status      string    `json:"status"`
	Assignie    *bool     `json:"assignie,omitempty"`
	Board       *bool     `json:"board,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetTasksResponse []*TaskResponse

type GetTaskByIDResponse struct {
	Task        *TaskResponse      `json:"task"`
	Comments    []*CommentResponse `json:"comments,omitempty"`
	Author      *UserResponse      `json:"author"`
	Assignie    *UserResponse      `json:"assignie,omitempty"`
	LinkedBoard struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"linkedBoard"`
}

type CommentResponse struct {
	ID         string    `json:"id,omitempty"`
	AuthorName string    `json:"author"`
	Text       string    `json:"text"`
	DateTime   time.Time `json:"dateTime"`
}

type DashboardResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Discription string    `json:"discription"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetDashboardsResponse []*DashboardResponse

type GetDashboardByIDResponse struct {
	Dashboard *DashboardResponse `json:"dashboard"`
	Tasks     []*TaskResponse    `json:"tasks"`
	Admins    []*UserResponse    `json:"admin"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
