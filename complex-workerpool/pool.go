package main

import (
	"sync"
)

// Pool represents a pool of workers.
type Pool interface {
	// Start starts the pool.
	Start()

	// Stop stops the pool.
	Stop()

	// SubmitJob submits a job to the pool.
	SubmitJob(job Job)
}

// WorkerPool is a struct that represents a worker pool.
type WorkerPool struct {
	workers []*Worker
	jobs    []Job
	start   sync.Once
	stop    sync.Once
	quit    chan struct{}
}
