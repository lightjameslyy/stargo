package core

import "log"

// Pool implements IPool
type Pool struct {
	dag      IDag
	doneChan chan bool
	workers  int
	taskChan chan ITask
}

func (p *Pool) Init(workers int) {
	p.doneChan = make(chan bool, workers)
	p.taskChan = make(chan ITask, workers)
}

func (p *Pool) Bind(dag IDag) {
	dag.Lock()
	p.dag = dag
}

func (p *Pool) Process() {
	p.dag.Update(p.taskChan)

	for i := 0; i < p.workers; i++ {
		createWorker(p.taskChan, p.doneChan)
	}

	for i := 0; i < p.workers; i++ {
		select {
		// 先不管是否正常退出
		case <-p.doneChan:
		}
	}
}

func createWorker(taskChan chan ITask, doneChan chan<- bool) {
	go func() {
		select {
		case task:= <- taskChan:
			_, err := task.Process()
			if err != nil {
				log.Println(err)
				break
			}
		}
		doneChan <- true
	}()
}
