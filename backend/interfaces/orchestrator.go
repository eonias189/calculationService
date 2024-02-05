package interfaces

type IOrchestrator interface {
	AddTask(Task) error
	GetTasksStatus() ([]TaskStatus, error)
	GetResult(string) (int, error)
	GetOperationsTimeouts() (OperationsTimeouts, error)
	GetTask() (Task, error)
	SetResult(string, int) error
	Register(string)
}
