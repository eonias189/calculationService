package use_timeouts

type GetTimeoutsResp struct {
	Timeouts TimeoutsSource `json:"timeouts"`
}

type TimeoutsSource struct {
	Add *uint `json:"add"`
	Sub *uint `json:"sub"`
	Mul *uint `json:"mul"`
	Div *uint `json:"div"`
}

type PatchTimeoutsBody struct {
	TimeoutsSource
}

type PatchTimeoutsResp struct {
	Timeouts TimeoutsSource `json:"timeouts"`
}
