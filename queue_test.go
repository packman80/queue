package queue

import (
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()

	// Test push front and back
	q.PushFront(1)
	q.PushBack(2)
	if q.Length() != 2 {
		t.Errorf("Expected length of 2, got %d", q.Length())
	}

	// Test pop front
	item, ok := q.PopFront()
	if !ok || item != 1 {
		t.Errorf("Expected to pop 1, got %d", item)
	}

	// Test pop back
	item, ok = q.PopBack()
	if !ok || item != 2 {
		t.Errorf("Expected to pop 2, got %d", item)
	}
}

func TestQueueRotate(t *testing.T) {
	q := NewQueue[int]()

	q.PushFront(1)
	q.PushBack(2)

	// Test rotate
	item, ok := q.RotateFrontToBack()
	if !ok || item != 1 {
		t.Errorf("Expected to rotate 1, got %d", item)
	}

	// Should be [2, 1] now
	item, ok = q.PopFront()
	if !ok || item != 2 {
		t.Errorf("Expected to pop 2, got %d", item)
	}
}

func TestQueueConcurrency(t *testing.T) {
	q := NewQueue[int]()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			q.PushBack(val)
		}(i)
	}

	wg.Wait()

	if q.Length() != 100 {
		t.Errorf("Expected length of 100, got %d", q.Length())
	}
}
