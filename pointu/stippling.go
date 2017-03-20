package pointu

import (
	"image"
)

func GetCentroids(sites Points, pdf PDF, bounds image.Rectangle, step float64) (Points, []float64) {

	var (
		siteCentroids   = make(map[Point]Point)
		siteIntensities = make(map[Point]float64)
		siteNPoints     = make(map[Point]float64)
	)

	for _, site := range sites {
		siteCentroids[site] = Point{}
	}

	var (
		box = rectangle{
			Point{float64(bounds.Min.X), float64(bounds.Min.Y)},
			Point{float64(bounds.Max.X), float64(bounds.Max.Y)},
		}
		kd = makeKdTree(sites, box)
	)

	for i := box.min.X; i < box.max.X; i += step {
		for j := box.min.Y; j < box.max.Y; j += step {
			p := Point{i, j}
			nn, _ := kd.findNearestNeighbour(p)
			w := pdf[int(i)][int(j)]
			centroid := siteCentroids[nn]
			centroid.X += w * p.X
			centroid.Y += w * p.Y
			siteCentroids[nn] = centroid
			siteIntensities[nn] += w
			siteNPoints[nn]++
		}
	}

	var (
		centroids = make(Points, len(sites))
		densities = make([]float64, len(sites))
		i         int
	)

	for site, density := range siteIntensities {
		centroid := siteCentroids[site]
		centroid.X /= density
		centroid.Y /= density
		centroids[i] = centroid
		densities[i] = siteIntensities[site] / siteNPoints[site]
		i++
	}

	return centroids, densities
}
