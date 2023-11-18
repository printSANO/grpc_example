package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	grpc_example "github.com/printSANO/grpc_example/test"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpc_example.NewGreeterClient(conn)

	fmt.Println("Type a message to send to the server. Type 'exit' to quit.")

	for {
		fmt.Print("Enter a message: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message := scanner.Text()

		if strings.ToLower(message) == "exit" {
			fmt.Println("Exiting the client.")
			break
		}

		req := &grpc_example.HelloRequest{
			Name: message,
		}

		res, err := client.SayHello(context.Background(), req)
		if err != nil {
			log.Fatalf("Error calling SayHello: %v", err)
		}

		fmt.Printf("Server response: %s\n", res.Message)
	}
}
