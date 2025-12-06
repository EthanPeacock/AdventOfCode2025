package main

import (
	"aoc/internal/utils"
	"fmt"
	"sort"
	"strings"
)

type FreshRange struct {
	min int
	max int
}

func getInput() ([]FreshRange, []int) {
	rawInput := utils.ReadFileLines("input.txt")

	emptyLineReached := false
	var fresh []FreshRange
	var available []int
	for _, line := range rawInput {
		if line == "" {
			emptyLineReached = true
			continue
		}

		if emptyLineReached {
			available = append(available, utils.StringToInt(line))
		} else {
			splitLine := strings.Split(line, "-")
			lineRange := FreshRange{utils.StringToInt(splitLine[0]), utils.StringToInt(splitLine[1])}
			fresh = append(fresh, lineRange)
		}
	}

	return fresh, available
}

func part1() int {
	fresh, available := getInput()

	freshAvailable := 0
	for _, ingredient := range available {
		for _, freshRange := range fresh {
			if ingredient >= freshRange.min && ingredient <= freshRange.max {
				freshAvailable++
				break
			}
		}
	}

	return freshAvailable
}

func part2() int {
	ranges, _ := getInput()
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	var mergedRanges []FreshRange
	currentRange := ranges[0]
	for _, nextRange := range ranges {
		// sorted so nextRange.min will always be >= currentRange.min
		if nextRange.min <= currentRange.max {
			currentRange.max = utils.Max(currentRange.max, nextRange.max)
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = nextRange
		}
	}
	mergedRanges = append(mergedRanges, currentRange)

	freshIngredients := 0
	for _, freshRange := range mergedRanges {
		freshIngredients += utils.Abs(freshRange.min, freshRange.max) + 1
	}

	return freshIngredients
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
