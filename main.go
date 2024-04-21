package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a new router using chi package
	router := chi.NewRouter()

	// Use the Logger middleware for logging requests
	router.Use(middleware.Logger)

	// Define a route and handle the route with a function
	router.Get("/", basicHandler)

	// Create an HTTP server with the specified address and router as the handler
	server := &http.Server{
		Addr:    ":2020",
		Handler: router,
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
