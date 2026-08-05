package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Service string `json:"service"`
	Status  string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Hello from Go API",
		Service: "api",
		Status:  "UP",
	}

	json.NewEncoder(w).Encode(response)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Service: "api",
		Status:  "UP",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health)

	log.Println("API service running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}