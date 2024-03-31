package pool

import (
	"context"
	"sync"
)

type Task interface {
	Do()
}

type WorkerPool interface {
	Start(context.Context)
	AddTask(Task)
	Close()
}

type workerPoolImpl struct {
	MaxWorkers int
	tasks      chan Task
	wg         sync.WaitGroup
}

func (wp *workerPoolImpl) Start(ctx context.Context) {
	for i := 0; i < wp.MaxWorkers; i++ {
		wp.wg.Add(1)

		go func() {
			defer wp.wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case task := <-wp.tasks:
					task.Do()
				}
			}
		}()
	}
}

func (wp *workerPoolImpl) AddTask(t Task) {
	wp.tasks <- t
}

func (wp *workerPoolImpl) Close() {
	wp.wg.Wait()
}

func NewWorkerPool(maxWorkers int) WorkerPool {
	return &workerPoolImpl{MaxWorkers: maxWorkers, wg: sync.WaitGroup{}, tasks: make(chan Task)}
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
