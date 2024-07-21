package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// 連接到 Redis 伺服器
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 伺服器地址
	})

	// 訂閱 Redis 頻道
	go func() {
		sub := rdb.Subscribe(ctx, "chat_channel")
		ch := sub.Channel()

		for msg := range ch {
			fmt.Printf("Received message: %s\n", msg.Payload)
		}
	}()

	// 捕捉中斷信號以優雅地關閉
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 從標準輸入讀取消息並發布到 Redis 頻道
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message := scanner.Text()
			publishMessage(rdb, message)
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading from stdin: %v\n", err)
		}
	}()

	<-sigs

	fmt.Println("Shutting down...")
	rdb.Close()
}

// 發送消息到 Redis 頻道
func publishMessage(rdb *redis.Client, message string) {
	err := rdb.Publish(ctx, "chat_channel", message).Err()
	if err != nil {
		log.Fatalf("Could not publish message: %v", err)
	}
	fmt.Println("Message sent:", message)
}
