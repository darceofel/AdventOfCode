package day5

import (
	"sort"
	"strconv"
	"strings"
)

func formatIngredients(lines []string, onlyRanges bool) [][]int {
	ingredients, seenBreak := [][]int{}, false
	for _, line := range lines {
		if line == "" {
			if onlyRanges {
				break
			}

			seenBreak = true
			continue
		}

		if seenBreak {
			val, _ := strconv.Atoi(line)
			ingredients = append(ingredients, []int{val, 0})
			continue
		}

		bounds := strings.Split(line, "-")
		a, _ := strconv.Atoi(bounds[0])
		b, _ := strconv.Atoi(bounds[1])

		ingredients = append(ingredients, []int{a, -1})
		ingredients = append(ingredients, []int{b, 1})
	}

	sort.Slice(ingredients, func(i, j int) bool {
		if ingredients[i][0] == ingredients[j][0] {
			return ingredients[i][1] < ingredients[j][1]
		}
		return ingredients[i][0] < ingredients[j][0]
	})

	return ingredients
}
