package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
)

func part1() int {
	input := utils.ReadFileLineBytes("input.txt")

	startLocationRegex, _ := regexp.Compile("S")
	startLocation := startLocationRegex.FindStringIndex(string(input[0]))

	beams := []int{startLocation[0]}
	splits := 0
	for i := 1; i < len(input); i++ {
		line := input[i]

		var newBeams []int
		for _, beam := range beams {
			switch line[beam] {
			case '.':
				line[beam] = '|'
				newBeams = append(newBeams, beam)
			case '^':
				left := beam - 1
				right := beam + 1

				if left >= 0 {
					line[left] = '|'
					newBeams = append(newBeams, left)
				}
				if right < len(line) {
					line[right] = '|'
					newBeams = append(newBeams, right)
				}

				splits++
			}
		}

		beams = newBeams
	}

	return splits
}

func part2() int {
	input := utils.ReadFileLineBytes("input.txt")
	var splitterRegex, _ = regexp.Compile(`\^`)

	var grid [][]byte
	for _, line := range input {
		if splitterRegex.Match(line) {
			grid = append(grid, line)
		}
	}

	startSplitterLoc := splitterRegex.FindStringIndex(string(grid[0]))

	beamCounts := make(map[int]int)
	beamCounts[startSplitterLoc[0]] = 1

	for r := 0; r < len(grid); r++ {
		nextRowCounts := make(map[int]int)

		for col, count := range beamCounts {
			if col < 0 || col >= len(grid[r]) {
				continue
			}

			if grid[r][col] == '^' {
				nextRowCounts[col-1] += count
				nextRowCounts[col+1] += count
			} else {
				nextRowCounts[col] += count
			}
		}

		beamCounts = nextRowCounts
	}

	total := 0
	for _, count := range beamCounts {
		total += count
	}

	return total
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
