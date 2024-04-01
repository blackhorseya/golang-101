package main

import (
	"log"
)

func main() {
	pool, err := NewWorkerPool(5)
	if err != nil {
		log.Fatalf("could not create worker pool: %v", err)
	}
	pool.Start()

	for i := 0; i < 10; i++ {
		job := NewSimpleJob(i)
		pool.SubmitJob(job)
	}

	pool.Stop()
}
