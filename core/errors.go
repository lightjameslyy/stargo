package core

import "errors"

var (
	ErrDagLocked = errors.New("Dag is already locked!")
)
