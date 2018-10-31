package core

import "sync"

// Set implements ISet
type Set struct {
	mp    map[T]bool
	mutex sync.Mutex
}

// Insert an element. Thread safe.
func (s *Set) Insert(v T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.mp[v] = true
}

// Remove an element. Thread safe.
func (s *Set) Remove(v T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.mp, v)
}

// Return if the dedicated element is in the set. NOT thread safe.
// NOTE: May not be consistent between two executions.
func (s *Set) Has(v T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.mp[v]
	return ok
}

// Return if the set is empty. NOT thread safe.
func (s *Set) Empty() bool {
	return len(s.mp) == 0
}

// Return size of the set. NOT thread safe.
func (s *Set) Size() int {
	return len(s.mp)
}

// Return slice of all elements in the set.
func (s *Set) All() []T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	res := []T{}
	for k := range s.mp {
		res = append(res, k)
	}
	return res
}
