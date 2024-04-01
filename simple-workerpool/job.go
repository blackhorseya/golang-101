package simpleworkerpool

import (
	"fmt"
)

// Job represents the job to be run.
type Job struct {
	ID int
}

// NewJob creates a new job with the given ID.
func NewJob(id int) *Job {
	return &Job{ID: id}
}

func (x *Job) Execute() {
	fmt.Printf("Job %d is running\n", x.ID) //nolint:forbidigo // This is just an example
}
