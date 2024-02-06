package interfaces

type IAgent interface {
	GetTaskStatus(id string) (TaskStatus, error)
	Run(string)
}

type GetTaskStatusApi ApiMethod[None, GetTaskStatusRestParams, GetTaskStatusResponse]

type GetTaskStatusResponse struct {
	ErrorResponse
	TaskStatus TaskStatus `json:"taskStatus"`
}

type GetTaskStatusRestParams struct {
	ID string `json:"id"`
}
