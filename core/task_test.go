package core

import (
	"fmt"
	"testing"
)

func TestTask_State(t *testing.T) {
	task := TaskFactory{}.Create()
	state := task.State()
	if state != TASK_NOTDONE {
		t.Errorf("task state expected %v, but got %v", TASK_NOTDONE, state)
	}
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

func TestTask_IsReady(t *testing.T) {
	var tasks [2]ITask
	for i := 0; i < 2; i++ {
		tasks[i] = TaskFactory{}.Create()
	}

	if !tasks[0].IsReady() {
		t.Errorf("task 0 should be ready!")
	}

	tasks[1].AddParent(tasks[0])
	if tasks[1].IsReady() {
		t.Errorf("task 1 should't be ready!")
	}

	res, err := tasks[0].Process()
	fmt.Println(res, err)
	if !tasks[1].IsReady() {
		t.Errorf("task 1 should be ready!")
	}

}
