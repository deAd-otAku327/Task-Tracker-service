package models

import "time"

type UserModel struct {
	ID             int
	Username       string
	Email          string
	HashedPassword string
}

type UserListModel []*UserModel

type TaskModel struct {
	ID          int
	Title       string
	Description *string
	Status      string
	AssignieID  *int
	BoardID     *int
	UpdatedAt   time.Time
}

type TaskListModel []*TaskModel

type CommentModel struct {
	ID       int
	AuthorID int
	Text     string
	DateTime time.Time
}

type TaskSummaryModel struct {
	Task        *TaskModel
	Comments    []*CommentModel
	Author      *UserModel
	Assignie    *UserModel
	LinkedBoard *DashboardModel
}

type DashboardModel struct {
	ID          int
	Title       string
	Description *string
	UpdatedAt   time.Time
}

type DashboardListModel []*DashboardModel

type DashboardSummaryModel struct {
	Dashboard *DashboardModel
	Tasks     []*TaskModel
	Admins    []*UserModel
}
