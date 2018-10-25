package core

import (
	"fmt"
	"sync"
	"testing"
)

func TestSet_Insert(t *testing.T) {
	s := Set{mp: map[T]bool{}}

	var wg sync.WaitGroup
	wg.Add(4)
	for _, rank := range []int{0, 1, 2, 3} {
		go func(rank int) {
			offset := rank * 10
			for i := 0; i < 10; i++ {
				s.Insert(offset + i)
			}
			wg.Done()
		}(rank)
	}

	wg.Wait()

	fmt.Println(s.mp)
	size := s.Size()
	if size != 40 {
		t.Errorf("expected size: %d, but got %d", 40, size)
	}

}

func TestSet_Remove(t *testing.T) {
	s := Set{mp: map[T]bool{}}

	for i := 0; i < 10; i++ {
		s.Insert(i)
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for _, rank := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		go func(rank int) {
			s.Remove(rank)
			wg.Done()
		}(rank)
	}

	wg.Wait()
	fmt.Println(s.mp)

	if s.Size() != 0 {
		t.Errorf("set should be empty!")
	}
}

func TestSet_Has(t *testing.T) {
	s := Set{mp: map[T]bool{}}

	s.Insert(2)
	s.Insert(3)
	s.Insert(4)

	if s.Has(3) != true {
		t.Errorf("set should have %d", 3)
	}

	if s.Has(5) {
		t.Errorf("set doesn't have %d", 5)
	}
}

func TestSet_Empty(t *testing.T) {
	s := Set{mp: map[T]bool{}}

	if s.Empty() != true {
		t.Errorf("set should be empty!")
	}
}

func TestSet_Size(t *testing.T) {
	s := Set{mp: map[T]bool{}}

	size := s.Size()
	if size != 0 {
		t.Errorf("expected size: %d, but got %d", 0, size)
	}

	s.Insert(0)
	s.Insert(1)
	s.Insert(2)

	size = s.Size()
	if size != 3 {
		t.Errorf("expected size: %d, but got %d", 3, size)
	}

	var wg sync.WaitGroup
	wg.Add(8)
	for _, rank := range []int{2, 3, 4, 5, 6, 7, 8, 9} {
		go func(rank int) {
			s.Insert(rank)
			wg.Done()
		}(rank)
	}
	wg.Wait()

	fmt.Println(s.mp)

	size = s.Size()
	if size != 10 {
		t.Errorf("expected size: %d, but got %d", 10, size)
	}

}
