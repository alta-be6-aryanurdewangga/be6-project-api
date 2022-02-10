package response

import "time"

type TaskResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	Project_id   int    `json:"project_id"`
	Project_name string `json:"project_name"`
}
