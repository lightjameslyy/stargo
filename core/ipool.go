package core

// IPool interface definition 
type IPool interface {
	Init(int)
	Bind(IDag)
	Process()
} 
