package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Register routes
	http.HandleFunc("/", basicHandler)
	http.HandleFunc("/hello", handleHello)

	// Create an HTTP server with the specified address and router as the handler
	server := &http.Server{
		Addr: ":2020",
	}

	// Start the server and listen for incoming requests
	err := server.ListenAndServe()

	// Print an error message if the server fails to start
	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}

// Respond to HTTP requests with for example a message
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

// Respond to HTTP requests with for example a message
func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
