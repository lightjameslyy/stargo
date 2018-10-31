package core

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// tables driven testing
var queueTests = []struct {
	q   IQueue
	res bool
}{
	{QueueFactory{}.Create(), true},
	{&Queue{items: []interface{}{1, 2, 3}}, false},
}

func TestQueue_isEmpty(t *testing.T) {
	for _, tt := range queueTests {
		res := tt.q.Empty()
		if res != tt.res {
			t.Errorf("queue empty? expected %v, but got %v", tt.res, res)
		}
	}
}

func TestQueue_Push(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(4)

	q := QueueFactory{}.Create()

	ranks := []int{0, 1, 2, 3}

	for rank := range ranks {
		go func(rank int) {
			defer wg.Done()
			offset := rank * 10
			for j := 0; j < 10; j++ {
				time.Sleep(time.Millisecond)
				q.Push(offset + j)
			}
		}(rank)
	}
	wg.Wait()

	curSize := q.Size()
	if curSize != 40 {
		t.Errorf("q.size: expected %d, but got %d", 40, curSize)
	}
}

func TestQueue_Pop(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(4)

	q := QueueFactory{}.Create()

	for i := 0; i < 40; i++ {
		q.Push(i)
	}

	for rank := range []int{0, 1, 2, 3} {
		go func(rank int) {
			defer wg.Done()
			var nums []int
			time.Sleep(time.Millisecond)
			for j := 0; j < 10; j++ {
				nums = append(nums, q.Pop().(int))
			}
			fmt.Printf("from goroutine %d: %v\n", rank, nums)
		}(rank)
	}
	wg.Wait()

	if q.Empty() == false {
		t.Errorf("after pop, queue should be empty!")
	}

}
