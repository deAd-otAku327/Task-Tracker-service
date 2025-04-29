package dtocomment

type PostTasksCommentRequest struct {
	TaskID string `json:"taskId"`
	Text   string `json:"text"`
}
