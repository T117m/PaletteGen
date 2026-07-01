package utils

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening %s: %s", filePath, err)
	}
	defer file.Close()

	var (
		extension = strings.SplitAfter(file.Name(), ".")
		img       image.Image
	)

	switch extension[len(extension)-1] {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	default:
		err = image.ErrFormat
	}

	if err != nil {
		return nil, fmt.Errorf("Error decoding file: %s", err)
	}

	return img, nil
}
