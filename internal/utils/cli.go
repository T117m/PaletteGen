package utils

import (
	"fmt"
	"image/color"
	"log"
)

func PrintPalette(p []color.Color) {
	for i, c := range p {
		if c != nil {
			var (
				r, g, b, _ = c.RGBA()
				s          = fmt.Sprintf("\x1b[38;2;%d;%d;%dm██\x1b[0m", uint8(r), uint8(g), uint8(b))
			)

			fmt.Printf("%d: %s %+v\n", i, s, c)
		}
	}
}

func DisplaImage(filePath string) {
	img, err := LoadImage(filePath)
	if err != nil {
		log.Fatalf("Error displaying image: %s", err)
	}

	var (
		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()
	)

	for y := range height {
		for x := range width {
			var (
				pixel      = img.At(x, y)
				rgba       = color.RGBAModel.Convert(pixel)
				r, g, b, _ = rgba.RGBA()
				s          = fmt.Sprintf("\x1b[38;2;%d;%d;%dm██\x1b[0m", uint8(r), uint8(g), uint8(b))
			)
			fmt.Print(s)
		}
		fmt.Println()
	}
}

