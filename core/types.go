package core

// type alias
type T = interface{}
type F = func(T) T

// enumeration for task state
type TaskState int

const (
	TASK_NOTDONE TaskState = iota
	TASK_DONE
	TASK_WRONG
)

// enumeration for dag state
type DagState int

const (
	DAG_INIT DagState = iota
	DAG_UPDATING
	DAG_NOTDONE
	DAG_DONE
)
