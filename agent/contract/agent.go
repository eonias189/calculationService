package contract

type PingResponse struct {
	ErrorResponse
	Ok bool `json:"ok"`
}
