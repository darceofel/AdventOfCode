package day10

import (
	"advent_of_code/utils"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type input struct {
	objective int
	buttons   []int
}

var objectivePattern = regexp.MustCompile(`\[.+\]`)
var buttonsPattern = regexp.MustCompile(`\(.+\)`)

func formatInput(line string) input {
	stringObjective := objectivePattern.FindString(line)
	stringButtons := buttonsPattern.FindAllString(line, -1)

	buttons := []int{}
	for _, stringButton := range stringButtons {

		connections := strings.Split(stringButton, " ")
		for _, connection := range connections {
			connection = strings.ReplaceAll(connection, "(", "")
			connection = strings.ReplaceAll(connection, ")", "")

			total := 0
			for _, strNum := range strings.Split(connection, ",") {
				num, _ := strconv.Atoi(strNum)
				total += utils.PowInt(2, num)
			}

			buttons = append(buttons, total)
		}
	}

	objective := 0
	for ix, char := range stringObjective[1 : len(stringObjective)-1] {
		if char == '#' {
			objective += utils.PowInt(2, ix)
		}
	}

	return input{objective: objective, buttons: buttons}
}

func generateAllSubsets(n int) [][]int {
	total := 1 << n
	res := make([][]int, 0, total)

	for mask := 0; mask < total; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, i)
			}
		}
		res = append(res, subset)
	}

	sort.Slice(res, func(i, j int) bool {
		if len(res[i]) != len(res[j]) {
			return len(res[i]) < len(res[j])
		}
		for k := range res[i] {
			if res[i][k] != res[j][k] {
				return res[i][k] < res[j][k]
			}
		}
		return false
	})

	return res
}
