package main

// Job is an interface that represents a job.
type Job interface {
	// Execute runs the task.
	Execute() error

	// OnFailure is called when the task fails.
	OnFailure(err error)
}
