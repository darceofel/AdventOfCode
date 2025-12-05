package day2

import (
	"advent_of_code/utils"
	"math"
	"strconv"
)

func isInvalid(n string) bool {
	length := len(n)
	if length%2 == 1 {
		return false
	}
	left, right := n[:length/2], n[length/2:]
	return left == right
}

func lowerInvalidsSum(n string, maxReps int) int {
	length, upperBound := len(n), 0
	if length%2 == 1 {
		upperBound = utils.PowInt(10, length/2) - 1
	} else {
		upperBound, _ = strconv.Atoi(n[:length/2])
		right, _ := strconv.Atoi(n[length/2:])
		if upperBound > right {
			upperBound--
		}
	}

	hardLim, _ := strconv.Atoi(n)
	total, seen := 0, map[int]bool{}
	for i := 1; i <= upperBound; i++ {
		log := int(math.Floor(math.Log10(float64(i)))) + 1
		r := utils.PowInt(10, log) + 1
		for rep := 1; rep < maxReps; rep++ {
			if r*i > hardLim {
				break
			} else if !seen[r*i] {
				total += r * i
				seen[r*i] = true
			}

			r = r*utils.PowInt(10, log) + 1
		}
	}

	return total
}
