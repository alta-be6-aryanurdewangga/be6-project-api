package response

import "time"

type TaskResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name_Task string `json:"name_task"`
	Priority  int    `json:"priority"`
}
