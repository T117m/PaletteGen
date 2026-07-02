package core

import (
	"github.com/T117m/PaletteGen/internal/utils"

	"image/color"
	"log"
	"math"
)

func MPA(filePath string, k int) []color.Color {
	img, err := utils.LoadImage(filePath)
	if err != nil {
		log.Fatalf("Error loading image: %s", err)
	}

	var (
		palette  = make([]color.Color, k)
		colorMap = make(map[color.Color]float64)

		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()
	)

	for y := range height {
		for x := range width {
			colorMap[color.RGBAModel.Convert(img.At(x, y))]++
		}
	}

	for i := range k {
		palette[i] = findMax(colorMap)

		delete(colorMap, palette[i])

		for c, n := range colorMap {
			d := distance(c, palette[i])
			colorMap[c] = n / (1 - math.Exp(-0.75*d*d))
		}
	}

	return palette
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
