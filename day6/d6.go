package main

import (
	"aoc/internal/utils"
	"fmt"
	"strings"
)

func getPart1Input() [][]string {
	rawInput := utils.ReadFileLines("input.txt")

	var lines [][]string
	for _, line := range rawInput {
		columns := strings.Fields(line)
		lines = append(lines, columns)
	}

	return lines
}

func part1() int {
	lines := getPart1Input()

	grandTotal := 0
	for c := range len(lines[0]) {
		var numbers []int
		var operator string
		for r := range len(lines) {
			item := lines[r][c]
			if item == "+" || item == "*" {
				operator = item
			} else {
				numbers = append(numbers, utils.StringToInt(item))
			}
		}

		total := 0
		for _, number := range numbers {
			switch operator {
			case "+":
				total += number
			case "*":
				total = utils.Max(1, total) * number
			}
		}
		grandTotal += total
	}

	return grandTotal
}

func getPart2Input() [][]string {
	rawInput := utils.ReadFileLines("input.txt")

	var cols [][]string

	lastColDivider := 0
	for col := range rawInput[0] {
		spaces := 0
		for row := range rawInput {
			if rawInput[row][col] == ' ' {
				spaces++
			}
		}

		lastCol := col == len(rawInput[0])-1
		if spaces == len(rawInput) || lastCol {
			var newCol []string
			for row := range rawInput {
				dividePoint := col
				if lastCol {
					dividePoint = col + 1
				}
				newCol = append(newCol, rawInput[row][lastColDivider:dividePoint])
			}

			cols = append(cols, newCol)
			lastColDivider = col + 1
		}
	}

	return cols
}

func part2() int {
	problems := getPart2Input()

	grandTotal := 0
	for i := len(problems) - 1; i >= 0; i-- {
		problem := problems[i]
		operator := strings.TrimSpace(problem[len(problem)-1])

		total := 0
		for c := range len(problem[0]) {
			number := 0
			for r := range len(problem) - 1 { // -1 to ignore the operator
				col := string(problem[r][c])
				if col != " " {
					number = (number * 10) + utils.StringToInt(col)
				}
			}

			switch operator {
			case "+":
				total += number
			case "*":
				total = utils.Max(1, total) * number
			}
		}

		grandTotal += total
	}

	return grandTotal
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
