package main

import (
	"log"
	"math/rand"
	"time"
)

// Job is an interface that represents a job.
type Job interface {
	// GetID returns the job ID.
	GetID() int

	// Execute runs the task.
	Execute(workerID int) error

	// OnFailure is called when the task fails.
	OnFailure(err error)
}

var _ Job = (*simpleJob)(nil)

type simpleJob struct {
	id int
}

func NewSimpleJob(id int) Job {
	return &simpleJob{id: id}
}

func (j *simpleJob) GetID() int {
	return j.id
}

func (j *simpleJob) Execute(workerID int) error {
	log.Printf("worker %d is executing job %d", workerID, j.id)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second) //nolint:gosec // simulate the job processing
	log.Printf("worker %d finished job %d", workerID, j.id)
	return nil
}

func (j *simpleJob) OnFailure(err error) {
	// TODO implement me
	panic("implement me")
}
