package request

import "part3/models/task"

type TaskRequest struct {
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	Project_id uint   `json:"project_id"`
}

func (t *TaskRequest) ToTask() task.Task {
	return task.Task{
		Name:       t.Name,
		Priority:   t.Priority,
		Project_id: t.Project_id,
	}
}
