package core

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

func KMeans(img image.Image, k int) color.Palette {
	var (
		p color.Palette

		colors     = getColors(img)
		centroids  = getRandomCentroids(colors, k)
		iterations = 20
	)

	for range iterations {
		clusters := make([][]color.Color, k)

		for _, clr := range colors {
			idx := findClosestCentroid(clr, centroids)
			clusters[idx] = append(clusters[idx], clr)
		}

		newCentroids := make([]color.Color, k)
		for i := range k {
			if len(clusters[i]) == 0 {
				newCentroids[i] = colors[rand.Intn(len(colors))]
			} else {
				newCentroids[i] = getAverage(clusters[i])
			}
		}

		if checkConvergence(&centroids, newCentroids) {
			break
		}

		centroids = newCentroids
	}

	for _, c := range centroids {
		p = append(p, c)
	}

	return p
}

func findClosestCentroid(clr color.Color, centroids []color.Color) int {
	var (
		closestIndex = 0
		minDistance  = distance(clr, centroids[0])
	)

	for i := 1; i < len(centroids); i++ {
		d := distance(clr, centroids[i])

		if d < minDistance {
			minDistance = d
			closestIndex = i
		}
	}

	return closestIndex
}

func getColors(img image.Image) []color.Color {
	var (
		colors []color.Color

		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()
	)

	for y := range height {
		for x := range width {
			colors = append(colors, color.RGBAModel.Convert(img.At(x, y)))
		}
	}

	return colors
}

func checkConvergence(old *[]color.Color, nu []color.Color) bool {
	var (
		converged = true
		eps       = 1.0
	)

	for i, oldC := range *old {
		oldR, oldG, oldB, _ := oldC.RGBA()
		newR, newG, newB, _ := nu[i].RGBA()

		difR := math.Abs(float64(int(oldR>>8) - int(newR>>8)))
		difG := math.Abs(float64(int(oldG>>8) - int(newG>>8)))
		difB := math.Abs(float64(int(oldB>>8) - int(newB>>8)))

		if difR > eps || difG > eps || difB > eps {
			(*old)[i] = nu[i]
			converged = false
		}
	}

	return converged
}

func getAverage(colors []color.Color) color.Color {
	var (
		rSum, gSum, bSum int
		avg              color.RGBA

		n = len(colors)
	)

	for _, c := range colors {
		r, g, b, _ := c.RGBA()
		rSum += int(r >> 8)
		gSum += int(g >> 8)
		bSum += int(b >> 8)
	}

	avg.R = uint8(rSum / n)
	avg.G = uint8(gSum / n)
	avg.B = uint8(bSum / n)
	avg.A = uint8(255)

	return avg
}

func getRandomCentroids(colors []color.Color, k int) []color.Color {
	centroids := make([]color.Color, k)

	for i := range k {
		centroids[i] = colors[rand.Intn(len(colors))]
	}

	return centroids
}
