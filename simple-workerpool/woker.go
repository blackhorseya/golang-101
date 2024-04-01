package main

// Worker represents the worker that executes the job.
type Worker struct {
	ID       int
	jobQueue chan *Job
}

// NewWorker creates a new worker with the given id.
func NewWorker(id int) *Worker {
	return &Worker{
		ID:       id,
		jobQueue: make(chan *Job),
	}
}
