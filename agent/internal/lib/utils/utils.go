package utils

import (
	"context"
	"time"
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
