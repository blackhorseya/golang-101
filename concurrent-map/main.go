package main

import (
	"sync"
)

// AccessCounter is a struct that holds the access counter.
type AccessCounter struct {
	sync.RWMutex
	visits map[string]int
}

// NewAccessCounter is a constructor that returns a new AccessCounter.
func NewAccessCounter() *AccessCounter {
	return &AccessCounter{
		visits: make(map[string]int),
	}
}

// Visit is a method that increments the access counter for the given username.
func (c *AccessCounter) Visit(username string) {
	c.Lock()
	defer c.Unlock()

	c.visits[username]++
}

// GetVisit is a method that returns the access counter for the given username.
func (c *AccessCounter) GetVisit(username string) int {
	c.RLock()
	defer c.RUnlock()

	return c.visits[username]
}

func main() {
}
