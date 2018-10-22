package core

import (
	"fmt"
	"reflect"
	"runtime"
)

type Task struct {
	Command func([]interface{})
	Args    []interface{}
}

func (t *Task) Run() {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(t.Command).Pointer()).Name())
	t.Command(t.Args)
}
