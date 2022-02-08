package task

import (
	"part3/models/task"
	"part3/models/task/request"
	"part3/models/task/response"
)

type Task interface {
	Create(user_id int, newTask task.Task) (task.Task, error)
	// GetById(id int, user_id int) (task.Task, error)
	UpdateById(id int, user_id int, taskReg request.TaskRequest) (task.Task, error)
	// DeleteById(id int, user_id int) (gorm.DeletedAt, error)
	GetAll(user_id int) ([]response.TaskResponse, error)
}
