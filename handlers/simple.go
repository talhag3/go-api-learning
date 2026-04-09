package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// UserHandler holds dependencies for user-related handlers
// This pattern is called "dependency injection"
// Instead of using global variables, we pass dependencies through struct
type UserHandler struct {
	// In later phases, we'll add database connections here
}

// NewUserHandler creates a new UserHandler
// This is a "constructor function" - Go doesn't have constructors like OOP languages
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetAllUsers handles GET /users
// http.ResponseWriter: Used to write the response back to client
// *http.Request: Contains all information about the incoming request
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// In-memory data (we'll replace with database later)
	users := []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		{ID: 1, Name: "Alice", Email: "ali@example.com"},
		{ID: 2, Name: "Bob", Email: "bilal@example.com"},
		{ID: 3, Name: "Charlie", Email: "umar@example.com"},
	}

	// Set response header BEFORE writing body
	// Content-Type tells the client what format the data is in
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder creates a JSON encoder that writes to w
	// Encode marshals the data to JSON and writes it
	if err := json.NewEncoder(w).Encode(users); err != nil {
		// If encoding fails, we can't send JSON error because headers might be sent
		// So we just log and let it fail
		fmt.Printf("Error encoding users: %v\n", err)
	}
}

// GetUser handles GET /users/{id}
// This demonstrates path parameter extraction
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path gives us the full path like "/users/123"
	path := r.URL.Path
	parts := strings.Split(path, "/")

	// parts would be ["", "users", "123"]
	// We need the last part
	if len(parts) < 3 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "User ID must be a number", http.StatusBadRequest)
		return
	}

	// Simulate finding a user
	user := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    id,
		Name:  "Ali",
		Email: "ali@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles POST /users
// This demonstrates request body parsing
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check Content-Type header
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Limit request body size to prevent abuse
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB

	// Declare the variable that will hold the decoded data
	var reqBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	// Decode reads from r.Body and unmarshals JSON into reqBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Strict mode - reject unknown fields
	if err := decoder.Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the input
	if reqBody.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if reqBody.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Create response (in real app, save to DB and get actual ID)
	response := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}{
		ID:    4, // Simulated ID
		Name:  reqBody.Name,
		Email: reqBody.Email,
		Age:   reqBody.Age,
	}

	// 201 Created - standard status for resource creation
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
