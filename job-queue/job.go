package jobqueue

import (
	"sync"
)

// Job is an interface for submitting work to the job queue.
type Job interface {
	Execute()
}

// JobQueue represents a queue of jobs that processes jobs in a sequential order.
type JobQueue struct {
	jobs chan Job
	wg   sync.WaitGroup
	quit chan bool
}

// NewJobQueue creates a new job queue.
func NewJobQueue() *JobQueue {
	jq := &JobQueue{
		jobs: make(chan Job),
		quit: make(chan bool),
	}

	jq.wg.Add(1)
	go jq.worker()

	return jq
}

// Enqueue adds a job to the queue.
func (jq *JobQueue) Enqueue(job Job) {
	jq.jobs <- job
}

// worker runs in a goroutine and processes jobs in the order they are received.
func (jq *JobQueue) worker() {
	defer jq.wg.Done()

	for {
		select {
		case job := <-jq.jobs:
			job.Execute()
		case <-jq.quit:
			return
		}
	}
}

// Close waits for the worker to finish and closes the job queue.
func (jq *JobQueue) Close() {
	close(jq.quit)
	jq.wg.Wait()
	close(jq.jobs)
}
