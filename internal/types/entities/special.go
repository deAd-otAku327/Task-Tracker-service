package entities

type TaskFilter struct {
	Status []string

	CreatorID  *int
	AssignieID *int
}

type TaskUpdate struct {
	ID          int
	Title       *string
	Status      *string
	Description *string
	AssignieID  *int
	BoardID     *int

	InitiatorID int
}

type DashboardUpdate struct {
	ID          int
	Title       *string
	Description *string

	InitiatorID int
}

type DashboardDelete struct {
	BoardID int

	InitiatorID int
}

type DashboardAdminAction struct {
	BoardID int
	UserID  int

	InitiatorID int
}
