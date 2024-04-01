package main

// Worker is a struct that represents a worker.
type Worker struct {
	ID       int
	jobQueue chan Job
}

// NewWorker creates a new worker.
func NewWorker(id int, jobQueue chan Job) *Worker {
	return &Worker{
		ID:       id,
		jobQueue: jobQueue,
	}
}

// Run is a method that runs the worker.
func (w *Worker) Run() {
	go func() {
		for job := range w.jobQueue {
			err := job.Execute()
			if err != nil {
				job.OnFailure(err)
			}
		}
	}()
}
