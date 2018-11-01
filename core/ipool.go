package core

// IPool interface definition 
type IPool interface {
	Init(int)
	Process(IDag)
}
