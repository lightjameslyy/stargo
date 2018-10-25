package core

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type T = interface{}

type Func = func([]T) T

type ITask interface {
	Process() (T, error)
	State() T
}

type Task struct {
	Command Func
	Args    []T
}

func (t *Task) Run() T {
	return t.Command(t.Args)
}

func (t *Task) TaskInfo() string {
	var res strings.Builder
	res.WriteString(runtime.FuncForPC(reflect.ValueOf(t.Command).Pointer()).Name() + "( ")
	for _, arg := range t.Args {
		res.WriteString(fmt.Sprintf("%T: %v, ", arg, arg))
	}
	res.WriteString(")\n")
	return res.String()
}

type F = func(T) T

type State int

const (
	NOTDONE State = iota
	DONE
	WRONG
)

type STask struct {
	Func    F
	Args    T
	Parents ISet
	state   State
}

func (t *STask) Process() (res T, err error) {
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

func (t *STask) State() T {
	return t.state
}
