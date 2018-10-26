package core

import "sync"

// Dag implements IDag
type Dag struct {
	notReadyTaskQueue IQueue
	readyTaskQueue    IQueue
	state             DagState
	stateMutex        sync.Mutex
	nAll              int
	nDone             int
}

func (d *Dag) AddTask(task ITask) {
	d.notReadyTaskQueue.Push(task)
}

func (d *Dag) AddEdge(from ITask, to ITask) {
	to.AddParent(from)
}

func (d *Dag) GetReadyTask() ITask {
	return d.readyTaskQueue.Pop().(ITask)
}

func (d *Dag) State() T {
	d.stateMutex.Lock()
	defer d.stateMutex.Unlock()
	return d.state
}

func (d *Dag) lock() {

}

func (d *Dag) update() {

}
