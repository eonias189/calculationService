package interfaces

import "fmt"

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

type ApiMethod[Body any, RestParams any, Response any] struct {
	Body       Body
	RestParams RestParams
	Response   *Response
}

func NewApiMethod[Body any, RestParams any, Response any](body Body, restParams RestParams, response *Response) ApiMethod[Body, RestParams, Response] {
	return ApiMethod[Body, RestParams, Response]{Body: body, RestParams: restParams, Response: response}
}

func f[B any, RP any, R any](m ApiMethod[B, RP, R]) ApiMethod[B, RP, R] {
	return ApiMethod[B, RP, R]{
		Body:       m.Body,
		RestParams: m.RestParams,
		Response:   m.Response,
	}

}
