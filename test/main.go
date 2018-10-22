package main

import (
	"fmt"
	"lt.go/stargo/core"
	"reflect"
)

type command func(args ...interface{})

func run(cmd command, args ...interface{}) {
	cmd(args...)
}

func cmd1(args ...interface{}) {
	for _, val := range args {
		fmt.Println(reflect.TypeOf(val), ":", val)
	}
}

func main() {
	t := core.Task{Name: "task1", Command: cmd1, Args: []interface{}{1, "hello", []int{1, 2, 3}}}
	//run(cmd1, 1, "hello", []int{1, 2, 3})
	t.Run()
}
