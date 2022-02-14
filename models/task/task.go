package task

import (
	"part3/models/task/response"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	User_ID    uint
	Name       string `gorm:"not null;type:varchar(100)"`
	Status     bool
	Priority   int  `gorm:"not null;index;type:int"`
	Project_id uint `gorm:"not null"`
}

func (t *Task) ToTaskResponse() response.TaskResponse {
	return response.TaskResponse{
		ID:         t.ID,
		Name:       t.Name,
		Status:     t.Status,
		Priority:   t.Priority,
		Project_id: int(t.Project_id),
	}
}
