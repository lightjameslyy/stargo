package core

import "sync"

type IQueue interface {
	Push(T)
	Pop() T
	Empty() bool
	Size() int
}

// A thread-safe FIFO queue.
type Queue struct {
	items []T
	mutex sync.Mutex
}

// Push an element into the queue.
func (q *Queue) Push(v T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, v)
}

// Pop a element from head.
// If the queue is empty, return nil.
func (q *Queue) Pop() T {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Empty() {
		return nil
	}
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

// Returns if the queue is empty.
// Not concurrently consistent.
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

// Return queue size.
// Not concurrently consistent.
func (q *Queue) Size() int {
	return len(q.items)
}
