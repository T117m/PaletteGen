package core

import (
	"image"
	"image/color"
)

func MedianCut(img image.Image, k int) color.Palette {
	var (
		p color.Palette
		colors  []rgb8

		width, height = getBounds(img)
	)

	for y := range height {
		for x := range width {
			var (
				r, g, b, _ = color.RGBAModel.Convert(img.At(x, y)).RGBA()
				clr        = rgb8{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
			)

			colors = append(colors, clr)
		}
	}

	for _, c := range medianCut(colors, k) {
		p = append(p, c)
	}

	return p
}

func medianCut(colors []rgb8, n int) []rgb8 {
	l := len(colors)

	if n <= 0 || l == 0 {
		return nil
	}
	if n == 1 || l == 1 {
		return []rgb8{getAverageSimple(colors)}
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

func getGreatestRange(colors []rgb8) int {
	var rRange, gRange, bRange = getRanges(colors)

	if rRange >= gRange && rRange >= bRange {
		return 0
	} else if gRange >= rRange && gRange >= bRange {
		return 1
	} else {
		return 2
	}
}

func getRanges(colors []rgb8) (rRange, gRange, bRange uint8) {
	var (
		rMin, gMin, bMin uint8 = 255, 255, 255
		rMax, gMax, bMax uint8 = 0, 0, 0
	)

	for _, c := range colors {
		if c[R] < rMin {
			rMin = c[R]
		}
		if c[R] > rMax {
			rMax = c[R]
		}

		if c[G] < gMin {
			gMin = c[G]
		}
		if c[G] > gMax {
			gMax = c[G]
		}

		if c[B] < bMin {
			bMin = c[B]
		}
		if c[B] > bMax {
			bMax = c[B]
		}
	}

	return rMax - rMin, gMax - gMin, bMax - bMin
}

func sortBy(colors []rgb8, rangeIndex int) []rgb8 {
	var (
		temp   = make([][]rgb8, 256)
		sorted []rgb8
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
