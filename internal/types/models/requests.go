package models

type UserLoginModel struct {
	Username string
	Password string
}

type UserRegisterModel struct {
	Email    string
	Username string
	Password string
}

type UserBoardAdminActionModel struct {
	BoardID int
	UserID  int

	InitiatorID int
}

type TaskFilterModel struct {
	Relation string
	Status   []string

	CreatorID  *int
	AssignieID *int
}

type TaskIDParamModel struct {
	TaskID int
}

type TaskCreateModel struct {
	Title         string
	Description   *string
	AuthorID      int
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

	InitiatorID int
}

type CommentCreateModel struct {
	TaskID   int
	AuthorID int
	Text     string
}

type DashboardIDParamModel struct {
	BoardID int
}

type DashboardCreateModel struct {
	Title       string
	Description *string

	CreatorID int
}

type DashboardUpdateModel struct {
	ID          int
	Title       *string
	Description *string

	InitiatorID int
}

type DashboardDeleteModel struct {
	BoardID int

	InitiatorID int
}
