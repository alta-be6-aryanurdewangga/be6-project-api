package task

type GetTaskResponFormat struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Name_Task string `json:"name_task"`
}
