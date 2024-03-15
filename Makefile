# Simple Makefile for a Go project

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

docker-run:
	@docker build -t image-processor .
	@docker run -p 8080:8080 image-processor