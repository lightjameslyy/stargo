package core

import (
	"fmt"
	"reflect"
	"runtime"
)

type T = interface{}

type F = func(T) T

type State int

const (
	NOTDONE State = iota
	DONE
	WRONG
)

type ITask interface {
	Process() (T, error)
	State() T
}

type Task struct {
	Func    F
	Args    T
	Parents ISet
	state   State
}

func (t *Task) Process() (res T, err error) {
	defer func() {
		if r := recover(); r != nil {
			t.state = WRONG
			res = nil
			err = fmt.Errorf("fatal error: panic at %s, args: %+v", runtime.FuncForPC(reflect.ValueOf(t.Func).Pointer()).Name(), t.Args)
		}
	}()
	res = t.Func(t.Args)
	t.state = DONE
	return res, nil
}

func (t *Task) State() T {
	return t.state
}
