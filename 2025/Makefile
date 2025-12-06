all: aoc01 aoc02 aoc03 aoc04 aoc05

# Run the application
aoc01:
	@go run cmd/aoc01/main.go
aoc02:
	@go run cmd/aoc02/main.go
aoc03:
	@go run cmd/aoc03/main.go
aoc04:
	@go run cmd/aoc04/main.go
aoc05:
	@go run cmd/aoc05/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -short
format:
	@echo "Formatting..."
	@gofmt -w .

.PHONY: all aoc01 aoc02 test format
