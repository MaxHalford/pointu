package pointu

import (
	"image"
	"os"
)

// LoadImage reads and loads an image from a file path.
func LoadImage(path string) (image.Image, error) {
	infile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func maxFloat64s(floats []float64) (max float64) {
	for _, f := range floats {
		if f > max {
			max = f
		}
	}
	return
}

func RescaleFloat64s(floats []float64, newMin, newMax float64) []float64 {
	var (
		oldMin   = maxFloat64s(floats)
		rescaled = make([]float64, len(floats))
	)
	for i, f := range floats {
		rescaled[i] = newMin + (newMax-newMin)*f/oldMin
	}
	return rescaled
}
