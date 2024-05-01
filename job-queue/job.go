package jobqueue

import (
	"sync"
)

// Job is an interface that represents a job that can be executed by a worker.
type Job interface {
	Execute() error
}

// JobQueue is a struct that represents a job queue.
type JobQueue struct {
	jobs chan Job
	wg   sync.WaitGroup
	quit chan struct{}
}

// NewJobQueue creates a new job queue with the given capacity.
func NewJobQueue(size int) *JobQueue {
	jq := &JobQueue{
		jobs: make(chan Job, size),
		quit: make(chan struct{}),
	}

	jq.wg.Add(1)
	go jq.worker()

	return jq
}

// AddJob adds a job to the job queue.
func (jq *JobQueue) AddJob(job Job) {
	jq.jobs <- job
}

func (jq *JobQueue) worker() {
	defer jq.wg.Done()

	for {
		select {
		case job := <-jq.jobs:
			_ = job.Execute()
		case <-jq.quit:
			return
		}
	}
}
