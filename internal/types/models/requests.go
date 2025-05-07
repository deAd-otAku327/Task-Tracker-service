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

type UserBoardAdminModel struct {
	BoardID int
	UserID  int
}

type TaskFilterModel struct {
	Relation string
	Status   []string
}

type TaskIDParamModel struct {
	TaskID int
}

type TaskCreateModel struct {
	Title         string
	Discription   *string
	AuthorID      int
	AssignieID    *int
	LinkedBoardID *int
}

type TaskUpdateModel struct {
	ID            int
	Title         *string
	Status        *string
	Discription   *string
	AssignieID    *int
	LinkedBoardID *int
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
	Discription *string
}

type DashboardUpdateModel struct {
	ID          int
	Title       *string
	Discription *string
}

type DashboardDeleteModel struct {
	BoardID int
}
