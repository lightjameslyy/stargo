package core

import (
	"fmt"
	"testing"
)

func TestTask_State(t *testing.T) {
	var state State
	fmt.Println(state)
}

func TestTask_Process(t *testing.T) {
	task := Task{Func: func(t T) T {
		fmt.Println(t)
		panic("fake panic")
	}, Args: 1, state: NOTDONE}
	res, err := task.Process()
	fmt.Println(res, err)
}
