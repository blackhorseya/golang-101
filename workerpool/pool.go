package main

// Pool represents a pool of workers that can execute jobs.
type Pool struct {
	workerNums int
	workers    []*Worker
	jobs       []*Job
}

// NewPool creates a new pool of workers.
func NewPool(workerNums int) *Pool {
	pool := &Pool{
		workerNums: workerNums,
		workers:    make([]*Worker, workerNums),
		jobs:       make([]*Job, 0),
	}

	for i := 0; i < workerNums; i++ {
		pool.workers[i] = NewWorker(i)
	}

	return pool
}
