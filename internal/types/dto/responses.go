package dto

import (
	"task-tracker-service/internal/types/enum"
	"time"
)

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUsersResponse []*User

type Task struct {
	ID          int32           `json:"id"`
	Title       string          `json:"title"`
	Status      enum.TaskStatus `json:"status"`
	Discription *string         `json:"discription,omitempty"`
	Assignie    *bool           `json:"assignie,omitempty"`
	Board       *bool           `json:"board,omitempty"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type GetTasksResponse []*Task

type GetTaskByIDResponse struct {
	Task        *Task      `json:"task"`
	Comments    []*Comment `json:"comments,omitempty"`
	Author      *User      `json:"author"`
	Assignie    *User      `json:"assignie,omitempty"`
	LinkedBoard struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"linkedBoard"`
}

type Comment struct {
	ID         string    `json:"id,omitempty"`
	AuthorName string    `json:"author"`
	Text       string    `json:"text"`
	DateTime   time.Time `json:"dateTime"`
}

type Dashboard struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Discription string    `json:"discription"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetDashboardsResponse []*Dashboard

type GetDashboardByIDResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
	Tasks     []*Task    `json:"tasks"`
	Admins    []*User    `json:"admin"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
