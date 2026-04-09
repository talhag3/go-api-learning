package models

type User struct {
	ID    int    `json:"id"` // Struct Tags
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// The backtick annotations are called "struct tags"
// `json:"id"` tells the JSON encoder/decoder to use "id" as the key
// Without tags, it would use the field name (which is uppercase in Go)
