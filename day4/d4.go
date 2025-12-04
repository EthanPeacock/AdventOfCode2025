package main

import (
	"aoc/internal/utils"
	"fmt"
)

var FORKLIFT_MAX int = 4

type Direction struct {
	dx int
	dy int
}

var DIRECTIONS = [...]Direction{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

func part1() int {
	paperGrid := utils.ReadFileLineBytes("input.txt")

	total := 0
	for row := range paperGrid {
		for col := range paperGrid[row] {
			adjacentRolls := 0

			if paperGrid[row][col] != '@' {
				continue
			}

			for _, dir := range DIRECTIONS {
				adjacentRow := row + dir.dx
				adjacentCol := col + dir.dy
				if (adjacentRow < 0 || adjacentRow > len(paperGrid)-1) ||
					(adjacentCol < 0 || adjacentCol > len(paperGrid[row])-1) {
					continue
				}

				if paperGrid[adjacentRow][adjacentCol] == '@' {
					adjacentRolls++
				}

				if adjacentRolls == FORKLIFT_MAX {
					break
				}
			}

			if adjacentRolls < FORKLIFT_MAX {
				total++
			}
		}
	}

	return total
}

type RollPosition struct {
	row int
	col int
}

func part2() int {
	paperGrid := utils.ReadFileLineBytes("input.txt")

	total := 0

	for {
		var removedRolls []RollPosition

		for row := range paperGrid {
			for col := range paperGrid[row] {
				adjacentRolls := 0

				if paperGrid[row][col] != '@' {
					continue
				}

				for _, dir := range DIRECTIONS {
					adjacentRow := row + dir.dx
					adjacentCol := col + dir.dy
					if (adjacentRow < 0 || adjacentRow > len(paperGrid)-1) ||
						(adjacentCol < 0 || adjacentCol > len(paperGrid[row])-1) {
						continue
					}

					if paperGrid[adjacentRow][adjacentCol] == '@' {
						adjacentRolls++
					}

					if adjacentRolls == FORKLIFT_MAX {
						break
					}
				}

				if adjacentRolls < FORKLIFT_MAX {
					removedRolls = append(removedRolls, RollPosition{row, col})
					total++
				}
			}
		}

		if len(removedRolls) == 0 {
			break
		}

		for _, roll := range removedRolls {
			paperGrid[roll.row][roll.col] = 'x'
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
