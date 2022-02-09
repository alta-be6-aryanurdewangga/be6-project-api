package response

type TaskResponse struct {
	ID        uint   `json:"id"`
	Name_Task string `json:"name_task"`
	Priority  int    `json:"priority"`
}
