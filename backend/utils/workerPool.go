package utils

import "sync"

type Task interface {
	Do()
}

type WorkerPool struct {
	ExecutingWorkers int
	MaxWorkers       int
	tasks            chan Task
	wg               sync.WaitGroup
	exit             chan struct{}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.MaxWorkers; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case <-wp.exit:
					return
				case task := <-wp.tasks:
					wp.ExecutingWorkers++
					task.Do()
					wp.ExecutingWorkers--

				}
			}
		}()
	}
}

func (wp *WorkerPool) Close() {
	close(wp.exit)
	wp.wg.Wait()
}

func (wp *WorkerPool) AddTask(t Task) {
	wp.tasks <- t
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{MaxWorkers: maxWorkers, wg: sync.WaitGroup{}, tasks: make(chan Task), exit: make(chan struct{})}
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
