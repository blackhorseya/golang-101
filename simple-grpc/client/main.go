package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blackhorseya/golang-101/simple-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello, World! This is a client.")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user, err := client.GetUser(context.Background(), &pb.UserIdRequest{Id: "123"})
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return
	}
	log.Printf("user: %v", user.GetName())
}
