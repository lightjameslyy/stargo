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
	Parents ISet
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
	res = t.Func(t.Args)
	t.state = TASK_DONE
	return res, nil
}

// Get state. Check if it's type is TaskState.
func (t *Task) State() T {
	return t.state
}

// Add a parent task.
func (t *Task) AddParent(p ITask) {
	t.Parents.Insert(p)
}

// Get number of parents.
func (t *Task) ParentsSize() int {
	return t.Parents.Size()
}
