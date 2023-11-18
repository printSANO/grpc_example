package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	url := "http://localhost:8080/sayhello"

	// Dummy request data
	requestData := HelloRequest{Name: "World"}

	// Convert request data to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Send POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Decode the response
	var response HelloResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	// Print the response
	fmt.Printf("Server response: %s\n", response.Message)
}
