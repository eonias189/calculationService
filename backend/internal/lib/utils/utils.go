package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func Await[R any](f func() R) <-chan R {
	out := make(chan R)

	go func() {
		defer close(out)

		out <- f()
	}()

	return out
}

func TryUntilSuccess(ctx context.Context, f func() error, interval time.Duration) {
	err := f()
	for err != nil {
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
			err = f()
		}
	}
}

func HandleError(w http.ResponseWriter, r *http.Request, err error, status int) {
	render.Status(r, status)
	render.JSON(w, r, render.M{"reason": err.Error()})
}

func BindAndValidate(w http.ResponseWriter, r *http.Request, body any) error {
	err := render.DecodeJSON(r.Body, &body)
	if err != nil {
		return err
	}

	err = validator.New().Struct(body)
	return err
}
