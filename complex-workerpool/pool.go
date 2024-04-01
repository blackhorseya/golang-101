package main

// Pool represents a pool of workers.
type Pool interface {
	// Start starts the pool.
	Start()

	// Stop stops the pool.
	Stop()

	// SubmitJob submits a job to the pool.
	SubmitJob(job Job)
}
