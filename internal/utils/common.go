package utils

import (
	"errors"
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
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("Файл %s не найден", filePath)
		}
		return nil, err
	}
	defer file.Close()

	var (
		split     = strings.SplitAfter(file.Name(), ".")
		extension = split[len(split)-1]

		img image.Image
	)

	switch extension {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	default:
		return nil, fmt.Errorf("Неподдерживаемый формат: %s", extension)
	}

	if err != nil {
		return nil, fmt.Errorf("Ошибка при декодировании файла: %s", err)
	}

	return img, nil
}
