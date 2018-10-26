package core

// factory interface definition
type Factory interface {
	Create() T
}

// factory for tasks
type ITaskFactory interface {
	Create() ITask
}

type TaskFactory struct{}

func (TaskFactory) Create() ITask {
	return &Task{
		Parents: SetFactory{}.Create(),
		state:   TASK_NOTDONE}
}

// factory for set
type ISetFactory interface {
	Create() ISet
}

type SetFactory struct{}

func (SetFactory) Create() ISet {
	return &Set{mp: map[T]bool{}}
}

// factory for queue
type IQueueFactory interface {
	Create() IQueue
}

type QueueFactory struct{}

func (QueueFactory) Create() IQueue {
	return &Queue{}
}
