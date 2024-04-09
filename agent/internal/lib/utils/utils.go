package utils

func Await[R any](f func() R) <-chan R {
	out := make(chan R)

	go func() {
		defer close(out)
		out <- f()
	}()

	return out
}
