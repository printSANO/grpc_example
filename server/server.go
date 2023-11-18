package server

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_example "github.com/printSANO/grpc_example/test"
	"google.golang.org/grpc"
)

type server struct {
	grpc_example.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *grpc_example.HelloRequest) (*grpc_example.HelloResponse, error) {
	message := req.Name

	return &grpc_example.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", message),
	}, nil
}

func RunGRPCServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpc_example.RegisterGreeterServer(s, &server{})

	fmt.Println("Server started. Listening on port 50051.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return nil
}
