package contract

import (
	"fmt"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func (er *ErrorResponse) GetError() error {
	if er.Error == "" {
		return nil
	}
	return fmt.Errorf(er.Error)
}
