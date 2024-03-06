package server

import (
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/images/convert-to-jpeg", s.ConvertToJPEGHandler)
	r.Post("/images/resize", s.ResizeImageHandler)
	r.Post("/images/compress", s.CompressImageHandler)

	return r
}

func (s *Server) ConvertToJPEGHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // limit your maxMultipartMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a zip writer
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// Retrieve the files from the form data
	files := r.MultipartForm.File["image"]

	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, fileHeader := range files {
		go func(fh *multipart.FileHeader) {
			defer wg.Done()

			file, err := fh.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// Create a temporary file
			tempFile, err := os.CreateTemp(os.TempDir(), "upload-*.png")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer os.Remove(tempFile.Name())

			// Write the uploaded file to the temporary file
			_, err = io.Copy(tempFile, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Convert the image to JPEG using ffmpeg
			outputFileName := fmt.Sprintf("%s.jpg", fh.Filename)
			cmd := exec.Command("ffmpeg", "-i", tempFile.Name(), "-f", "image2", "-vcodec", "mjpeg", outputFileName)
			err = cmd.Run()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Create a file in the zip
			zipFile, err := zipWriter.Create(outputFileName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Read the converted file
			jpeg, err := os.ReadFile(outputFileName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Write the JPEG data to the zip file
			_, err = zipFile.Write(jpeg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}(fileHeader)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/zip")
}

func (s *Server) ResizeImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implement this")
}

func (s *Server) CompressImageHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // limit your maxMultipartMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the file from the form data
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a temporary file
	tempFile, err := os.CreateTemp(os.TempDir(), "upload-*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	// Write the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Compress the image using ffmpeg
	cmd := exec.Command("ffmpeg", "-i", tempFile.Name(), "-compression_level", "90", "output.png")
	err = cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read the compressed file
	compressed, err := os.ReadFile("output.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the compressed data to the response
	w.Header().Set("Content-Type", "image/png")
	w.Write(compressed)
}
