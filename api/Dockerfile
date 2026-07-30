package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go API Microservice")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("API running on port 8080")
	http.ListenAndServe(":8080", nil)
}