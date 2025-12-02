package main

import (
	"aoc/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func getInput() [][]int {
	input := utils.ReadFile("input.txt")
	input = strings.TrimSuffix(input, "\n")
	rangesInput := strings.Split(input, ",")

	var ranges [][]int
	for _, r := range rangesInput {
		startAndEnd := strings.Split(r, "-")
		start := utils.StringToInt(startAndEnd[0])
		end := utils.StringToInt(startAndEnd[1])
		ranges = append(ranges, []int{start, end})
	}

	return ranges
}

func part1() int {
	ranges := getInput()

	total := 0
	for _, idRange := range ranges {
		start := idRange[0]
		end := idRange[1]

		for number := start; number <= end; number++ {
			numberString := strconv.Itoa(number)
			numberLength := len(numberString)

			if (numberLength % 2) != 0 {
				continue
			}

			midpoint := numberLength / 2
			firstHalf := numberString[:midpoint]
			secondHalf := numberString[midpoint:]

			if firstHalf == secondHalf {
				total += number
			}
		}
	}

	return total
}

func part2() int {
	ranges := getInput()

	total := 0
	for _, idRange := range ranges {
		start := idRange[0]
		end := idRange[1]

		for number := start; number <= end; number++ {
			numberString := strconv.Itoa(number)
			numberLength := len(numberString)

			midpoint := numberLength / 2
			for index := range numberString[:midpoint] {
				val := numberString[:index+1]
				repetitions := numberLength / len(val)

				if strings.Repeat(val, repetitions) == numberString {
					total += number
					break
				}
			}
		}
	}

	return total
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
