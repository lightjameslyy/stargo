package core

// IDag interface definition
type IDag interface {
	AddTask(ITask) error
	AddEdge(ITask, ITask) error
	GetReadyTask() ITask
	Lock() error
	Update(chan ITask)
	State() T
}
