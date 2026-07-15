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
		centroids  = kMeansPlusPlus(colors, k)
		iterations = 20
	)

	for range iterations {
		clusters := make([][]color.Color, k)

		for _, clr := range colors {
			idx, _ := findClosestCentroid(clr, centroids)
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

func findClosestCentroid(clr color.Color, centroids []color.Color) (Index int, Distance float64) {
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

	return closestIndex, minDistance
}

func getColors(img image.Image) []color.Color {
	var (
		colors []color.Color

		width, height = getBounds(img)
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
		var (
			oldR, oldG, oldB, _ = oldC.RGBA()
			newR, newG, newB, _ = nu[i].RGBA()

			difR = diff(oldR, newR)
			difG = diff(oldG, newG)
			difB = diff(oldB, newB)
		)

		if difR > eps || difG > eps || difB > eps {
			(*old)[i] = nu[i]
			converged = false
		}
	}

	return converged
}

func diff(a, b uint32) float64 {
	return math.Abs(float64(int(a>>8) - int(b>>8)))
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

func kMeansPlusPlus(colors []color.Color, k int) []color.Color {
	var centroids []color.Color
	centroids = append(centroids, colors[rand.Intn(len(colors))])

	for len(centroids) < k {
		var distances []float64

		for i := range colors {
			_, d := findClosestCentroid(colors[i], centroids)
			distances = append(distances, d*d)
		}

		var (
			total              = sum(distances)
			threshold          = rand.Float64() * total
			cumulative float64 = 0
		)

		for i := range colors {
			cumulative += distances[i]
			if cumulative >= threshold {
				centroids = append(centroids, colors[i])
				break
			}
		}
	}

	return centroids
}

func sum(arr []float64) float64 {
	var s float64
	for _, v := range arr {
		s += v
	}
	return s
}
