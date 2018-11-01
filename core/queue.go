package core

import "sync"

// A thread-safe FIFO queue implements IQueue
type Queue struct {
	items []T
	mutex sync.Mutex
}

// Push an element into the queue.
// Here we assuming there is no limit on the queue's capacity.
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
// NOT thread safe.
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

// Return queue size.
// NOT thread safe.
func (q *Queue) Size() int {
	return len(q.items)
}
