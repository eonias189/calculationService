package utils

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/go-chi/jwtauth/v5"
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

func Map[T, R any](s []T, f func(T) R) []R {
	res := make([]R, len(s))
	for i, item := range s {
		res[i] = f(item)
	}
	return res
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

func GetUserIdFromClaims(r *http.Request) (int64, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return 0, err
	}

	userIdAny, ok := claims["user_id"]
	if !ok {
		return 0, errors.ErrInvalidClaims
	}

	userIdFloat64, ok := userIdAny.(float64)
	if !ok {
		return 0, errors.ErrInvalidClaims
	}

	return int64(userIdFloat64), nil
}

func GetIntQuery(query url.Values, key string, deflt int) int {
	val, err := strconv.Atoi(query.Get(key))
	if err != nil {
		return deflt
	}
	return val
}
