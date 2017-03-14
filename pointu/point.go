package pointu

import (
	"fmt"
	"sort"
)

// A Point represents an (X, Y) pair.
type Point struct {
	X float64
	Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

// getDist computes the squared euclidean distance to another Point.
func (p Point) getDist(q Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return dx*dx + dy*dy
}

type Points []Point

// sortByX sort points inplace by their x-axis.
func (pts Points) sortByX() {
	sort.Slice(pts, func(i, j int) bool { return pts[i].X < pts[j].X })
}

// sortByY sort Points inplace by their y-axis.
func (pts Points) sortByY() {
	sort.Slice(pts, func(i, j int) bool { return pts[i].Y < pts[j].Y })
}
