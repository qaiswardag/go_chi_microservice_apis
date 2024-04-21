package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Create an HTTP server with the specified address and router as the handler
	server := &http.Server{
		Addr:    ":2020",
		Handler: http.HandlerFunc(basicHandler),
	}

	// Start the server and listen for incoming requests
	err := server.ListenAndServe()

	// Print an error message if the server fails to start
	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}

// handle all requests
// Respond to HTTP requests with for example a message
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}
