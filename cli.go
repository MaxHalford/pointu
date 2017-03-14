package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/MaxHalford/pointu/pointu"
	"github.com/fogleman/gg"
)

var (
	inPath     = flag.String("in", "", "Input path")
	outPath    = flag.String("out", "", "Output path")
	nPoints    = flag.Int("points", 500, "Number of points")
	threshold  = flag.Int("threshold", 200, "8 bit threshold used for importance sampling")
	resolution = flag.Float64("resolution", 1, "Resolution ratio")
	iterations = flag.Int("iterations", 10, "Number of iterations")
	rMin       = flag.Float64("rmin", 1, "Minimum point radius")
	rMax       = flag.Float64("rmax", 1, "Maximum point radius")
	withColor  = flag.Bool("color", false, "Use input image color instead of black")
	rng        = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {

	// Parse the command-line arguments
	flag.Parse()

	// Open the input image
	img, err := pointu.LoadImage(*inPath)
	if err != nil {
		fmt.Printf("Failed to open '%s'\n", *inPath)
		os.Exit(1)
	}

	// Convert it to grayscale
	gray := pointu.ImageToGray(img)
	bounds := gray.Bounds()

	// Generate initial points by using importance sampling
	points := pointu.ImportanceSample(*nPoints, gray, uint8(*threshold), rng)

	// Apply Weighted Voronoi Stippling
	pdf := pointu.MakePDF(gray)
	step := 1 / *resolution
	stipples, densities := pointu.GetCentroids(points, pdf, bounds, step)
	for i := 1; i < *iterations-1; i++ {
		stipples, densities = pointu.GetCentroids(stipples, pdf, bounds, step)
	}

	// Normalize the densities to obtain radiuses with desired bounds
	radiuses := pointu.RescaleFloat64s(densities, *rMin, *rMax)

	// Start with a white image
	dc := gg.NewContext(img.Bounds().Dx(), img.Bounds().Dy())
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Draw each circle, with the initial image colors if specified
	if *withColor {
		for i, s := range stipples {
			dc.SetColor(img.At(int(s.X), int(s.Y)))
			dc.DrawCircle(s.X, s.Y, radiuses[i])
			dc.Fill()
		}
	} else {
		for i, s := range stipples {
			dc.DrawCircle(s.X, s.Y, radiuses[i])
		}
		dc.SetRGB(0, 0, 0)
		dc.Fill()
	}

	// Save the stippled image as a PNG file
	err = dc.SavePNG(*outPath)
	if err != nil {
		fmt.Printf("Failed to save output to '%s'\n", *outPath)
		os.Exit(1)
	}
}
