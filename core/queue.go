package core

import "sync"

// An FIFO queue.
type Queue struct {
	items []interface{}
	mutex *sync.Mutex
}

// Push an element into the queue.
// e.g.:
// 		q.push()
func (q *Queue) Push(v interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, v)
}

// Pop a element from head.
func (q *Queue) Pop() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.IsEmpty() {
		return nil
	}
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
