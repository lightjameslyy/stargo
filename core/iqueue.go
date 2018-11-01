package core

// IQueue interface definition
type IQueue interface {
	// push an element to the tail of the queue.
	Push(T)
	// return the element at the head of the queue and delete it.
	Pop() T
	// return if the queue is empty.
	Empty() bool
	// return size of the queue.
	Size() int
}
