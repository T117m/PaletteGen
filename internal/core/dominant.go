package core

import (
	"image"
	"image/color"
)

func Dominant(img image.Image, k int) color.Palette {
	var (
		p  color.Palette
		colorMap = make(map[color.Color]uint)

		width, height = getBounds(img)

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
			p = append(p, colorMapSorted[i]...)
			j -= len(colorMapSorted[i])
		}
	}

	return p[:k]
}
