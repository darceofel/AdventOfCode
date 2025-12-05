package day4

var neighborOffsets = [8][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countNeighbors(lines []string) [][]int {
	neighborCounts := [][]int{}
	for ix, line := range lines {
		currCounts := []int{}

		for jx, char := range line {
			if char == '.' {
				currCounts = append(currCounts, -1)
				continue
			}

			totalAts := 0
			for _, offset := range neighborOffsets {
				if ix+offset[0] < 0 || ix+offset[0] >= len(lines) ||
					jx+offset[1] < 0 || jx+offset[1] >= len(line) {
					continue
				}

				if lines[ix+offset[0]][jx+offset[1]] == '@' {
					totalAts++
				}
			}

			currCounts = append(currCounts, totalAts)
		}
		neighborCounts = append(neighborCounts, currCounts)
	}

	return neighborCounts
}

func updateLines(lines []string) ([]string, bool) {
	updatedLines := []string{}
	changed := false

	neighborCounts := countNeighbors(lines)
	for ix, line := range lines {
		newLine := ""
		for jx, char := range line {
			if char == '.' {
				newLine += "."
				continue
			}

			count := neighborCounts[ix][jx]
			if count < 4 {
				newLine += "."
				changed = true
			} else {
				newLine += "@"
			}
		}

		updatedLines = append(updatedLines, newLine)
	}

	return updatedLines, changed
}

func countAts(lines []string) int {
	atCount := 0
	for _, line := range lines {
		for _, char := range line {
			if char == '@' {
				atCount++
			}
		}
	}

	return atCount
}
