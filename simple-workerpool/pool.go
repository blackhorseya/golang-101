package simpleworkerpool

// Pool represents a worker pool.
type Pool struct {
	workerNums int
	jobsCh     chan *Job
}

// NewPool creates a new worker pool with the given capacity.
func NewPool(workerNums int) *Pool {
	return &Pool{
		workerNums: workerNums,
		jobsCh:     make(chan *Job),
	}
}

func (x *Pool) worker(id int) {

}

func (x *Pool) Run() {

}
