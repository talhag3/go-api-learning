package handlers

import (
	"encoding/json"
	"net/http"
)

// Response is a standard API response wrapper
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"` // omitempty: don't include if nil/empty
	Error   string      `json:"error,omitempty"`
}

// JSON sends a JSON response with status code
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    data,
	})
}

// JSONError sends a JSON error response
func JSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   message,
	})
}
