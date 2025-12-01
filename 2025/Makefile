all: aoc01

# Run the application
aoc01:
	@go run cmd/aoc01/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -short
format:
	@echo "Formatting..."
	@gofmt -w .

.PHONY: all aoc01 test format
