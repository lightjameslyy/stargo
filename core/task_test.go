package core

import (
	"fmt"
	"testing"
)

func sum(args []interface{}) interface{} {
	res := 0
	for _, n := range args {
		res += n.(int)
	}
	return res
}

var testIntPtr = new(int)

func addToIntPtr(args []interface{}) interface{} {
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
		func(i []interface{}) interface{} {
			fmt.Println("anonymous function")
			return nil
		},
		nil}},
	// 同类参数
	{Task{
		sum,
		[]interface{}{1, 2, 3, 4, 5}}},
	// 不同类参数
	{Task{
		addToIntPtr,
		[]interface{}{testIntPtr, []int{1, 2, 3, 4, 5}}}},
}

func TestTask_Run(t *testing.T) {
	for _, tt := range taskTests {
		fmt.Println(tt.task.Run())
		fmt.Println(tt.task.TaskInfo())
	}
}
