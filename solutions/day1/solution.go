package day1

import (
	"advent_of_code/utils"
	"fmt"
	"strconv"
)

var initial_dial_position int = 50

func solution1(lines []string) string {
	position, password := initial_dial_position, 0
	for _, line := range lines {
		direction, value := line[0], line[1:]

		intVal, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting value to integer:", err)
			return ""
		}

		intDirection := 1
		if direction == 'L' {
			intDirection = -1
		}

		position = (position + intDirection*intVal + 100) % 100
		if position == 0 {
			password++
		}
	}

	return strconv.Itoa(password)
}

func solution2(lines []string) string {
	position, password := initial_dial_position, 0
	for _, line := range lines {
		direction, value := line[0], line[1:]

		intVal, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting value to integer:", err)
			return ""
		}

		intDirection := 1
		if direction == 'L' {
			intDirection = -1
		}

		nextPosition := position + intDirection*intVal
		if (position > 0) && (nextPosition <= 0) {
			password++
		}

		password += intDirection * nextPosition / 100
		position = (nextPosition%100 + 100) % 100
	}

	return strconv.Itoa(password)
}

func Solution(inputNum string) (string, string) {
	lines := utils.GetInputPathLines(inputNum)

	return solution1(lines), solution2(lines)
}
