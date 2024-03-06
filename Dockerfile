# Use the official golang image as the base image
FROM golang:1.22-alpine

# Install ffmpeg
RUN apk update && apk add ffmpeg

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main cmd/api/main.go

# Set the entry point for the container
ENTRYPOINT ["./main"]
