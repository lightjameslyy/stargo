package core

// Pool implements IPool
type Pool struct {
	dag IDag
}

func (p *Pool) Init(workers int)  {

}

func (p *Pool) Bind(dag IDag)  {
	p.dag = dag
}

func (p *Pool) Process()  {

}
