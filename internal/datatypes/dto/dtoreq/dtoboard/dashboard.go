package dtoboard

type PostDashboardsCreateRequest struct {
	Title       string  `json:"title"`
	Discription *string `json:"discription,omitempty"`
}

type PostDashboardsUpdateRequest struct {
	BoardId     string  `json:"boardId"`
	Title       *string `json:"title,omitempty"`
	Discription *string `json:"discription,omitempty"`
}

type PostDashboardsDeleteRequest struct {
	BoardId string `json:"boardId"`
}

type PostDashboardsAddAdminRequest struct {
	BoardId string `json:"boardId"`
	UserId  string `json:"userId"`
}

type PostDashboardsDeleteAdminRequest struct {
	BoardId string `json:"boardId"`
	UserId  string `json:"userId"`
}
