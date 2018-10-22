package core

import "fmt"

type Task struct {
	Name    string
	Command func(...interface{})
	Args    []interface{}
}

func (t *Task) Run() {
	fmt.Println("Task Name: ", t.Name)
	t.Command(t.Args...)
}
