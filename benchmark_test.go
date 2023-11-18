package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/printSANO/grpc_example/rest"
	"github.com/printSANO/grpc_example/server"
	grpc_example "github.com/printSANO/grpc_example/test"
	"google.golang.org/grpc"
)

// BenchmarkGRPC benchmarks the gRPC server.
func BenchmarkGRPC(b *testing.B) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpc_example.NewGreeterClient(conn)

	for i := 0; i < b.N; i++ {
		req := &grpc_example.HelloRequest{
			Name: "World",
		}

		_, err := client.SayHello(context.Background(), req)
		if err != nil {
			b.Fatalf("Error calling SayHello: %v", err)
		}
	}
}

type HelloRequest struct {
	Name string `json:"name"`
}

// BenchmarkREST benchmarks the REST API server.
func BenchmarkREST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://localhost:8080/sayhello"

		// Dummy request data
		requestData := HelloRequest{Name: "World"}

		// Convert request data to JSON
		requestBody, err := json.Marshal(requestData)
		if err != nil {
			b.Fatalf("Error marshalling JSON: %v", err)
		}

		// Send POST request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			b.Fatalf("Error sending HTTP request: %v", err)
		}
		resp.Body.Close()
	}
}

func runRESTServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/sayhello", rest.SayHelloHandler).Methods("POST")

	http.Handle("/", r)

	fmt.Println("REST API server started. Listening on port 8080.")
	return http.ListenAndServe(":8080", nil)
}

func init() {
	// Start gRPC server
	go func() {
		if err := server.RunGRPCServer(); err != nil {
			panic(err)
		}
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Start REST API server
	go func() {
		if err := runRESTServer(); err != nil {
			panic(err)
		}
	}()

	// Wait for the servers to start
	time.Sleep(1 * time.Second)
}
