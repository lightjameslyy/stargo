package core

import (
	"fmt"
	"reflect"
	"runtime"
)

// Task implements ITask
type Task struct {
	Func    F
	Args    T
	parents ISet
	state   TaskState
}

// Set function.
func (t *Task) SetFunc(f F) {
	t.Func = f
}

// Set arguments.
func (t *Task) SetArgs(args T) {
	t.Args = args
}

// Run the task.
func (t *Task) Process() (res T, err error) {
	defer func() {
		if r := recover(); r != nil {
			t.state = TASK_WRONG
			res = nil
			err = fmt.Errorf("fatal error: panic at %s, args: %+v",
				runtime.FuncForPC(reflect.ValueOf(t.Func).Pointer()).Name(), t.Args)
		}
	}()
	switch t.state {
	case TASK_NOTDONE:
		if t.Func == nil {
			t.state = TASK_DONE
			break
		}
		res = t.Func(t.Args)
		t.state = TASK_DONE
		return res, nil
	case TASK_WRONG:
		t.state = TASK_WRONG
		return nil, fmt.Errorf("task is in TASK_WRONG state, " +
			"this is because the former tasks broke down or this task is" +
			" already processed to be wrong.")
	}
	return nil, nil
}

// Get state. Check if it's type is TaskState.
func (t *Task) State() T {
	return t.state
}

// Add a parent task.
func (t *Task) AddParent(p ITask) {
	t.parents.Insert(p)
}

func (t *Task) IsReady() bool {
	for _, parent := range t.parents.All() {
		task := parent.(ITask)
		switch task.State() {
		case TASK_DONE:
			t.parents.Remove(task)
		case TASK_WRONG:
			t.state = TASK_WRONG
			return true
		}
	}
	return t.ParentsSize() == 0
}

// Get number of parents.
func (t *Task) ParentsSize() int {
	return t.parents.Size()
}
