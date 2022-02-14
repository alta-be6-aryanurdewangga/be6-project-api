package task

import (
	"part3/models/task"
	"part3/models/task/request"
	"part3/models/task/response"

	"gorm.io/gorm"
)

type Task interface {
	Create(user_id int, newTask task.Task) (task.Task, error)
	UpdateById(id int, user_id int, taskReg request.TaskRequest) (task.Task, error)
	DeleteById(id int, user_id int) (gorm.DeletedAt, error)
	GetAll(user_id int) ([]response.TaskResponse, error)
	GetByIdResp(id int, user_id int) (response.TaskResponse, error)
	TaskCompleted(id int, user_id int, taskRequest request.TaskRequest) (task.Task, error)
	TaskReopened(id int, user_id int, taskRequest request.TaskRequest) (task.Task, error)
}
