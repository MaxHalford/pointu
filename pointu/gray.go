package pointu

import (
	"image"
)

// ImageToGray converts an image.Image into an image.Gray.
func ImageToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}
	return gray
}
