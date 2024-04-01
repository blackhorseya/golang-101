package main

func main() {
	pool := NewPool(5)
	pool.Start()

	for i := 0; i < 100; i++ {
		job := NewJob(i)
		pool.AddJob(job)
	}
}
