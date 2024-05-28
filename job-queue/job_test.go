package jobqueue

import (
	"sync"
	"testing"
	"time"
)

// Define a simple Job implementation for testing
type TestJob struct {
	id       int
	executed bool
	mu       sync.Mutex
}

func (t *TestJob) Execute() {
	t.mu.Lock()
	t.executed = true
	t.mu.Unlock()
}

func TestJobQueue(t *testing.T) {
	jq := NewJobQueue(2) // Use a small buffer to test the full condition
	defer jq.Close()

	job1 := &TestJob{id: 1}
	job2 := &TestJob{id: 2}
	job3 := &TestJob{id: 3}

	jq.Enqueue(job1)
	jq.Enqueue(job2)
	jq.Enqueue(job3) // This should be rejected as the buffer is only 2

	time.Sleep(100 * time.Millisecond) // Allow some time for jobs to be processed

	if !job1.executed || !job2.executed {
		t.Errorf("Expected job1 and job2 to be executed")
	}

	if job3.executed {
		t.Errorf("Expected job3 to not be executed")
	}
}
