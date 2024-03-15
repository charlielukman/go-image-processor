package server

import (
	"archive/zip"
	"fmt"
	imageprocessor "go-image-processor/internal/image-processor"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func parseMultipartForm(r *http.Request) ([]*multipart.FileHeader, error) {
	err := r.ParseMultipartForm(10 << 20) // limit your maxMultipartMemory
	if err != nil {
		return nil, err
	}

	// Retrieve the files from the form data
	files := r.MultipartForm.File["image"]

	return files, nil
}

func createTempFile(file multipart.File) (*os.File, error) {
	// Create a temporary file
	tempFile, err := os.CreateTemp(os.TempDir(), "upload-*.png")
	if err != nil {
		return nil, err
	}

	// Write the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}

func convertToJPEG(imgPrc imageprocessor.ImageProcessor, tempFile *os.File, filename string) (string, error) {
	outputFileName := fmt.Sprintf("%s.jpg", filename)

	err := imgPrc.ConvertToJPEG(tempFile.Name(), outputFileName)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}

func writeToZip(zipWriter *zip.Writer, outputFileName string) error {
	// Create a file in the zip
	zipFile, err := zipWriter.Create(outputFileName)
	if err != nil {
		return err
	}

	// Read the converted file
	jpeg, err := os.ReadFile(outputFileName)
	if err != nil {
		return err
	}

	// Write the JPEG data to the zip file
	_, err = zipFile.Write(jpeg)
	if err != nil {
		return err
	}

	return nil
}

func compressImage(imgPrc imageprocessor.ImageProcessor, tempFile *os.File, filename string) (string, error) {
	outputFileName := fmt.Sprintf("%s-compressed.png", filename)

	err := imgPrc.CompressImage(tempFile.Name(), outputFileName)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}

func resizeImage(imgPrc imageprocessor.ImageProcessor, tempFile *os.File, width int, height int, filename string) (string, error) {
	outputFileName := fmt.Sprintf("%s-resized.png", filename)

	err := imgPrc.ResizeImage(tempFile.Name(), width, height, outputFileName)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}
