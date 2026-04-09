# Terminal 1: Start the server
go run main.go

# Terminal 2: Test endpoints
# Get all users
curl http://localhost:8080/users

# Get specific user
curl http://localhost:8080/users/1

# Create user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com","age":30}'

# Test error cases
curl http://localhost:8080/users/abc    # Invalid ID
curl -X POST http://localhost:8080/users -d '{}'  # Missing required fields