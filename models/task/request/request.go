package request

import "part3/models/task"

type TaskRequest struct {
	Name_Task  string `json:"name_task"`
	Priority   int    `json:"priority"`
	Project_id uint   `json:"project_id"`
}

func (t *TaskRequest) ToTask() task.Task {
	return task.Task{
		Name_Task: t.Name_Task,
		Priority:  t.Priority,
		Project_id: t.Project_id,
	}
}

func (t *TaskRequest) ToTaskCont(name string, pri int) *task.Task {
	return &task.Task{
		Name_Task: name,
		Priority:  pri,
	}
}

func (t *TaskRequest) ToTaskCont1(name string, pri int) *task.Task {
	return &task.Task{
		Name_Task: name,
		Priority:  pri,
	}
}
