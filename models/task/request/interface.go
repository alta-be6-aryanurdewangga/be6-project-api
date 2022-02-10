package request

import "part3/models/task"

type TaskReq interface {
	ToTaskCont(name string, pri int) *task.Task
}
