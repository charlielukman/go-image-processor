package main

import (
	"fmt"
	imageprocessor "go-image-processor/internal/image-processor"
	"go-image-processor/internal/server"
)

func main() {

	imageProcessor := imageprocessor.NewImageProcessor()
	server := server.NewServer(&imageProcessor)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
