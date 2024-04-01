package main

import (
	"log"
	"sync"
)

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
	start     sync.Once
	stop      sync.Once
	quit      chan struct{}
}

// NewWorkerPool creates a new worker pool.
func NewWorkerPool(workerNums int) Pool {
	panic("implement me")
}

func (wp *workerPool) Start() {
	wp.start.Do(func() {
		log.Println("starting worker pool")
		for i := 0; i < len(wp.workers); i++ {
			wp.workers[i].Run()
		}
	})
}

func (wp *workerPool) Stop() {
	wp.stop.Do(func() {
		log.Println("stopping worker pool")
		close(wp.quit)
	})
}

func (wp *workerPool) SubmitJob(job Job) {
	wp.jobs = append(wp.jobs, job)
}
