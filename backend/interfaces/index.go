package interfaces

import (
	"fmt"
)

type None struct{}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (er *ErrorResponse) GetError() error {
	if er.Error == "" {
		return nil
	}
	return fmt.Errorf(er.Error)
}

type ApiMethod[Body any, RP any, Response any] struct {
	Body       Body
	RestParams RP
	Response   *Response
}

func NewApiMethod[Body any, RP any, Response any](body Body, restParams RP, response *Response) ApiMethod[Body, RP, Response] {
	return ApiMethod[Body, RP, Response]{Body: body, RestParams: restParams, Response: response}
}
