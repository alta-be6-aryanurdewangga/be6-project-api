package task

import (
	"part3/models/task/response"

	"gorm.io/gorm"
)



type Task struct {
	gorm.Model

	User_ID   uint   
	Name_Task string `gorm:"not null;type:varchar(100)"`
	Priority  int    `gorm:"not null;indext;type:int"`
}

func (t *Task) ToTaskResponse() response.TaskResponse {
	return response.TaskResponse{
		ID:        t.ID,
		Name_Task: t.Name_Task,
		Priority:  t.Priority,
	}
}
