package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	// todo: 2024/3/30|sean|implement producer
}

func consumer(ch <-chan time.Time) {
	// todo: 2024/3/30|sean|implement consumer
}
