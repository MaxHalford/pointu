package pointu

import (
	"image"
)

func GetCentroids(sites Points, pdf PDF, bounds image.Rectangle, step float64) (Points, []float64) {
	var box = rectangle{
		Point{float64(bounds.Min.X), float64(bounds.Min.Y)},
		Point{float64(bounds.Max.X), float64(bounds.Max.Y)},
	}

	var siteCentroids = make(map[Point]Point)
	var siteDensities = make(map[Point]float64)
	var siteAreas = make(map[Point]float64)

	for _, site := range sites {
		siteCentroids[site] = Point{}
	}

	var kd = makeKdTree(sites, box)

	for i := box.min.X; i < box.max.X; i += step {
		for j := box.min.Y; j < box.max.Y; j += step {
			p := Point{i, j}
			nn, _, _ := kd.nearest(p)
			w := pdf[int(i)][int(j)]
			centroid := siteCentroids[nn]
			centroid.X += w * p.X
			centroid.Y += w * p.Y
			siteCentroids[nn] = centroid
			siteDensities[nn] += w
			siteAreas[nn]++
		}
	}

	var centroids = make(Points, len(sites))
	var densities = make([]float64, len(sites))
	var i int
	for site, density := range siteDensities {
		centroid := siteCentroids[site]
		centroid.X /= density
		centroid.Y /= density
		centroids[i] = centroid
		densities[i] = siteDensities[site] / siteAreas[site]
		i++
	}

	return centroids, densities
}
