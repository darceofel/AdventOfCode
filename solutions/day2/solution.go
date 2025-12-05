package day2

import (
	"advent_of_code/utils"
	"strconv"
	"strings"
)

func solution1(input string) string {
	ranges := strings.Split(input, ",")

	totalInvalids := 0
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		a, b := bounds[0], bounds[1]

		currentInvalids := lowerInvalidsSum(b, 2) - lowerInvalidsSum(a, 2)
		if isInvalid(a) {
			aVal, _ := strconv.Atoi(a)
			currentInvalids += aVal
		}

		totalInvalids += currentInvalids
	}

	return strconv.Itoa(totalInvalids)
}

func solution2(input string) string {
	maxReps := 100000000
	ranges := strings.Split(input, ",")

	totalInvalids := 0
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		a, b := bounds[0], bounds[1]

		currentInvalids := lowerInvalidsSum(b, maxReps) - lowerInvalidsSum(a, maxReps)
		if isInvalid(a) {
			aVal, _ := strconv.Atoi(a)
			currentInvalids += aVal
		}

		totalInvalids += currentInvalids
	}

	return strconv.Itoa(totalInvalids)
}

func Solution(inputNum string) (string, string) {
	input := utils.GetInputPathContent(inputNum)

	return solution1(input), solution2(input)
}
