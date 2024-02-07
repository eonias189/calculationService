package contract

type AddTaskBody struct {
	Task Task `json:"task"`
}

type GetTasksStatusResponse struct {
	ErrorResponse
	TasksStatus []TaskStatus `json:"tasksStatus"`
}

type GetResultRestParams struct {
	ID string `json:"id"`
}

type GetResultResponse struct {
	ErrorResponse
	Number int `json:"number"`
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
