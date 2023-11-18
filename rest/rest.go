package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	var request HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Hello, %s!", request.Name)
	response := HelloResponse{Message: message}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/sayhello", SayHelloHandler).Methods("POST")

// 	http.Handle("/", r)

// 	fmt.Println("REST API server started. Listening on port 8080.")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }
