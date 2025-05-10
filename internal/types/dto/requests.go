package dto

type PostUsersLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostUsersRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostUsersBoardAdminRequest struct {
	BoardID int `json:"boardId"`
	UserID  int `json:"userId"`
}

// Defaults are from internal/enum.
type GetTasksParams struct {
	Relation string   `schema:"relation,omitempty,default:assigned_to_me"`
	Status   []string `schema:"status,omitempty,default:created,in_progress"`
}

type GetTaskByIDParam struct {
	TaskID int
}

type PostTasksCreateRequest struct {
	Title         string  `json:"title"`
	Description   *string `json:"discription,omitempty"`
	AssignieID    *int    `json:"assignie_id,omitempty"`
	LinkedBoardID *int    `json:"linkedBoard_id,omitempty"`
}

type PostTasksUpdateRequest struct {
	TaskID        int     `json:"taskId"`
	Title         *string `json:"title,omitempty"`
	Status        *string `json:"status,omitempty"`
	Description   *string `json:"discription,omitempty"`
	AssignieID    *int    `json:"assignie_id,omitempty"`
	LinkedBoardID *int    `json:"linkedBoard_id,omitempty"`
}

type PostCommentRequest struct {
	TaskID int    `json:"taskId"`
	Text   string `json:"text"`
}

type GetDashboardByIDParam struct {
	BoardID int
}

type PostDashboardsCreateRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"discription,omitempty"`
}

type PostDashboardsUpdateRequest struct {
	BoardID     int     `json:"boardId"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"discription,omitempty"`
}

type PostDashboardsDeleteRequest struct {
	BoardID int `json:"boardId"`
}
