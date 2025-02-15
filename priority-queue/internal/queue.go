package internal

import "sync"

type Queue[T any] struct {
	items []T
	mu    sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{make([]T, 0), sync.Mutex{}}
}

func (q *Queue[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (*T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return &item, true
}
