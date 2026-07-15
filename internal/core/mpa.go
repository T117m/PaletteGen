package core

import (
	"image"
	"image/color"
	"math"
)

func MPA(img image.Image, k int) color.Palette {
	var (
		p  = make(color.Palette, k)
		colorMap = make(map[color.Color]float64)

		width, height = getBounds(img)
	)

	for y := range height {
		for x := range width {
			colorMap[color.RGBAModel.Convert(img.At(x, y))]++
		}
	}

	for i := range k {
		p[i] = findMax(colorMap)

		delete(colorMap, p[i])

		for c, n := range colorMap {
			d := distance(c, p[i])
			colorMap[c] = n / (1 - math.Exp(-0.75*d*d))
		}
	}

	return p
}

func findMax(colorMap map[color.Color]float64) color.Color {
	var (
		m  float64 = 0
		mC color.Color
	)

	for c, n := range colorMap {
		if n > m {
			m = n
			mC = c
		}
	}

	return mC
}

func distance(x, y color.Color) float64 {
	r1, g1, b1, _ := x.RGBA()
	r2, g2, b2, _ := y.RGBA()

	r1_f, g1_f, b1_f := float64(r1), float64(g1), float64(b1)
	r2_f, g2_f, b2_f := float64(r2), float64(g2), float64(b2)

	return math.Sqrt(2*math.Pow((r1_f-r2_f), 2) + 4*math.Pow((g1_f-g2_f), 2) + 3*math.Pow((b1_f-b2_f), 2))
}
