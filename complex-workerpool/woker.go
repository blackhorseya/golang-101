package main

// Worker is a struct that represents a worker.
type Worker struct {
	ID       int
	jobQueue chan Job
}

// Run is a method that runs the worker.
func (w *Worker) Run() {
	for job := range w.jobQueue {
		err := job.Execute()
		if err != nil {
			job.OnFailure(err)
		}
	}
}
