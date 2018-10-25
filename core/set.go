package core

import "sync"

type ISet interface {
	Insert(T)
	Remove(T)
	Has(T) bool
	Empty() bool
	Size() int
}

type Set struct {
	mp    map[T]bool
	mutex sync.Mutex
}

func (s *Set) Insert(v T) {

}

func (s *Set) Remove(v T) {

}

func (s *Set) Has(v T) bool {
	return false
}

func (s *Set) Empty() bool {
	return false
}

func (s *Set) Size() int {
	return 0
}
