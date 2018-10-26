package core

// IDag interface definition
type IDag interface {
	AddTask(ITask)
	AddEdge(ITask, ITask)
	GetReadyTask() ITask
	State() T
}
