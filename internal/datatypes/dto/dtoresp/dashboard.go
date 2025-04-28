package responses

import "time"

type Dashboard struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Discription string    `json:"discription"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetDashboardsResponse []*Dashboard

type GetDashboardByIDResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
	Tasks     []*Task    `json:"tasks"`
	Admins    []*User    `json:"admin"`
}
