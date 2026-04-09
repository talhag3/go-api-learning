package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/talhag3/go-api-learning/handlers"
)

func main() {
	// Create handler with dependencies
	userHandler := handlers.NewUserHandler()

	// Register routes using ServeMux
	// ServeMux is Go's built-in HTTP request multiplexer (router)
	mux := http.NewServeMux()

	// Register handlers for specific paths
	// The pattern is: method + path -> handler function
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	mux.HandleFunc("POST /users", userHandler.CreateUser)

	// Custom server configuration
	server := &http.Server{
		Addr:         ":8080",          // Address to listen on (:8080 means all interfaces, port 8080)
		Handler:      mux,              // The handler to use (our router)
		ReadTimeout:  10 * time.Second, // Max time to read request (including body)
		WriteTimeout: 10 * time.Second, // Max time to write response
		IdleTimeout:  60 * time.Second, // Max time to keep connection alive between requests
	}

	// Print startup message
	fmt.Println("Server starting on :8080")
	fmt.Println("Try these commands:")
	fmt.Println("  curl http://localhost:8080/users")
	fmt.Println("  curl http://localhost:8080/users/1")
	fmt.Println(`  curl -X POST http://localhost:8080/users -d '{"name":"John","email":"asim@example.com"}' -H "Content-Type: application/json"`)

	// ListenAndServe starts the HTTP server
	// It blocks until an error occurs (like server shutdown)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
