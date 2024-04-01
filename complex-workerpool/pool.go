package main

import (
	"errors"
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
	jobsQueue chan Job
	start     sync.Once
	stop      sync.Once
	quit      chan struct{}
}

// NewWorkerPool creates a new worker pool.
func NewWorkerPool(workerNums int) (Pool, error) {
	if workerNums <= 0 {
		return nil, errors.New("workerNums must be greater than 0")
	}

	workers := make([]*Worker, workerNums)
	jobsQueue := make(chan Job)
	quit := make(chan struct{})

	for i := 0; i < workerNums; i++ {
		workers[i] = NewWorker(i, jobsQueue)
	}

	return &workerPool{
		workers:   workers,
		jobsQueue: jobsQueue,
		start:     sync.Once{},
		stop:      sync.Once{},
		quit:      quit,
	}, nil
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
	select {
	case <-wp.quit:
		log.Println("worker pool is stopped, cannot submit job")
		return
	default:
		log.Printf("submitting job %d\n", job.GetID())
		wp.jobsQueue <- job
		log.Printf("job %d is submitted\n", job.GetID())
	}
}
