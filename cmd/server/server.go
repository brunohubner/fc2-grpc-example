package main

import (
	"fmt"
	"log"
	"net"

	"github.com/brunohubner/fc2-grpc/pb"
	"github.com/brunohubner/fc2-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	fmt.Println("gRPC server running at http://localhost:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not Serve: %v", err)
	}
}
