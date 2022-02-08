package task

type TaskRequest struct {
	Name_Task string `json:"name_task"`
	Priority  int    `json:"priority"`
}

type GetTaskResponFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
