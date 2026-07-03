package core

import (
	"image"
	"image/color"
)

func Dominant(img image.Image, k int) []color.Color {
	var (
		palette  []color.Color
		colorMap = make(map[color.Color]uint)

		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()

		colorMapSorted = make([][]color.Color, width*height)
	)

	for y := range height {
		for x := range width {
			colorMap[color.RGBAModel.Convert(img.At(x, y))]++
		}
	}

	for c, n := range colorMap {
		colorMapSorted[n] = append(colorMapSorted[n], c)
	}

	for i, j := (height*width)-1, k; i > 0 && j > 0; i-- {
		if colorMapSorted[i] != nil {
			palette = append(palette, colorMapSorted[i]...)
			j -= len(colorMapSorted[i])
		}
	}

	return palette[:k]
}
