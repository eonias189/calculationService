package errors

import "errors"

var (
	ErrNotFound                     = errors.New("not found")
	ErrChanClosed                   = errors.New("channel closed")
	ErrMetadataInvalidOrNotProvided = errors.New("metadata is invalid or not provided")
	ErrConnectionDoesntExists       = errors.New("connection doesn`t exists")
	ErrNegativeRunningTasksCount    = errors.New("running negative tasks count")
)

func ErrMissingEnvParam(param string) error {
	return errors.New("Missing required env param: " + param)
}
