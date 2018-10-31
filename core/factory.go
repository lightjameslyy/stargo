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
		parents: SetFactory{}.Create(),
		state:   TASK_NOTDONE,
	}
}

// factory for set
type ISetFactory interface {
	Create() ISet
}

type SetFactory struct{}

func (SetFactory) Create() ISet {
	return &Set{
		mp: map[T]bool{},
	}
}

// factory for queue
type IQueueFactory interface {
	Create() IQueue
}

type QueueFactory struct{}

func (QueueFactory) Create() IQueue {
	return &Queue{}
}

//factory for dag
type IDagFactory interface {
	Create() IDag
}

type DagFactory struct{}

func (DagFactory) Create() IDag {
	return &Dag{
		readyTaskQueue:  QueueFactory{}.Create(),
		notReadyTaskSet: SetFactory{}.Create(),
		state:           DAG_INIT,
	}
}

// factory for pool
type IPoolFactory interface {
	Create() IPool
}

type PoolFactory struct{}

func (PoolFactory) Create() IPool {
	return &Pool{
		workers: 0,
	}
}
