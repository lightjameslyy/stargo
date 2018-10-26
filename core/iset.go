package core

// ISet interface definition
type ISet interface {
	Insert(T)
	Remove(T)
	Has(T) bool
	Empty() bool
	Size() int
}
