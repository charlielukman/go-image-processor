package server

import (
	"errors"
	"go-image-processor/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_convertToJPEG(t *testing.T) {
	type args struct {
		imgPrc func(imageProcessor *mocks.ImageProcessor)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success convert to jpeg",
			args: args{
				imgPrc: func(imageProcessor *mocks.ImageProcessor) {
					imageProcessor.On("ConvertToJPEG", mock.Anything, mock.Anything).Return(nil)
				},
			},
			want:    "test.jpg",
			wantErr: false,
		},
		{
			name: "failed convert to jpeg",
			args: args{
				imgPrc: func(imageProcessor *mocks.ImageProcessor) {
					imageProcessor.On("ConvertToJPEG", mock.Anything, mock.Anything).Return(errors.New("error"))
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewImageProcessor(t)

			if tt.args.imgPrc != nil {
				tt.args.imgPrc(mock)
			}

			tempFile, err := os.CreateTemp("", "prefix")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())
			got, err := convertToJPEG(mock, tempFile, "test")
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToJPEG() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertToJPEG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressImage(t *testing.T) {
	type args struct {
		imgPrc func(imageProcessor *mocks.ImageProcessor)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success compress",
			args: args{
				imgPrc: func(imageProcessor *mocks.ImageProcessor) {
					imageProcessor.On("CompressImage", mock.Anything, mock.Anything).Return(nil)
				},
			},
			want:    "test-compressed.png",
			wantErr: false,
		},
		{
			name: "failed compress",
			args: args{
				imgPrc: func(imageProcessor *mocks.ImageProcessor) {
					imageProcessor.On("CompressImage", mock.Anything, mock.Anything).Return(errors.New("error"))
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewImageProcessor(t)

			if tt.args.imgPrc != nil {
				tt.args.imgPrc(mock)
			}

			tempFile, err := os.CreateTemp("", "prefix")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name()) // clean up
			got, err := compressImage(mock, tempFile, "test")
			if (err != nil) != tt.wantErr {
				t.Errorf("compressImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("compressImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resizeImage(t *testing.T) {
	type args struct {
		imgPrc func(imageProcessor *mocks.ImageProcessor)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success resize",
			args: args{
				imgPrc: func(imageProcessor *mocks.ImageProcessor) {
					imageProcessor.On("ResizeImage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				},
			},
			want:    "test-resized.png",
			wantErr: false,
		},
		{
			name: "failed resize",
			args: args{
				imgPrc: func(imageProcessor *mocks.ImageProcessor) {
					imageProcessor.On("ResizeImage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewImageProcessor(t)

			if tt.args.imgPrc != nil {
				tt.args.imgPrc(mock)
			}

			tempFile, err := os.CreateTemp("", "prefix")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name()) // clean up
			got, err := resizeImage(mock, tempFile, 100, 100, "test")
			if (err != nil) != tt.wantErr {
				t.Errorf("resizeImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("resizeImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
