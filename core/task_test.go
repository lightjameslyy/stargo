package core

import (
	"fmt"
	"testing"
)

func TestTask_State(t *testing.T) {
	var state TaskState
	fmt.Println(state)
}

func TestTask_Process(t *testing.T) {
	task := Task{Func: func(t T) T {
		fmt.Println(t)
		panic("fake panic")
	}, Args: 1, state: TASK_NOTDONE}
	res, err := task.Process()
	fmt.Println(res, err)
}

func TestTask_AddParent(t *testing.T) {
	task := TaskFactory{}.Create()
	task.SetFunc(func(t T) T {
		fmt.Println("task:", t)
		return nil
	})
	task.SetArgs(1)
	parent := TaskFactory{}.Create()
	parent.SetFunc(func(t T) T {
		fmt.Println("parent:", t)
		return nil
	})
	parent.SetArgs(1)

	task.AddParent(parent)

	task.Process()
	parent.Process()

	size := task.ParentsSize()
	if size != 1 {
		t.Errorf("task.Parents.Size expected %d, but got %d", 1, size)
	}
}
