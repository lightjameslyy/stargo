package core

import "sync"

// Set implements ISet
type Set struct {
	mp    map[T]bool
	mutex sync.Mutex
}

func (s *Set) Insert(v T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.mp[v] = true
}

func (s *Set) Remove(v T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.mp, v)
}

func (s *Set) Has(v T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.mp[v]
	return ok
}

func (s *Set) Empty() bool {
	return len(s.mp) == 0
}

func (s *Set) Size() int {
	return len(s.mp)
}
