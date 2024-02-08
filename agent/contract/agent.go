package contract

type GetAgentStatusResponse struct {
	ErrorResponse
	Status AgentStatus `json:"status"`
}

type AgentStatus struct {
	ThreadsRuning    int `json:"threadsRuning"`
	MaxThreadsNumber int `json:"maxThreadsNumber"`
}
