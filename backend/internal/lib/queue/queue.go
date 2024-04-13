package queue

import (
	"sync"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Last  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	back *Node[T]
	mu   *sync.Mutex
}

func (q *Queue[T]) Push(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	newNode := &Node[T]{Value: item}

	if q.head == nil && q.back == nil {
		q.head = newNode
		q.back = newNode
	}

	last := q.back
	newNode.Next = last
	last.Last = newNode
	q.back = newNode
}

func (q *Queue[T]) Shift() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.head == nil {
		var res T
		return res, false
	}

	head := q.head
	if q.back == head {
		q.back = nil
		q.head = nil
	} else {
		last := head.Last
		last.Next = nil
		q.head = last
	}

	return head.Value, true
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		mu: &sync.Mutex{},
	}
}
