package core

import (
	"image"
	"image/color"
)

type simpleColor [3]uint8 // 0 - R, 1 - G, 2 - B

func MedianCut(img image.Image, k int) []color.Color {
	var (
		palette []color.Color
		colors  []simpleColor

		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()
	)

	for y := range height {
		for x := range width {
			var (
				r, g, b, _ = color.RGBAModel.Convert(img.At(x, y)).RGBA()
				clr        = simpleColor{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
			)

			colors = append(colors, clr)
		}
	}

	for _, c := range medianCut(colors, k) {
		palette = append(palette, c)
	}

	return palette
}

// Implements the color.Color interface
func (c simpleColor) RGBA() (r, g, b, a uint32) {
	return uint32(c[0]) << 8, uint32(c[1]) << 8, uint32(c[2]) << 8, uint32(255) << 8
}

func medianCut(colors []simpleColor, n int) []simpleColor {
	l := len(colors)

	if n <= 0 || l == 0 {
		return nil
	}
	if n == 1 || l == 1 {
		return []simpleColor{getAverage(colors)}
	}

	var (
		colorRange = getGreatestRange(colors)
		sorted     = sortBy(colors, colorRange)
		mid        = l / 2
		nLeft      = n / 2
		nRight     = n - nLeft
	)

	return append(medianCut(sorted[:mid], nLeft), medianCut(sorted[mid:], nRight)...)
}

func getAverage(colors []simpleColor) simpleColor {
	var (
		r, g, b int
		avg     simpleColor

		n = len(colors)
	)

	for _, c := range colors {
		r += int(c[0])
		g += int(c[1])
		b += int(c[2])
	}

	avg[0] = uint8(r / n)
	avg[1] = uint8(g / n)
	avg[2] = uint8(b / n)

	return avg
}

func getGreatestRange(colors []simpleColor) int {
	var rRange, gRange, bRange = getRanges(colors)

	if rRange >= gRange && rRange >= bRange {
		return 0
	} else if gRange >= rRange && gRange >= bRange {
		return 1
	} else {
		return 2
	}
}

func getRanges(colors []simpleColor) (rRange, gRange, bRange uint8) {
	var (
		rMin, gMin, bMin uint8 = 255, 255, 255
		rMax, gMax, bMax uint8 = 0, 0, 0
	)

	for _, c := range colors {
		if c[0] < rMin {
			rMin = c[0]
		}
		if c[0] > rMax {
			rMax = c[0]
		}

		if c[1] < gMin {
			gMin = c[1]
		}
		if c[1] > gMax {
			gMax = c[1]
		}

		if c[2] < bMin {
			bMin = c[2]
		}
		if c[2] > bMax {
			bMax = c[2]
		}
	}

	return rMax - rMin, gMax - gMin, bMax - bMin
}

func sortBy(colors []simpleColor, rangeIndex int) []simpleColor {
	var (
		temp   = make([][]simpleColor, 256)
		sorted []simpleColor
	)

	for _, c := range colors {
		temp[c[rangeIndex]] = append(temp[c[rangeIndex]], c)
	}

	for _, batch := range colors {
		if colors != nil {
			sorted = append(sorted, batch)
		}
	}

	return sorted
}
