package pointu

import (
	"image"
	"math/rand"
	"sort"
)

func ImportanceSample(n int, gray *image.Gray, threshold uint8, rng *rand.Rand) Points {
	var (
		pts    = make(Points, n)
		bounds = gray.Bounds()
		hist   = make(map[uint8]Points)
	)
	// Build a histogram
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var intensity = gray.GrayAt(x, y).Y
			if intensity <= threshold {
				hist[intensity] = append(hist[intensity], Point{float64(x), float64(y)})
			}
		}
	}
	// Build a roulette wheel
	var roulette = make([]int, threshold+1)
	roulette[0] = 256 * len(hist[0])
	for i := 1; i < len(roulette); i++ {
		roulette[i] = roulette[i-1] + (256-i)*len(hist[uint8(i)])
	}
	// Run the wheel
	for i := range pts {
		var (
			ball = rng.Intn(roulette[len(roulette)-1])
			bin  = uint8(sort.SearchInts(roulette, ball))
			p    = hist[bin][rng.Intn(len(hist[bin]))]
		)
		p.X += rng.Float64()
		p.Y += rng.Float64()
		pts[i] = p
	}
	return pts
}
