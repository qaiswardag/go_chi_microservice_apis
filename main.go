package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)
	server := &http.Server{
		Addr: ":2020",
		// handler function runs when the server receives a request
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}

// handle all requests
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
