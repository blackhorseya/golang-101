package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("Hello, World!")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	ctx, cancelFunc := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		fmt.Println("Waiting for signal")
		<-ctx.Done()
		fmt.Println("Signal received")
	}(ctx)

	<-signalChan
	cancelFunc()

	wg.Wait()

	fmt.Println("Goodbye, World!")
}
