package day4

import (
	"advent_of_code/utils"
	"strconv"
)

func solution1(lines []string) string {
	total := 0
	for _, line := range countNeighbors(lines) {
		for _, count := range line {
			if count >= 0 && count < 4 {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}

func solution2(lines []string) string {
	initialAtCount := countAts(lines)
	for {
		newLines, changed := updateLines(lines)
		if !changed {
			break
		}
		lines = newLines
	}

	finalAtCount := countAts(lines)
	return strconv.Itoa(initialAtCount - finalAtCount)
}

func Solution(day string) (string, string) {
	lines := utils.GetInputPathLines(day)
	return solution1(lines), solution2(lines)
}
