package core

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 表格驱动测试
var queueTests = []struct {
	q   *Queue
	res bool
}{
	{new(Queue), true},
	{&Queue{items: []interface{}{1, 2, 3}}, false},
}

func TestQueue_isEmpty(t *testing.T) {
	for _, tt := range queueTests {
		fmt.Println(*tt.q)
		res := tt.q.isEmpty()
		if res != tt.res {
			t.Errorf("queue: %v, expected %v, but got %v", *tt.q, tt.res, res)
		}
	}
}

func TestQueue_Push(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(4)

	q := new(Queue)

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

	fmt.Println(*q)

	curSize := q.size()
	if curSize != 40 {
		t.Errorf("q.size: expected %d, but got %d", 40, curSize)
	}
}

func TestQueue_Pop(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(4)

	q := new(Queue)

	for i := 0; i < 40; i++ {
		q.Push(i)
	}

	fmt.Println("before pop:", q.items)

	ranks := []int{0, 1, 2, 3}

	for rank := range ranks {
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

	if q.isEmpty() == false {
		t.Errorf("after pop, queue should be empty!")
	}

}
