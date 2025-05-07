package models

import "time"

type UserModel struct {
	ID             int
	Username       string
	Email          string
	HashedPassword string
}

type TaskModel struct {
	ID          int
	Title       string
	Discription *string
	Status      string
	AssignieID  *int
	BoardID     *int
	UpdatedAt   time.Time
}

type CommentModel struct {
	ID         int
	AuthorID   int
	AuthorName string
	Text       string
	DateTime   time.Time
}

type DashboardModel struct {
	ID          int
	Title       string
	Discription *string
	UpdatedAt   time.Time
}
