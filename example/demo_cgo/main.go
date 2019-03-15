package main

/*

#cgo CFLAGS: -I.

#cgo LDFLAGS: -L. -lfoo

#include "foo.h"
    
*/
import "C"

import (
    "fmt"
	"stargo/core"
	"time"
)

type GoFoo struct {
    foo C.Foo
}

func New(a int) GoFoo {
    var ret GoFoo
    ret.foo = C.FooInit(C.int(a))
    return ret
}
func (f GoFoo) Free() {
    C.FooFree(f.foo)
}
func (f GoFoo) Bar() {
    C.FooBar(f.foo)
}

type T = core.T

func main() {
	dag := core.DagFactory{}.Create()

	ids := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	var tasks [12]core.ITask
	// each task sleep 1 second and print "task i"
	for i, id := range ids {
		tasks[i] = core.TaskFactory{}.Create()
        foo := New(id)
		tasks[i].SetFunc(func(t T) T {
			time.Sleep(time.Second)
            foo.Bar()
            foo.Free()
			return nil
		})
        /*
		tasks[i].SetFunc(func(t T) T {
			time.Sleep(time.Second)
			fmt.Println("task", t)
			return nil
		})
        */
		tasks[i].SetArgs(id)
		dag.AddTask(tasks[i])
	}
	dag.AddEdge(tasks[0], tasks[1])
	dag.AddEdge(tasks[1], tasks[2])
	dag.AddEdge(tasks[2], tasks[3])
	dag.AddEdge(tasks[3], tasks[6])
	dag.AddEdge(tasks[6], tasks[7])
	dag.AddEdge(tasks[2], tasks[4])
	dag.AddEdge(tasks[4], tasks[5])
	dag.AddEdge(tasks[5], tasks[6])
	dag.AddEdge(tasks[8], tasks[5])
	dag.AddEdge(tasks[9], tasks[5])
	dag.AddEdge(tasks[5], tasks[10])

	pool := core.PoolFactory{}.Create()
	pool.Init(5)
	pool.Process(dag)

    fmt.Println("exit...")
}
