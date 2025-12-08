package day5

import (
	"advent_of_code/utils"
	"math"
	"strconv"
)

func solution1(lines []string) string {
	insideInterval, totalFresh := 0, 0

	ingredients := formatIngredients(lines, false)
	for _, ingredient := range ingredients {
		if ingredient[1] == -1 {
			insideInterval++
		} else if ingredient[1] == 1 {
			insideInterval--
		} else if insideInterval > 0 {
			totalFresh++
		}
	}

	return strconv.Itoa(totalFresh)
}

func solution2(lines []string) string {
	ingredients := formatIngredients(lines, true)

	insideInterval, lowerBound, freshCount := 0, math.MaxInt, 0
	for _, ingredient := range ingredients {
		val, typ := ingredient[0], ingredient[1]
		if typ == -1 {
			insideInterval++
			lowerBound = min(val, lowerBound)
			continue
		}

		insideInterval--
		if insideInterval == 0 {
			freshCount += val - lowerBound + 1
			lowerBound = math.MaxInt
		}
	}

	return strconv.Itoa(freshCount)
}

func Solution(day string) (string, string) {
	lines := utils.GetInputPathLines(day)

	return solution1(lines), solution2(lines)
}
