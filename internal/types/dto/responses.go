package dto

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUsersResponse []*UserResponse

type TaskResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"discription,omitempty"`
	Status      string  `json:"status"`
	Assignie    *bool   `json:"assignie,omitempty"`
	Board       *bool   `json:"board,omitempty"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetTasksResponse []*TaskResponse

type GetTaskByIDResponse struct {
	Task        *TaskResponse      `json:"task"`
	Comments    []*CommentResponse `json:"comments,omitempty"`
	Author      *UserResponse      `json:"author"`
	Assignie    *UserResponse      `json:"assignie,omitempty"`
	LinkedBoard *BoardData         `json:"linkedBoard,omitempty"`
}

type BoardData struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type CommentResponse struct {
	ID       int    `json:"id,omitempty"`
	AuthorID int    `json:"authorId"`
	Text     string `json:"text"`
	DateTime string `json:"dateTime"`
}

type DashboardResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"discription,omitempty"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetDashboardsResponse []*DashboardResponse

type GetDashboardByIDResponse struct {
	Dashboard *DashboardResponse `json:"dashboard"`
	Tasks     []*TaskResponse    `json:"tasks"`
	Admins    []*UserResponse    `json:"admin"`
}

type Token string

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
