package core

// ITask interface definition
type ITask interface {
	// Set the function which will be executed by this task.
	SetFunc(F)
	// Set arguments for the function.
	SetArgs(T)
	// Get state of the task. State is defined by yourself.
	State() T
	// Add parents of this task. When there's no parent,
	// we consider this task to be READY.
	AddParent(ITask)
	// Return number of parents.
	ParentsSize() int
	// Check if ready.
	IsReady() bool
	// Run the function with supplied arguments.
	Process() (T, error)
}
