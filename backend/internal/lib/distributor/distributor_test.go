package distributor

import (
	"context"
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"golang.org/x/sync/errgroup"
)

func WithTimeout[R any](f func() R, timeout time.Duration) (R, error) {
	select {
	case <-time.After(timeout):
		var res R
		return res, errors.New("timeout exceeded")
	case res := <-utils.Await(f):
		return res, nil
	}
}

func TestDistributor(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		d := NewDistributor[int](10)
		ctx, cancel := context.WithCancel(context.Background())
		defer d.Close()
		defer cancel()
		d.Start(ctx)

		err, timeoutExceeded := WithTimeout(func() error { return d.Distribute(1) }, time.Second)
		if timeoutExceeded != nil {
			t.Fatal(timeoutExceeded.Error())
		}

		if err != nil {
			t.Fatal(err.Error())
		}

		ch := d.Subscribe(1, 5)

		task, timeoutExceeded := WithTimeout(func() int { return <-ch }, time.Second)
		if timeoutExceeded != nil {
			t.Fatal(timeoutExceeded.Error())
		}

		if task != 1 {
			t.Fatalf("expected task to be 1, got %v", task)
		}

		err = d.Done(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err.Error())
		}

		for i := 1; i <= 6; i++ {
			err, timeoutExceeded := WithTimeout(func() error { return d.Distribute(i) }, time.Second)
			if timeoutExceeded != nil {
				t.Fatal(timeoutExceeded.Error())
			}

			if err != nil {
				t.Fatal(err.Error())
			}
		}

		res, err := WithTimeout(func() []int {
			res := []int{}
			for i := 0; i < 5; i++ {
				res = append(res, <-ch)
			}
			return res
		}, time.Second*5)
		if err != nil {
			t.Fatal(err.Error())
		}

		if !slices.ContainsFunc(res, func(i int) bool {
			return slices.Contains([]int{1, 2, 3, 4, 5}, i)
		}) {
			t.Fatalf("expected to get [1, 2, 3, 4, 5], got %v", res)
		}

		extraTimeCh := make(chan time.Time)
		startTime := time.Now()
		go func() {
			select {
			case <-time.After(time.Second * 10):
				extraTimeCh <- time.Now()
			case <-ch:
				extraTimeCh <- time.Now()
			}
		}()

		time.Sleep(time.Second * 2)
		err = d.Done(1)
		doneTime := time.Now()
		if err != nil {
			t.Fatal(err.Error())
		}

		extraTime := <-extraTimeCh
		if extraTime.Before(doneTime) {
			t.Fatal("expected to get extra after done, got before")
		}

		if extraTime.After(startTime.Add(time.Second * 9)) {
			t.Fatal("expected to get extra")
		}

		for i := 0; i < 5; i++ {
			err = d.Done(1)
			if err != nil {
				t.Fatal(err.Error())
			}
		}

		err = d.Done(1)
		if err == nil {
			t.Fatal("expected to got error")
		}

		g := errgroup.Group{}
		g.Go(func() error {
			closed, err := WithTimeout(func() bool {
				_, ok := <-ch
				return !ok
			}, time.Second*5)
			if !closed {
				return errors.New("expected to close channel")
			}
			return err
		})

		err = d.Unsubscribe(1)
		if err != nil {
			t.Fatal(err.Error())
		}

		err = g.Wait()
		if err != nil {
			t.Fatal(err.Error())
		}
	})
}
