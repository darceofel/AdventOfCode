package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetInputPathContent(day string) string {
	filePath := fmt.Sprintf("inputs/day%s.txt", day)
	if day == "-1" {
		filePath = "inputs/_test.txt"
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	return string(file)
}

func GetInputPathLines(day string) []string {
	content := GetInputPathContent(day)

	return strings.Split(strings.TrimSpace(content), "\n")
}

func GetInputPathLinesFormatted[T any](day string, f func(s string) T) []T {
	result := []T{}
	for _, line := range GetInputPathLines(day) {
		result = append(result, f(line))
	}

	return result
}
