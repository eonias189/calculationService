package queue

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestQueue(t *testing.T) {
	t.Run("test sync", func(t *testing.T) {
		q := NewQueue[int]()

		_, ok := q.Shift()
		if ok {
			t.Fatal("expected empty queue to shift false")
		}

		q.Push(1)
		val, ok := q.Shift()
		if !ok {
			t.Fatal("expected queue to shift true")
		}
		if val != 1 {
			t.Fatalf("expected queue to shift 1, got %v", val)
		}

		q.Push(2)
		q.Push(3)
		q.Push(4)
		val, ok = q.Shift()
		if !ok {
			t.Fatal("expected queue to shift true")
		}
		if val != 2 {
			t.Fatalf("expected queue to shift 2, got %v", val)
		}
		val, ok = q.Shift()
		if !ok {
			t.Fatal("expected queue to shift true")
		}
		if val != 3 {
			t.Fatalf("expected queue to shift 3, got %v", val)
		}
		val, ok = q.Shift()
		if !ok {
			t.Fatal("expected queue to shift true")
		}
		if val != 4 {
			t.Fatalf("expected queue to shift 4, got %v", val)
		}
	})

	t.Run("test parallel", func(t *testing.T) {
		q := NewQueue[int]()
		g, ctx := errgroup.WithContext(context.Background())

		g.Go(func() error {
			for i := 0; i < 1000; i++ {
				select {
				case <-ctx.Done():
					return nil
				default:
					q.Push(i)
				}
			}
			return nil
		})

		time.Sleep(time.Millisecond * 2)

		g.Go(func() error {
			last, ok := q.Shift()
			if !ok {
				return fmt.Errorf("expected queue to shift true")
			}

			for i := 1; i < 1000; i++ {
				val, ok := q.Shift()
				if !ok {
					return fmt.Errorf("expected queue to shift true")
				}
				if val != last+1 {
					return fmt.Errorf("expected val to be %v, got %v", last+1, val)
				}
				last++
			}
			return nil
		})

		err := g.Wait()
		if err != nil {
			t.Fatal(err.Error())
		}
	})
}
