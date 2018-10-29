package core

import (
	"sync"
)

// Dag implements IDag
type Dag struct {
	notReadyTaskQueue IQueue
	readyTaskQueue    IQueue
	state             DagState
	stateMutex        sync.Mutex
	nAll              int
	nSent             int
	locked            bool
	lockedMutex       sync.Mutex
}

func (d *Dag) AddTask(task ITask) (err error) {
	d.lockedMutex.Lock()
	defer d.lockedMutex.Unlock()
	if d.locked == true {
		return ErrDagLocked
	}
	d.notReadyTaskQueue.Push(task)
	return nil
}

func (d *Dag) AddEdge(from ITask, to ITask) (err error) {
	d.lockedMutex.Lock()
	defer d.lockedMutex.Unlock()
	if d.locked == true {
		return ErrDagLocked
	}
	to.AddParent(from)
	return nil
}

func (d *Dag) GetReadyTask() ITask {
	return d.readyTaskQueue.Pop().(ITask)
}

func (d *Dag) State() T {
	d.stateMutex.Lock()
	defer d.stateMutex.Unlock()
	return d.state
}

func (d *Dag) Lock() (err error) {
	d.lockedMutex.Lock()
	defer d.lockedMutex.Unlock()
	d.locked = true
	d.nAll = d.notReadyTaskQueue.Size()
	d.nSent = 0
	return nil
}

func (d *Dag) Update(taskChan chan ITask) {
	// update in a goroutine
	go func() {
		for d.State() != DAG_DONE {
			if d.readyTaskQueue.Empty() == false {
				taskChan <- d.readyTaskQueue.Pop().(ITask)
				d.nSent++
				if d.nSent == d.nAll {
					d.state = DAG_DONE
				}
			} else {
				for d.readyTaskQueue.Empty() {
				}
			}
		}
	}()
}
