package jobqueue

import (
	"log"
	"testing"
	"time"
)

type ExampleJob struct {
	Name string
}

func (j *ExampleJob) Execute() {
	log.Println("Executing job", j.Name)
	time.Sleep(1 * time.Second)
}

func TestNewJobQueue(t *testing.T) {
	jq := NewJobQueue(10)
	if jq == nil {
		t.Error("expected job queue to be created")
	}

	jq.Enqueue(&ExampleJob{Name: "job1"})
	jq.Enqueue(&ExampleJob{Name: "job2"})
	jq.Enqueue(&ExampleJob{Name: "job3"})

	time.Sleep(5 * time.Second)
	jq.Close()
}
