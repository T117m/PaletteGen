package core

import (
	"image"
	"image/color"
	"math/rand"
)

func KMeans(img image.Image, k int) color.Palette {
	var (
		p color.Palette

		bounds        = img.Bounds()
		width, height = bounds.Dx(), bounds.Dy()

		centroids = getRandomCentroids(img, k, width, height)
		converged = false
	)

	for !converged {
		clusters := make([][]color.Color, k)

		for y := range height {
			for x := range width {
				var (
					pixel = img.At(x, y)
					clr   = color.RGBAModel.Convert(pixel)

					closestIndex = 0
					minDistance  = distance(clr, centroids[0])
				)

				for i := 1; i < k; i++ {
					d := distance(clr, centroids[i])

					if d < minDistance {
						minDistance = d
						closestIndex = i
					}
				}

				clusters[closestIndex] = append(clusters[closestIndex], clr)
			}
		}

		newCentroids := make([]color.Color, k)
		for i := range k {
			if len(clusters[i]) == 0 {
				converged = true
				break
			}

			newCentroids[i] = getAverage(clusters[i])
		}

		if !converged {
			converged = checkConvergence(&centroids, newCentroids)
		}
	}

	for _, c := range centroids {
		p = append(p, c)
	}

	return p
}

func checkConvergence(old *[]color.Color, nu []color.Color) bool {
	converged := true

	for i, oldC := range *old {
		oldR, oldG, oldB, _ := oldC.RGBA()
		newR, newG, newB, _ := nu[i].RGBA()

		if oldR != newR || oldG != newG || oldB != newB {
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

func getRandomCentroids(img image.Image, k, width, height int) []color.Color {
	centroids := make([]color.Color, k)

	for i := range k {
		pixel := img.At(rand.Intn(width-1), rand.Intn(height-1))
		centroids[i] = color.RGBAModel.Convert(pixel)
	}

	return centroids
}
