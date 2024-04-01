package main

import (
	"log"
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
func (j *Job) Execute() {
	log.Printf("Job %d is running\n", j.ID)
}
