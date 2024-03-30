package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	interval = 1 * time.Second
)

func main() {
	ch := make(chan time.Time, 10)

	go producer(ch)
	go consumer(ch)

	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		log.Println("signal received:", sig)

		close(ch)
	}
}

func producer(ch chan<- time.Time) {
	ticker := time.NewTicker(interval)

	for {
		t := <-ticker.C
		ch <- t
		log.Println("produce a new job")
	}
}

func consumer(ch <-chan time.Time) {
	for t := range ch {
		log.Println("receive a job", t)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second) //nolint:gosec // simulate the job processing
	}
}
