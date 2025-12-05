package day3

import (
	"advent_of_code/utils"
	"strconv"
)

func solution(lines [][]int, length int) string {
	total := 0
	for _, line := range lines {
		total += getMaxComb(line, length)
	}

	return strconv.Itoa(total)
}

func Solution(inputNum string) (string, string) {
	lines := utils.GetInputPathLinesFormatted(inputNum, func(s string) []int {
		res := []int{}
		for _, char := range s {
			val, _ := strconv.Atoi(string(char))
			res = append(res, val)
		}
		return res
	})

	return solution(lines, 2), solution(lines, 12)
}
