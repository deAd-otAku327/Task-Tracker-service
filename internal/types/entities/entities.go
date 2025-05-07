package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type UserBoardAdmin struct {
	BoardID int
	UserID  int
}

type Task struct {
	ID          int
	Title       string
	Discription sql.NullString
	Status      string
	AuthorID    int
	AssignieID  sql.NullInt32
	BoardID     sql.NullInt32
	UpdatedAt   time.Time
}

type TaskFilter struct {
	Relation string
	Status   []string
}

type TaskUpdate struct {
	ID          int
	Title       *string
	Status      *string
	Discription *string
	AssignieID  *int
	BoardID     *int
}

type Comment struct {
	ID       int
	AuthorID int
	Text     string
	DateTime time.Time
}

type Dashboard struct {
	ID          int
	Title       string
	Discription sql.NullString
	UpdatedAt   time.Time
}

type DashboardUpdate struct {
	ID          int
	Title       *string
	Discription *string
}
