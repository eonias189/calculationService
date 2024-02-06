package interfaces

type IOrchestrator interface {
	AddTask(Task) error
	GetTasksStatus() ([]TaskStatus, error)
	GetResult(string) (int, error)
	GetOperationsTimeouts() (OperationsTimeouts, error)
	SetOperationsTimeouts(OperationsTimeouts) error
	GetTask() (Task, error)
	SetResult(string, int) error
	Register(string) error
	Run(string)
}

type AddTaskApi ApiMethod[AddTaskBody, None, ErrorResponse]
type GetTasksStatusApi ApiMethod[None, None, GetTasksStatusResponse]
type GetResultApi ApiMethod[None, GetResultRestParams, GetResultResponse]
type GetOperationsTimeoutsApi ApiMethod[None, None, GetOperationsTimeoutsResponse]
type SetOperationsTimeoutsApi ApiMethod[SetOperationsTimeoutsBody, None, ErrorResponse]
type GetTaskApi ApiMethod[None, None, GetTaskResponse]
type SetResultApi ApiMethod[SetResultBody, None, ErrorResponse]
type RegisterApi ApiMethod[RegisterBody, None, ErrorResponse]

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
