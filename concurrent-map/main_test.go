package main

import (
	"sync"
	"testing"
)

func TestAccessCounter_ReadWrite(_ *testing.T) {
	var wg sync.WaitGroup
	counter := NewAccessCounter()

	// write operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter.Visit("test")
		}
	}()

	// read operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			_ = counter.GetVisit("test")
		}
	}()

	wg.Wait()
}

func BenchmarkAccessCounter_Visit(b *testing.B) {
	counter := NewAccessCounter()

	b.RunParallel(func(pb *testing.PB) {
		key := "test"
		for pb.Next() {
			counter.Visit(key)
		}
	})
}
