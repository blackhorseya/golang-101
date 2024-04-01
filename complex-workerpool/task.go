package main

// Task is an interface that all tasks must implement.
type Task interface {
	// Execute runs the task.
	Execute() error

	// OnFailure is called when the task fails.
	OnFailure(err error)
}
