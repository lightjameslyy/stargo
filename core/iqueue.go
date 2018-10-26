package core

// IQueue interface definition
type IQueue interface {
	Push(T)
	Pop() T
	Empty() bool
	Size() int
}
