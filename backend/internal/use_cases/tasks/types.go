package use_tasks

type PostTaskBody struct {
	Expression string `json:"expression" validate:"required"`
}

type PostTaskResp struct {
	Task TaskSource `json:"task"`
}

type TaskSource struct {
	Id         int64   `json:"id"`
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
	Status     string  `json:"status"`
	CreateTime int64   `json:"createTime"`
}

type GetTasksResp struct {
	Tasks []TaskSource `json:"tasks"`
}

type GetTaskResp struct {
	Task TaskSource `json:"task"`
}
