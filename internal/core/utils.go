package core

import "image"

type rgb8 [3]uint8 // 0 - R, 1 - G, 2 - B

const (
	R int = iota
	G
	B
)

// Implements the color.Color interface
func (c rgb8) RGBA() (r, g, b, a uint32) {
	return uint32(c[R]) << 8, uint32(c[G]) << 8, uint32(c[B]) << 8, uint32(255) << 8
}

func getAverageSimple(colors []rgb8) rgb8 {
	var (
		r, g, b int
		avg     rgb8

		n = len(colors)
	)

	for _, c := range colors {
		r += int(c[R])
		g += int(c[G])
		b += int(c[B])
	}

	avg[R] = uint8(r / n)
	avg[G] = uint8(g / n)
	avg[B] = uint8(b / n)

	return avg
}

func getBounds(img image.Image) (width, height int) {
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy()
}
