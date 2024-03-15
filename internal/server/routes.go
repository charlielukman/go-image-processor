package server

import (
	"archive/zip"
	imageprocessor "go-image-processor/internal/image-processor"
	"mime/multipart"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	h := NewHandler(s.ImageProcessor)

	r.Post("/images/convert-to-jpeg", h.ConvertToJPEGHandler)
	r.Post("/images/resize", h.ResizeImageHandler)
	r.Post("/images/compress", h.CompressImageHandler)

	return r
}

type Handler struct {
	imageProcessor imageprocessor.ImageProcessor
}

func NewHandler(imageProcessor imageprocessor.ImageProcessor) *Handler {
	return &Handler{
		imageProcessor: imageProcessor,
	}
}

func (h *Handler) ConvertToJPEGHandler(w http.ResponseWriter, r *http.Request) {
	files, err := parseMultipartForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

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

			tempFile, err := createTempFile(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer os.Remove(tempFile.Name())

			outputFile, err := convertToJPEG(h.imageProcessor, tempFile, fh.Filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = writeToZip(zipWriter, outputFile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}(fileHeader)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/zip")
}

func (h *Handler) ResizeImageHandler(w http.ResponseWriter, r *http.Request) {
	resizeWidth, resizeHeight := 800, 600

	files, err := parseMultipartForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

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

			tempFile, err := createTempFile(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer os.Remove(tempFile.Name())

			outputFileName, err := resizeImage(h.imageProcessor, tempFile, resizeWidth, resizeHeight, fh.Filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = writeToZip(zipWriter, outputFileName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}(fileHeader)
	}

	wg.Wait()
	w.Header().Set("Content-Type", "application/zip")
}

func (h *Handler) CompressImageHandler(w http.ResponseWriter, r *http.Request) {
	files, err := parseMultipartForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

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

			tempFile, err := createTempFile(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer os.Remove(tempFile.Name())

			outputFileName, err := compressImage(h.imageProcessor, tempFile, fh.Filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = writeToZip(zipWriter, outputFileName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}(fileHeader)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/zip")
}
