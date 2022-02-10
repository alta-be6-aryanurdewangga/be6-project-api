package request

import "part3/models/task"

type TaskRequest struct {
	Name_Task string `json:"name_task"`
	Priority  int    `json:"priority"`
}

func (t *TaskRequest) ToTask() task.Task {
	return task.Task{
		Name_Task: t.Name_Task,
		Priority:  t.Priority,
	}
}


func (t *TaskRequest) ToTaskCont(name string, pri int) *task.Task {
	return &task.Task{
		Name_Task: name,
		Priority: pri,
	}
}