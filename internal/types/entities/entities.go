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

type Task struct {
	ID          int
	Title       string
	Description sql.NullString
	Status      string
	AuthorID    int
	AssignieID  sql.NullInt32
	BoardID     sql.NullInt32
	UpdatedAt   time.Time
}

type Comment struct {
	ID       int
	TaskID   int
	AuthorID int
	Text     string
	DateTime time.Time
}

type Dashboard struct {
	ID          int
	Title       string
	CreatorID   int
	Description sql.NullString
	UpdatedAt   time.Time
}
