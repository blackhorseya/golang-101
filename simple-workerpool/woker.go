package main

// Worker represents the worker that executes the job.
type Worker struct {
	ID       int
	jobQueue chan *Job
}

// NewWorker creates a new worker with the given id.
func NewWorker(id int, jobQueue chan *Job) *Worker {
	return &Worker{
		ID:       id,
		jobQueue: jobQueue,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	go func() {
		for job := range w.jobQueue {
			job.Execute(w.ID)
		}
	}()
}
