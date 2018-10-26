package core

// ITask interface definition
type ITask interface {
	SetFunc(F)
	SetArgs(T)
	State() T
	AddParent(ITask)
	ParentsSize() int
	Process() (T, error)
}
