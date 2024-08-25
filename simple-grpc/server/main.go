package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/blackhorseya/golang-101/simple-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var _ pb.UserServiceServer = (*userService)(nil)

type userService struct {
}

func (x *userService) GetUser(ctx context.Context, request *pb.UserIdRequest) (*pb.User, error) {
	return &pb.User{
		Id:   request.Id,
		Name: "Sean",
	}, nil
}

func main() {
	fmt.Println("Hello, World! This is a server.")

	listen, err := net.Listen("tcp", ":50051") //nolint:gosec // this is a demo
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcserver := grpc.NewServer()

	// register user service
	pb.RegisterUserServiceServer(grpcserver, &userService{})

	// register health service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcserver, healthServer)
	healthServer.SetServingStatus("example", grpc_health_v1.HealthCheckResponse_SERVING)

	// register reflection service on gRPC server
	reflection.Register(grpcserver)

	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
