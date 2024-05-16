package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello world")
}

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	http.ListenAndServe(port, nil)
}
