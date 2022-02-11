package response

import (
	proResp "part3/models/project/response"
	taskResp "part3/models/task/response"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name     string                  `json:"name"`
	Email    string                  `json:"email"`
	Tasks    []taskResp.TaskResponse `json:"tasks"`
	Projects []proResp.ProResponse  `json:"projects"`
}
