package main

import (
	"log"
	"net"

	"user-service/internal/api"
	"user-service/internal/repository"
	"user-service/internal/service"

	"google.golang.org/grpc"
)

func main() {
	// Initialize our in-memory repository
	userRepository := repository.NewUserRepository()

	// Initialize our service layer with the repository
	userService := service.NewUserService(userRepository)

	// Initialize our gRPC server
	grpcServer := grpc.NewServer()

	// Register our UserService implementation with the gRPC server
	api.RegisterUserServiceServer(grpcServer, api.NewUserServiceServer(userService))

	// Start the gRPC server
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
