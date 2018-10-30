package core

import (
	"sync"
)

// Dag implements IDag
type Dag struct {
	readyTaskQueue  IQueue
	notReadyTaskSet ISet
	state           DagState
	stateMutex      sync.Mutex
	nAll            int
	nSent           int
	locked          bool
	lockedMutex     sync.Mutex
}

func (d *Dag) AddTask(task ITask) (err error) {
	d.lockedMutex.Lock()
	defer d.lockedMutex.Unlock()
	if d.locked == true {
		return ErrDagLocked
	}
	if d.notReadyTaskSet.Has(task) {
		return nil
	}
	d.notReadyTaskSet.Insert(task)
	return nil
}

func (d *Dag) AddEdge(from ITask, to ITask) (err error) {
	d.lockedMutex.Lock()
	defer d.lockedMutex.Unlock()
	if d.locked == true {
		return ErrDagLocked
	}
	if d.notReadyTaskSet.Has(from) == false || d.notReadyTaskSet.Has(to) == false {
		return ErrNoSuchTaskInDag
	}
	to.AddParent(from)
	return nil
}

func (d *Dag) GetReadyTask() ITask {
	task := d.readyTaskQueue.Pop()
	if task == nil {
		return nil
	}
	return task.(ITask)
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
	d.nAll = d.notReadyTaskSet.Size()
	d.nSent = 0
	return nil
}

func (d *Dag) Update(taskChan chan ITask) {
	// update in a goroutine
	go func() {
		for d.State() != DAG_DONE {
			if task := d.GetReadyTask(); task != nil {
				taskChan <- task
				d.notReadyTaskSet.Remove(task)
				d.nSent++
				if d.nSent == d.nAll {
					d.state = DAG_DONE
				}
			} else {
				for d.readyTaskQueue.Empty() {
					for _, val := range d.notReadyTaskSet.All() {
						task := val.(ITask)
						if task.IsReady() {
							d.notReadyTaskSet.Remove(task)
							d.readyTaskQueue.Push(task)
						}
					}
				}
			}
		}
	}()
}
