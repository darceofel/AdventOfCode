package day8

import (
	"advent_of_code/utils"
	"sort"
)

type Point struct {
	X int
	Y int
	Z int
}

type Distance struct {
	aIx    int
	bIx    int
	sqDist int
}

func SquaredDist(a, b Point) int {
	return utils.PowInt(a.X-b.X, 2) + utils.PowInt(a.Y-b.Y, 2) + utils.PowInt(a.Z-b.Z, 2)
}

func GetDistance(points []Point, aIx, bIx int) Distance {
	return Distance{
		aIx:    aIx,
		bIx:    bIx,
		sqDist: SquaredDist(points[aIx], points[bIx]),
	}
}

func GetPreviousNodeIx(nodes map[int]int, nodeIx, otherIx int) (int, bool) {
	node, currentIx := nodes[nodeIx], nodeIx
	for node != nodeIx {
		if node == otherIx {
			return -1, false
		}

		currentIx = node
		node = nodes[node]
	}

	return currentIx, true
}

func GetSortedDistances(points []Point) []Distance {
	dists := []Distance{}
	for ix := 0; ix < len(points)-1; ix++ {
		for jx := ix + 1; jx < len(points); jx++ {
			dists = append(dists, GetDistance(points, ix, jx))
		}
	}

	sort.Slice(dists, func(i, j int) bool {
		return dists[i].sqDist < dists[j].sqDist
	})

	return dists
}
