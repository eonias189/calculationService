package use_agents

type GetAgentsResp struct {
	Agents []AgentSource `json:"agents"`
}

type AgentSource struct {
	Id             int64 `json:"id"`
	Ping           int64 `json:"ping"`
	Active         bool  `json:"active"`
	MaxThreads     int   `json:"maxThreads"`
	RunningThreads int   `json:"runningThreads"`
}
