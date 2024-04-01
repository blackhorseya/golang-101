package main

import (
	"log"
)

// Job is an interface that represents a job.
type Job interface {
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

func (j *simpleJob) Execute(workerID int) error {
	log.Printf("worker %d is executing job %d", workerID, j.id)
	return nil
}

func (j *simpleJob) OnFailure(err error) {
	// TODO implement me
	panic("implement me")
}
