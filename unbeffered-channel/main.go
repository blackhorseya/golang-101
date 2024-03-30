package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	interval = 3 * time.Second
)

func main() {
	ch := make(chan string)

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

func producer(ch chan<- string) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			ch <- "ping"
			log.Println("send ping")
		}
	}
}

func consumer(ch <-chan string) {
	for s := range ch {
		if s == "ping" {
			log.Println("reply pong")
		}
	}
}
