package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// In-memory users (simulated database)
var users = []struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// GetAllUsers - GET /users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Success response - wraps in {"success": true, "data": [...]}
	JSON(w, http.StatusOK, users)
}

// GetUser - GET /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path: "/users/1" -> "1"
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		// Error response - wraps in {"success": false, "error": "..."}
		JSONError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		// Error response
		JSONError(w, http.StatusBadRequest, "User ID must be a number")
		return
	}

	// Find user
	for _, user := range users {
		if user.ID == id {
			// Success response - single object
			JSON(w, http.StatusOK, user)
			return
		}
	}

	// Error response - 404 Not Found
	JSONError(w, http.StatusNotFound, "User not found")
}

// CreateUser - POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Check method
	if r.Method != http.MethodPost {
		// Error response - 405 Method Not Allowed
		JSONError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	// Parse body
	var reqBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		// Error response - bad JSON
		JSONError(w, http.StatusBadRequest, "Invalid JSON in request body")
		return
	}

	// Validate
	if reqBody.Name == "" {
		// Error response - validation failed
		JSONError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if reqBody.Email == "" {
		// Error response
		JSONError(w, http.StatusBadRequest, "Email is required")
		return
	}

	// Create user (in real app, save to DB)
	newUser := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    len(users) + 1,
		Name:  reqBody.Name,
		Email: reqBody.Email,
	}
	users = append(users, newUser)

	// Success response - 201 Created
	JSON(w, http.StatusCreated, newUser)
}

// DeleteUser - DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		JSONError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		JSONError(w, http.StatusBadRequest, "User ID must be a number")
		return
	}

	// Find and delete
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			// Success - no content needed for delete
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	JSONError(w, http.StatusNotFound, "User not found")
}
