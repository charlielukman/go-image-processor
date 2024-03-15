package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	imageprocessor "go-image-processor/internal/image-processor"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port           int
	ImageProcessor imageprocessor.ImageProcessor
}

func NewServer(imageProcessor imageprocessor.ImageProcessor) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:           port,
		ImageProcessor: imageProcessor,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
