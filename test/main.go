package main

import (
	"fmt"
	"lt.go/stargo/core"
	"reflect"
)

func cmd1(args []interface{}) {
	for _, val := range args {
		fmt.Println(reflect.TypeOf(val), ":", val)
	}
}

func main() {
	t := core.Task{Command: cmd1, Args: []interface{}{1, "hello", []int{1, 2, 3, 4}}}
	t.Run()
}
