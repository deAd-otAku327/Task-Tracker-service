package entities

type UserBoardAdminAction struct {
	BoardID int
	UserID  int

	InitiatorID int
}

type TaskFilter struct {
	Status []string

	CreatorID  *int
	AssignieID *int
}

type TaskUpdate struct {
	ID          int
	Title       *string
	Status      *string
	Discription *string
	AssignieID  *int
	BoardID     *int

	InitiatorID int
}

type DashboardUpdate struct {
	ID          int
	Title       *string
	Discription *string

	InitiatorID int
}

type DashboardDelete struct {
	BoardID int

	InitiatorID int
}
