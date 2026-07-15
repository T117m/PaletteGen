package utils

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"golang.org/x/term"
)

const (
	ErrImgTooBig = "Слишком большое изображение для отображения в терминале"
)

func PrintPalette(p []color.Color) {
	for i, c := range p {
		if c != nil {
			fmt.Printf("%d: %s %+v\n", i, drawPixel(c), c)
		}
	}
}

func DisplayImage(filePath string) {
	img, err := LoadImage(filePath)
	if err != nil {
		log.Fatalf("Error loading image to display: %s", err)
	}

	var (
		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()

		fd = int(os.Stdout.Fd())
	)

	if !term.IsTerminal(fd) {
		return
	}

	tWidth, tHeight, err := term.GetSize(fd)
	if err != nil {
		fmt.Println("Unexpected error getting terminal size:", err)
		return
	}

	if tWidth < width*2 || tHeight*2 < height {
		fmt.Printf("[!] %s: %dx%d vs %dx%d\n", ErrImgTooBig, width, height, tWidth/2, tHeight)
		return
	}

	for y := range height {
		for x := range width {
			fmt.Print(drawPixel(color.RGBAModel.Convert(img.At(x, y))))
		}
		fmt.Println()
	}
}

func drawPixel(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm██\x1b[0m", uint8(r >> 8), uint8(g >> 8), uint8(b >> 8))
}
