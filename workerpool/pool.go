package main

// Pool represents a pool of workers that can execute jobs.
type Pool struct {
	workerNums int
	workers    []*Worker
	jobs       []*Job
}
