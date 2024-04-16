package errors

import "errors"

var (
	ErrExecution                    = errors.New("execution error")
	ErrNotFound                     = errors.New("not found")
	ErrChanClosed                   = errors.New("channel closed")
	ErrMetadataInvalidOrNotProvided = errors.New("metadata is invalid or not provided")
	ErrConnectionDoesntExists       = errors.New("connection doesn`t exists")
	ErrNegativeRunningTasksCount    = errors.New("running negative tasks count")
	ErrInvalidCrendials             = errors.New("invalid crendials")
	ErrInvalidClaims                = errors.New("invalid claims")
)

func ErrMissingEnvParam(param string) error {
	return errors.New("Missing required env param: " + param)
}
