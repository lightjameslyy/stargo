package core

import "errors"

var (
	ErrDagLocked = errors.New("Dag is already locked!")
	ErrNoSuchTaskInDag = errors.New("No such task in the Dag!")
)
