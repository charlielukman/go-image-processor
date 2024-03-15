package imageprocessor

import (
	"fmt"
	"os/exec"
)

type ImageProcessor interface {
	ConvertToJPEG(inputFileName string, outputFileName string) error
	ResizeImage(inputFileName string, width int, height int, outputFileName string) error
	CompressImage(inputFileName string, outputFileName string) error
}

type DefaultImageProcessor struct{}

func NewImageProcessor() DefaultImageProcessor {
	return DefaultImageProcessor{}
}

func (d *DefaultImageProcessor) ConvertToJPEG(inputFileName string, outputFileName string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFileName, "-f", "image2", "-vcodec", "mjpeg", outputFileName)
	return cmd.Run()
}

func (d *DefaultImageProcessor) ResizeImage(inputFileName string, width int, height int, outputFileName string) error {
	scaleCommand := fmt.Sprintf("scale=%d:%d", width, height)
	cmd := exec.Command("ffmpeg", "-i", inputFileName, "-vf", scaleCommand, outputFileName)
	return cmd.Run()
}

func (d *DefaultImageProcessor) CompressImage(inputFileName string, outputFileName string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFileName, "-compression_level", "90", outputFileName)
	return cmd.Run()
}
