package interfaces

type IAgent interface {
	GetTaskStatus(id string) (TaskStatus, error)
	IsWorking() bool
}
