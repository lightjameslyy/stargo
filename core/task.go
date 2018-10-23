package core

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Func func([]interface{}) interface{}

type Task struct {
	Command Func
	Args    []interface{}
}

func (t *Task) Run() interface{} {
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
