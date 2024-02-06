package interfaces

type Task struct {
	ID                 string             `json:"id"`
	Expression         string             `json:"expression"`
	OperationsTimeouts OperationsTimeouts `json:"operationsTimeouts"`
}

type TaskStatus struct {
	ID   string `json:"id"`
	Done bool   `json:"done"`
	Err  bool   `json:"err"`
}
