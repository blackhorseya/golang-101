package main

var _ Pool = (*workerPool)(nil)

// Pool represents a pool of workers.
type Pool interface {
	// Start starts the pool.
	Start()

	// Stop stops the pool.
	Stop()

	// SubmitJob submits a job to the pool.
	SubmitJob(job Job)
}

type workerPool struct {
	workers   []*Worker
	jobs      []Job
	jobsQueue chan Job
	quit      chan struct{}
}

// NewWorkerPool creates a new worker pool.
func NewWorkerPool(cap int) Pool {
	panic("implement me")
}

func (wp *workerPool) Start() {
	// TODO implement me
	panic("implement me")
}

func (wp *workerPool) Stop() {
	// TODO implement me
	panic("implement me")
}

func (wp *workerPool) SubmitJob(job Job) {
	wp.jobs = append(wp.jobs, job)
}
