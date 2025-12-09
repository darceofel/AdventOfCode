package day8

import (
	"advent_of_code/utils"
	"sort"
	"strconv"
	"strings"
)

var MAX_LENGTH int = 1000

func solution1(points []Point) string {
	dists := GetSortedDistances(points)
	dists = dists[:MAX_LENGTH]

	nodes := make(map[int]int)
	for _, dist := range dists {
		_, seenA := nodes[dist.aIx]
		_, seenB := nodes[dist.bIx]

		if !seenA && !seenB {
			nodes[dist.aIx] = dist.bIx
			nodes[dist.bIx] = dist.aIx
		} else if seenA && seenB {
			bPrev, shouldUpdate := GetPreviousNodeIx(nodes, dist.bIx, dist.aIx)

			if shouldUpdate {
				nodes[bPrev] = nodes[dist.aIx]
				nodes[dist.aIx] = dist.bIx
			}

		} else if seenA {
			prevIx, shouldUpdate := GetPreviousNodeIx(nodes, dist.aIx, dist.bIx)
			if shouldUpdate {
				nodes[prevIx] = dist.bIx
				nodes[dist.bIx] = dist.aIx
			}
		} else if seenB {
			prevIx, shouldUpdate := GetPreviousNodeIx(nodes, dist.bIx, dist.aIx)
			if shouldUpdate {
				nodes[prevIx] = dist.aIx
				nodes[dist.aIx] = dist.bIx
			}
		}
	}

	seenNodes, totalCircuits := make(map[int]bool), []int{}
	for ix := range nodes {
		currIx, circuitLength := ix, 0
		for !seenNodes[currIx] {
			circuitLength++
			seenNodes[currIx] = true
			currIx = nodes[currIx]
		}

		if circuitLength > 0 {
			totalCircuits = append(totalCircuits, circuitLength)
		}
	}

	sort.Ints(totalCircuits)
	res := 1
	for ix := 1; ix <= 3; ix++ {
		res *= totalCircuits[len(totalCircuits)-ix]
	}
	return strconv.Itoa(res)
}

func solution2(points []Point) string {
	totalNodes, dists := len(points), GetSortedDistances(points)

	nodes, circuitCount, nodesCount := make(map[int]int), 0, 0
	res := 0
	for _, dist := range dists {
		_, seenA := nodes[dist.aIx]
		_, seenB := nodes[dist.bIx]

		if !seenA && !seenB {
			nodes[dist.aIx] = dist.bIx
			nodes[dist.bIx] = dist.aIx
			nodesCount += 2
			circuitCount++
		} else if seenA && seenB {
			bPrev, shouldUpdate := GetPreviousNodeIx(nodes, dist.bIx, dist.aIx)

			if shouldUpdate {
				nodes[bPrev] = nodes[dist.aIx]
				nodes[dist.aIx] = dist.bIx
				circuitCount--
			}

		} else if seenA {
			prevIx, _ := GetPreviousNodeIx(nodes, dist.aIx, -1)

			nodes[prevIx] = dist.bIx
			nodes[dist.bIx] = dist.aIx
			nodesCount++
		} else if seenB {
			prevIx, _ := GetPreviousNodeIx(nodes, dist.bIx, -1)

			nodes[prevIx] = dist.aIx
			nodes[dist.aIx] = dist.bIx
			nodesCount++
		}

		if (nodesCount == totalNodes) && (circuitCount == 1) {
			a := points[dist.aIx]
			b := points[dist.bIx]

			res = a.X * b.X
			break
		}
	}

	return strconv.Itoa(res)
}

func Solution(day string) (string, string) {
	points := utils.GetInputPathLinesFormatted(day, func(line string) Point {
		coords := strings.Split(line, ",")
		X, Y, Z := coords[0], coords[1], coords[2]

		xVal, _ := strconv.Atoi(X)
		yVal, _ := strconv.Atoi(Y)
		zVal, _ := strconv.Atoi(Z)

		return Point{
			X: xVal,
			Y: yVal,
			Z: zVal,
		}
	})

	return solution1(points), solution2(points)
}
