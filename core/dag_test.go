package core

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDag_AddTask(t *testing.T) {
	dag := Dag{
		readyTaskQueue:  QueueFactory{}.Create(),
		notReadyTaskSet: SetFactory{}.Create(),
		state:           DAG_INIT,
		nAll:            0,
		nSent:           0,
		locked:          false,
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 4; i++ {
		go func() {
			for j := 0; j < 25; j++ {
				dag.AddTask(TaskFactory{}.Create())
				wg.Done()
			}
		}()
	}
	wg.Wait()

	size := dag.notReadyTaskSet.Size()
	if size != 100 {
		t.Errorf("size of notReadyTaskQueue expected %d, but got %d", 100, size)
	}
}

/*
Generate a dag using graphviz dot:
	digraph dag1 {
		rankdir = LR
		graph [fontname = "Inconsolata"];
		node [fontname = "Inconsolata"];
		edge [fontname = "Inconsolata"];
		node [
			shape = "circle"
		]
		0 -> 1 -> 2 -> 3 -> 6 -> 7
		2 -> 4 -> 5 -> 6
		8, 9 -> 5 -> 10
		11
	}
 */
func TestDag_AddEdge(t *testing.T) {
	dag := Dag{
		readyTaskQueue:  QueueFactory{}.Create(),
		notReadyTaskSet: SetFactory{}.Create(),
		state:           DAG_INIT,
		nAll:            0,
		nSent:           0,
		locked:          false,
	}

	var tasks [12]ITask
	for i := 0; i < 12; i++ {
		tasks[i] = TaskFactory{}.Create()
		dag.AddTask(tasks[i])
	}

	err := dag.AddEdge(tasks[0], tasks[1])
	fmt.Println(err)
	if err == ErrNoSuchTaskInDag {
		t.Errorf("task not in dag, expect such error: %v", ErrNoSuchTaskInDag)
	}
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

	for i, task := range tasks {
		fmt.Println(i, task, task.ParentsSize())
	}

}

// using same dag as TestDag_AddEdge.
func TestDag_Update(t *testing.T) {
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

	taskChan := make(chan ITask, 4)

	dag.Update(taskChan)

	var wg sync.WaitGroup
	wg.Add(12)

	for i := 0; i < 12; i++ {
		select {
		case task := <-taskChan:
			go func() {
				task.Process()
				wg.Done()
			}()
		}
	}

	wg.Wait()
}
