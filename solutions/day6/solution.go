package day6

import (
	"advent_of_code/utils"
	"regexp"
	"strconv"
	"strings"
)

func solution1(day string) string {
	pattern := regexp.MustCompile(`\s+`)
	lines := utils.GetInputPathLinesFormatted(day, func(line string) []string {
		line = strings.TrimSpace(line)
		return pattern.Split(line, -1)
	})

	neutrals, ops := OpGetter(lines[len(lines)-1])

	for _, line := range lines[:len(lines)-1] {
		for ix, val := range line {
			currVal, f := neutrals[ix], ops[ix]
			intVal, _ := strconv.Atoi(val)

			neutrals[ix] = f(currVal, intVal)
		}
	}

	total := 0
	for _, val := range neutrals {
		total += val
	}

	return strconv.Itoa(total)
}

func solution2(day string) string {
	lines := utils.GetInputPathLines(day)

	lastLine := strings.TrimSpace(lines[len(lines)-1])
	lastLineSplit := regexp.MustCompile(`\s+`).Split(lastLine, -1)

	neutrals, ops := OpGetter(lastLineSplit)
	charIx, total := len(lines[0])-1, 0
	for columnIx := len(lastLineSplit) - 1; columnIx >= 0; columnIx-- {
		currentColumn := []int{}
		for charIx >= 0 {
			toAdd, seen := 0, 0
			for lineIx := 2; lineIx <= len(lines); lineIx++ {
				line := lines[len(lines)-lineIx]
				val, err := strconv.Atoi(string(line[charIx]))

				if err != nil {
					val = 0
				} else {
					seen++
				}
				toAdd += val * utils.PowInt(10, seen-1)
			}

			charIx--
			if toAdd == 0 {
				break
			}

			currentColumn = append(currentColumn, toAdd)
		}

		currVal, f := neutrals[columnIx], ops[columnIx]
		for _, val := range currentColumn {
			currVal = f(currVal, val)
		}

		total += currVal
	}

	return strconv.Itoa(total)
}

func Solution(day string) (string, string) {
	return solution1(day), solution2(day)
}
