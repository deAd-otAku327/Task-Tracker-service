package responses

import (
	"task-tracker-service/internal/enum"
	"time"
)

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
