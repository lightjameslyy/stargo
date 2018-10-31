package core

import (
	"fmt"
	"testing"
	"time"
)

func TestPool_Process(t *testing.T) {
	dag := Dag{
		readyTaskQueue:  QueueFactory{}.Create(),
		notReadyTaskSet: SetFactory{}.Create(),
		state:           DAG_INIT,
		nAll:            0,
		nSent:           0,
		locked:          false,
	}

	ids := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	var tasks [12]ITask
	// each task sleep 1 second and print "task i"
	for i, id := range ids {
		tasks[i] = TaskFactory{}.Create()
		tasks[i].SetFunc(func(t T) T {
			time.Sleep(time.Second)
			fmt.Println("task", t)
			return nil
		})
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

	pool := PoolFactory{}.Create()
	pool.Init(4)
	pool.Process(&dag)
}
