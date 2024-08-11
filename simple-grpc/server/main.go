package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/blackhorseya/golang-101/simple-grpc/pb"
	"google.golang.org/grpc"
)

var _ pb.UserServiceServer = (*userService)(nil)

type userService struct {
}

func (x *userService) GetUser(ctx context.Context, request *pb.UserIdRequest) (*pb.User, error) {
	// TODO: 2024/8/11|sean|implement me
	panic("implement me")
}

func (x *userService) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	// TODO: 2024/8/11|sean|implement me
	panic("implement me")
}

func main() {
	fmt.Println("Hello, World! This is a server.")

	listen, err := net.Listen("tcp", ":50051") //nolint:gosec // this is a demo
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcserver := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcserver, &userService{})

	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
