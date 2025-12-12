package day10

import (
	"advent_of_code/utils"
	"strconv"
)

func solution1(inputs []input) string {
	total := 0
	for _, inp := range inputs {
		for _, sub := range generateAllSubsets(len(inp.buttons))[1:] {
			val := inp.buttons[sub[0]]
			for _, ix := range sub[1:] {
				val = val ^ inp.buttons[ix]
			}

			if val == inp.objective {
				total += len(sub)
				break
			}
		}
	}

	return strconv.Itoa(total)
}

func Solution(day string) (string, string) {
	inputs := utils.GetInputPathLinesFormatted(day, formatInput)

	return solution1(inputs), ""
}
