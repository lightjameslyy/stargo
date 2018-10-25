package core

import (
	"fmt"
	"testing"
)

func sum(args []T) T {
	res := 0
	for _, n := range args {
		res += n.(int)
	}
	return res
}

var testIntPtr = new(int)

func addToIntPtr(args []T) T {
	intPtr := args[0].(*int)
	*intPtr = 0
	nums := args[1].([]int)
	for _, n := range nums {
		*intPtr += n
	}
	return *intPtr
}

// 表格驱动测试
var taskTests = []struct {
	task Task
}{
	// 匿名函数，参数为空
	{Task{
		func(i []T) T {
			fmt.Println("anonymous function")
			return nil
		},
		nil}},
	// 同类参数
	{Task{
		sum,
		[]T{1, 2, 3, 4, 5}}},
	// 不同类参数
	{Task{
		addToIntPtr,
		[]T{testIntPtr, []int{1, 2, 3, 4, 5}}}},
}

func TestTask_Run(t *testing.T) {
	for _, tt := range taskTests {
		fmt.Println(tt.task.Run())
		fmt.Println(tt.task.TaskInfo())
	}
}

func TestSTask_State(t *testing.T) {
	var state State
	fmt.Println(state)
}

func TestSTask_Process(t *testing.T) {
	task := STask{Func: func(t T) T {
		fmt.Println(t)
		panic("fake panic")
	}, Args: 1, state: NOTDONE}
	res, err := task.Process()
	fmt.Println(res, err)
}
