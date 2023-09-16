package queue

import "sync"

type Queue[T any] struct {
	mu    sync.Mutex
	items []T
}

func NewQueue() *Queue[T] {
	return &Queue{}
}

func (q *Queue[T]) PushFront(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append([]T{item}, q.items...)
}

func (q *Queue[T]) PushBack(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue[T]) PopFront() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) PopBack() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item, true
}

func (q *Queue[T]) RotateFrontToBack() (T, bool) {
	item, ok := q.PopFront()
	if !ok {
		return nil, false
	}
	q.PushBack(item)
	return item, true
}

func (q *Queue[T]) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
