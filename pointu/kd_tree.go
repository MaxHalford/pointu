package pointu

import (
	"fmt"
	"math"
)

// A rectangle has an upper left point (Min) and a lower right point (Max).
type rectangle struct {
	min, max Point
}

type kdNode struct {
	p           Point
	splitByX    bool
	left, right *kdNode
}

type kdTree struct {
	root   *kdNode
	bounds rectangle
}

func makeKdTree(pts Points, bounds rectangle) kdTree {

	var split func(pts Points, splitByX bool) *kdNode

	split = func(pts Points, splitByX bool) *kdNode {

		if len(pts) == 0 {
			return nil
		}

		if splitByX {
			pts.sortByX()
		} else {
			pts.sortByY()
		}

		med := len(pts) / 2

		return &kdNode{
			p:        pts[med],
			splitByX: splitByX,
			left:     split(pts[:med], !splitByX),
			right:    split(pts[med+1:], !splitByX),
		}

	}

	return kdTree{
		root:   split(pts, true),
		bounds: bounds,
	}
}

func nn(node *kdNode, target Point, r rectangle, maxDistSqd float64) (nearest Point, distSqd float64, nodesVisited int) {

	if node == nil {
		return Point{}, math.Inf(1), 0
	}

	nodesVisited++

	var targetInLeft bool
	var leftBox, rightBox rectangle
	var nearestNode, furthestNode *kdNode
	var nearestBox, furthestBox rectangle

	if node.splitByX {
		leftBox = rectangle{r.min, Point{node.p.X, r.max.Y}}
		rightBox = rectangle{Point{node.p.X, r.min.Y}, r.max}
		targetInLeft = target.X <= node.p.X
	} else {
		leftBox = rectangle{r.min, Point{r.max.X, node.p.Y}}
		rightBox = rectangle{Point{r.min.X, node.p.Y}, r.max}
		targetInLeft = target.Y <= node.p.Y
	}

	if targetInLeft {
		nearestNode, nearestBox = node.left, leftBox
		furthestNode, furthestBox = node.right, rightBox
	} else {
		nearestNode, nearestBox = node.right, rightBox
		furthestNode, furthestBox = node.left, leftBox
	}

	var nv int
	nearest, distSqd, nv = nn(nearestNode, target, nearestBox, maxDistSqd)
	nodesVisited += nv
	if distSqd < maxDistSqd {
		maxDistSqd = distSqd
	}

	var d float64
	if node.splitByX {
		d = node.p.X - target.X
	} else {
		d = node.p.Y - target.Y
	}
	d *= d

	if d > maxDistSqd {
		return
	}

	if d = node.p.getDist(target); d < distSqd {
		nearest = node.p
		distSqd = d
		maxDistSqd = distSqd
	}

	tmpNearest, tmpSqd, nv := nn(furthestNode, target, furthestBox, maxDistSqd)
	nodesVisited += nv
	if tmpSqd < distSqd {
		nearest = tmpNearest
		distSqd = tmpSqd
	}
	return
}

func (t kdTree) nearest(p Point) (best Point, bestSqd float64, nv int) {
	return nn(t.root, p, t.bounds, math.Inf(1))
}

func showNearest(heading string, kd kdTree, p Point) {
	fmt.Println()
	fmt.Println(heading)
	fmt.Println("point:           ", p)
	nn, ssq, nv := kd.nearest(p)
	fmt.Println("nearest neighbor:", nn)
	fmt.Println("distance:        ", math.Sqrt(ssq))
	fmt.Println("nodes visited:   ", nv)
}
