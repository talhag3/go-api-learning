package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/talhag3/go-api-learning/handlers"
)

func main() {
	userHandler := handlers.NewUserHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	mux.HandleFunc("POST /users", userHandler.CreateUser)

	// Apply middleware chain
	// Order: Request → Recovery → CORS → Logging → Handler
	wrappedHandler := handlers.Chain(
		handlers.RecoveryMiddleware,
		handlers.CORSMiddleware,
		handlers.LoggingMiddleware,
	)(mux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      wrappedHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server starting on :8080")
	log.Fatal(server.ListenAndServe())
}
