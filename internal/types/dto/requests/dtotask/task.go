package dtotask

// Defaults are from internal/enum.
type GetTasksParams struct {
	Relation string   `schema:"relation,omitempty,default:assigned_to_me"`
	Status   []string `schema:"status,omitempty,default:created,in_progress"`
}

type GetTaskByIDParam string

type PostTasksCreateRequest struct {
	Title       string  `json:"title"`
	Discription *string `json:"discription,omitempty"`
	LinkedBoard *string `json:"linkedBoard,omitempty"`
}

type PostTasksUpdateRequest struct {
	TaskId      string  `json:"taskId"`
	Title       *string `json:"title,omitempty"`
	Status      *string `json:"status,omitempty"`
	Discription *string `json:"discription,omitempty"`
	Assignie    *string `json:"assignie,omitempty"`
	LinkedBoard *string `json:"linkedBoard,omitempty"`
}
