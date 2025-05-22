package models

type UserLoginModel struct {
	Username string
	Password string
}

type UserRegisterModel struct {
	Email          string
	Username       string
	Password       string
	HashedPassword string // Service calculation.
}

type TaskFilterModel struct {
	Relation *string
	Status   []string

	CreatorID  *int // Service calculation.
	AssignieID *int // Service calculation.
}

type TaskSummaryParamModel struct {
	TaskID int
}

type TaskCreateModel struct {
	Title         string
	Description   *string
	AuthorID      int // Service calculation.
	AssignieID    *int
	LinkedBoardID *int
}

type TaskUpdateModel struct {
	ID            int
	Title         *string
	Status        *string
	Description   *string
	AssignieID    *int
	LinkedBoardID *int

	InitiatorID int // Service calculation.
}

type CommentCreateModel struct {
	TaskID   int
	AuthorID int // Service calculation.
	Text     string
}

type DashboardSummaryParamModel struct {
	BoardID int
}

type DashboardCreateModel struct {
	Title       string
	Description *string

	CreatorID int // Service calculation.
}

type DashboardUpdateModel struct {
	ID          int
	Title       *string
	Description *string

	InitiatorID int // Service calculation.
}

type DashboardDeleteModel struct {
	BoardID int

	InitiatorID int // Service calculation.
}

type DashboardAdminActionModel struct {
	BoardID int
	UserID  int

	InitiatorID int // Service calculation.
}
