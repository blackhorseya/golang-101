package main

// Pool represents a pool of workers that can execute jobs.
type Pool struct {
	workerNums int
	jobsQueue  chan *Job
}

// NewPool creates a new pool of workers.
func NewPool(workerNums int) *Pool {
	return &Pool{
		workerNums: workerNums,
		jobsQueue:  make(chan *Job),
	}
}

// Start starts the pool of workers.
func (p *Pool) Start() {
	for i := 0; i < p.workerNums; i++ {
		worker := NewWorker(i, p.jobsQueue)
		worker.Run()
	}
}

// AddJob adds a job to the pool.
func (p *Pool) AddJob(job *Job) {
	p.jobsQueue <- job
}
