package responses

import "time"

type Comment struct {
	ID         string    `json:"id,omitempty"`
	AuthorName string    `json:"author"`
	Text       string    `json:"text"`
	DateTime   time.Time `json:"dateTime"`
}
