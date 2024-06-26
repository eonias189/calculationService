package distributor

import (
	"context"
	"sync"
	"time"

	"github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/lib/queue"
)

type Connection[T any] struct {
	MaxTasks     int
	RunningTasks int
	Chan         chan T
}

type Distributor[T any] struct {
	queue *queue.Queue[T]
	conns map[int64]Connection[T]
	mu    *sync.RWMutex
	wg    *sync.WaitGroup
}

func (d *Distributor[T]) GetFreeConn() (int64, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for id, conn := range d.conns {
		if conn.RunningTasks < conn.MaxTasks {
			return id, true
		}
	}
	return 0, false
}

func (d *Distributor[T]) ShiftAndDistribute(id int64) (bool, error) {
	task, ok := d.queue.Shift()
	if !ok {
		return false, nil
	}

	d.mu.RLock()
	conn, ok := d.conns[id]
	d.mu.RUnlock()
	if !ok {
		return false, errors.ErrConnectionDoesntExists
	}

	conn.Chan <- task
	conn.RunningTasks++

	d.mu.Lock()
	d.conns[id] = conn
	d.mu.Unlock()

	return true, nil
}

func (d *Distributor[T]) Distribute(task T) error {
	d.queue.Push(task)
	id, ok := d.GetFreeConn()

	if !ok {
		return nil
	}

	_, err := d.ShiftAndDistribute(id)
	return err
}

func (d *Distributor[T]) Subscribe(id int64, maxTasks int) <-chan T {
	out := make(chan T)
	d.mu.Lock()
	d.conns[id] = Connection[T]{MaxTasks: maxTasks, Chan: out}
	d.mu.Unlock()

	return out
}

func (d *Distributor[T]) Done(id int64) error {
	d.mu.RLock()
	conn, ok := d.conns[id]
	d.mu.RUnlock()
	if !ok {
		return errors.ErrConnectionDoesntExists
	}

	if conn.RunningTasks > 0 {
		conn.RunningTasks--
	}
	d.mu.Lock()
	d.conns[id] = conn
	d.mu.Unlock()

	_, err := d.ShiftAndDistribute(id)
	return err
}

func (d *Distributor[T]) Unsubscribe(id int64) error {
	d.mu.RLock()
	conn, ok := d.conns[id]
	d.mu.RUnlock()
	if !ok {
		return errors.ErrConnectionDoesntExists
	}

	close(conn.Chan)

	d.mu.Lock()
	delete(d.conns, id)
	d.mu.Unlock()

	return nil
}

func (d *Distributor[T]) StartPushing(ctx context.Context, interval time.Duration) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(interval):
				free, found := d.GetFreeConn()
				if !found {
					continue
				}
				d.ShiftAndDistribute(free)
			}
		}
	}()
}

func (d *Distributor[T]) Close() {
	d.wg.Wait()
}

func NewDistributor[T any](workers int) *Distributor[T] {
	return &Distributor[T]{
		queue: queue.NewQueue[T](),
		conns: make(map[int64]Connection[T]),
		mu:    &sync.RWMutex{},
		wg:    &sync.WaitGroup{},
	}
}
