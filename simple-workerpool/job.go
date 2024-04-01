package main

import (
	"log"
	"math/rand"
	"time"
)

// Job represents the job to be run.
type Job struct {
	ID int
}

// NewJob creates a new job with the given id.
func NewJob(id int) *Job {
	return &Job{
		ID: id,
	}
}

// Execute runs the job.
func (j *Job) Execute(workerID int) {
	log.Printf("Worker %d is running job %d\n", workerID, j.ID)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second) //nolint:gosec // simulate the job processing
	log.Printf("Worker %d finished job %d\n", workerID, j.ID)
}
