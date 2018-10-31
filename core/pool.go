package core

import (
	"fmt"
	"log"
	"sync"
)

// Pool implements IPool
type Pool struct {
	workers  int
	taskChan chan ITask
}

func (p *Pool) Init(workers int) {
	p.workers = workers
	p.taskChan = make(chan ITask, workers)
}

func (p *Pool) Process(dag IDag) {
	dag.Lock()
	dag.Update(p.taskChan)

	var wg sync.WaitGroup
	/*
	wg.Add(p.workers)

	for i := 0; i < p.workers; i++ {
		createWorker(p.taskChan, wg)
	}
	*/
	wg.Add(dag.Size())

	for i := 0; i < dag.Size(); i++ {
		select {
		case task := <-p.taskChan:
			go func() {
				_, err := task.Process()
				wg.Done()
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}

	wg.Wait()
}

func worker(task ITask, wg sync.WaitGroup) {
	go func() {
		_, err := task.Process()
		wg.Done()
		if err != nil {
			log.Println(err)
		}
	}()
}

func createWorker(taskChan chan ITask, wg sync.WaitGroup) {
	go func() {
		/*
		select {
		case task, ok := <-taskChan:
			if !ok {
				fmt.Println("closed")
				break
			}
			_, err := task.Process()
			if err != nil {
				log.Println(err)
				break
			}
		}
		*/
		for {
			task, open := <-taskChan
			if !open {
				break
			}
			_, err := task.Process()
			wg.Done()
			if err != nil {
				log.Println(err)
				break
			}
		}
		fmt.Println("worker done")
	}()
}
