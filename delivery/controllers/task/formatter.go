package task

type GetTaskResponFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type TaskRequest struct {
	Name_Task string `json:"name_task"`
	Priority  int    `json:"priority"`
}
