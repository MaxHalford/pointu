package pointu

import (
	"image"
)

// A PDF is a probability density function.
type PDF [][]float64

// makePDF generates a probability density function from an image.Gray.
func MakePDF(gray *image.Gray) PDF {
	var (
		bounds = gray.Bounds()
		pdf    = make(PDF, bounds.Dx())
	)
	for x := 0; x < bounds.Dx(); x++ {
		pdf[x] = make([]float64, bounds.Dy())
		for y := 0; y < bounds.Dy(); y++ {
			pdf[x][y] = 255 - float64(gray.GrayAt(x, y).Y)
		}
	}
	return pdf
}
