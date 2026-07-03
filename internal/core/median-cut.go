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
				clr        = simpleColor{uint8(r), uint8(g), uint8(b)}
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
	return uint32(c[0]), uint32(c[1]), uint32(c[2]), uint32(255)
}

func medianCut(colors []simpleColor, n int) []simpleColor {
	if n <= 0 || len(colors) == 0 {
		return nil
	}

	var (
		i, _            = getGreatestRange(colors)
		clr, nextBucket = cut(sortBy(colors, i))
	)

	return append(medianCut(nextBucket, n-1), clr)
}

func cut(sorted []simpleColor) (average simpleColor, bucket []simpleColor) {
	var (
		mid          = len(sorted) / 2
		left, right  = sorted[:mid], sorted[mid:]
		_, lGreatest = getGreatestRange(left)
		_, rGreatest = getGreatestRange(right)
	)

	if lGreatest > rGreatest {
		return getAverage(right), left
	}

	return getAverage(left), right
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

func getGreatestRange(colors []simpleColor) (index int, colorRange uint8) {
	var rRange, gRange, bRange = getRanges(colors)

	if rRange >= gRange && rRange >= bRange {
		return 0, rRange
	} else if gRange >= rRange && gRange >= bRange {
		return 1, gRange
	} else {
		return 2, bRange
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

func sortBy(colors []simpleColor, clrIndex int) []simpleColor {
	var (
		temp   = make([][]simpleColor, 256)
		sorted []simpleColor
	)

	for _, c := range colors {
		temp[c[clrIndex]-1] = append(temp[c[clrIndex]-1], c)
	}

	for _, batch := range colors {
		if colors != nil {
			sorted = append(sorted, batch)
		}
	}

	return sorted
}
