# Project go-image-processor

Golang image processing:
1. Convert image files from PNG to JPEG.
2. Resize images according to specified dimensions.
3. Compress images to reduce file size while maintaining reasonable quality.

Processed images output as a zip file

## Getting Started

To run build and run execute:
```bash
make docker-run
```

To execute image conversion:
endpoint will be:
Converted image(s) will be resulted as .zip
`content-type` multipart/form-data, image(s) on body key: `image`
- POST http://localhost:8080/images/convert-to-jpeg
- POST http://localhost:8080/images/compress
- POST http://localhost:8080/images/resize, for simplicity, resized with width 800, height 600


This repo use ffmpeg to convert image, the installation is added on docker.

## MakeFile

run all make commands with clean tests

run the test suite
```bash
make test
```

build and run using docker
```bash
make docker-run
```