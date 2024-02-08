package contract

type AddTaskBody struct {
	Task Task `json:"task"`
}

type GetTasksDataResponse struct {
	ErrorResponse
	TasksData []TaskData `json:"data"`
}

type GetAgentsDataResponse struct {
	ErrorResponse
	Data []AgentData `json:"data"`
}

type AgentData struct {
	Ping int `json:"ping"`
	AgentStatus
}

type GetOperationsTimeoutsResponse struct {
	ErrorResponse
	OperationsTimeouts OperationsTimeouts `json:"operationsTimeouts"`
}

type SetOperationsTimeoutsBody struct {
	OperationsTimeouts OperationsTimeouts `json:"operationsTimeouts"`
}

type GetTaskResponse struct {
	ErrorResponse
	Task Task `json:"task"`
}
type SetResultBody struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
}

type RegisterBody struct {
	Url string `json:"url"`
}
