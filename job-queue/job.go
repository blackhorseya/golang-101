package jobqueue

import (
	"log"
	"sync"
)

// Job is an interface for submitting work to the job queue.
type Job interface {
	Execute()
}

// JobQueue represents a queue of jobs that processes jobs in a sequential order.
type JobQueue struct {
	jobs   chan Job
	wg     sync.WaitGroup
	quit   chan struct{}
	closed bool
	mu     sync.Mutex
}

// NewJobQueue creates a new job queue.
func NewJobQueue(size int) *JobQueue {
	jq := &JobQueue{
		jobs:   make(chan Job, size),
		wg:     sync.WaitGroup{},
		quit:   make(chan struct{}),
		closed: false,
		mu:     sync.Mutex{},
	}

	jq.wg.Add(1)
	go jq.worker()

	return jq
}

// Enqueue adds a job to the queue.
func (jq *JobQueue) Enqueue(job Job) {
	jq.mu.Lock()
	defer jq.mu.Unlock()

	if jq.closed {
		log.Println("Attempt to enqueue on closed queue")
		return
	}

	select {
	case jq.jobs <- job:
		// Job enqueued successfully
	default:
		// Handle full queue, e.g., log, metrics, or alternative storage
		log.Println("Queue is full. Job rejected.")
	}
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
	jq.mu.Lock()
	if !jq.closed {
		jq.closed = true
		close(jq.quit)
		jq.wg.Wait()
		close(jq.jobs)
	}
	jq.mu.Unlock()
}
