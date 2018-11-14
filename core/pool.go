package core

import (
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

	wg.Add(p.workers)

	for i := 0; i < p.workers; i++ {
		rank := i
		createWorker(rank, p.taskChan, &wg)
	}

	wg.Wait()
}

func createWorker(rank int, taskChan chan ITask, wg *sync.WaitGroup) {
	go func() {
		for {
			task, ok := <-taskChan
			if !ok {
				break
			}
			_, err := task.Process()
			if err != nil {
				log.Println(err)
				break
			}
		}
		wg.Done()
		log.Printf("INFO: from POOL: worker[%d] done", rank)
	}()
}
