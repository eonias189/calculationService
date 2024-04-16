package pool

import (
	"context"
	"sync"
)

type Task interface {
	Do()
}

type WorkerPool struct {
	MaxWorkers     int
	tasks          chan Task
	wg             sync.WaitGroup
	executingTasks int
	mu             *sync.RWMutex
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.MaxWorkers; i++ {
		wp.wg.Add(1)

		go func() {
			defer wp.wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case task := <-wp.tasks:
					wp.mu.Lock()
					wp.executingTasks++
					wp.mu.Unlock()
					task.Do()
					wp.mu.Lock()
					wp.executingTasks--
					wp.mu.Unlock()
				}
			}
		}()
	}
}

func (wp *WorkerPool) AddTask(t Task) {
	wp.tasks <- t
}

func (wp *WorkerPool) Close() {
	wp.wg.Wait()
}

func (wp *WorkerPool) ExecutingTasks() int {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	return wp.executingTasks
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{MaxWorkers: maxWorkers, wg: sync.WaitGroup{}, tasks: make(chan Task), mu: &sync.RWMutex{}}
}

type simpleTask struct {
	f func()
}

func (st *simpleTask) Do() {
	st.f()
}

func NewTask(f func()) Task {
	return &simpleTask{f: f}
}
