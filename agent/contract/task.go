package contract

type Task struct {
	ID                 string             `json:"id"`
	Expression         string             `json:"expression"`
	OperationsTimeouts OperationsTimeouts `json:"operationsTimeouts"`
}

type TaskStatus struct {
	Done bool `json:"done"`
	Err  bool `json:"err"`
}

type TaskData struct {
	Task   Task       `json:"task"`
	Status TaskStatus `json:"status"`
}
