package contract

type ErrorResponse struct {
	Error error `json:"-"`
}

func (er *ErrorResponse) GetError() error {
	return er.Error
}
