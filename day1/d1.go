package main

import (
	"aoc/internal/utils"
	"fmt"
)

var DIAL_CLICKS int = 100

type Rotation struct {
	direction byte
	distance  int
}

func getInput() []Rotation {
	input := utils.ReadFileLines("input.txt")

	var instructions []Rotation
	for _, line := range input {
		dir := line[0]
		dist := utils.StringToInt(line[1:])
		instructions = append(instructions, Rotation{dir, dist})
	}

	return instructions
}

func part1() int {
	instructions := getInput()

	dial := 50
	count := 0
	for _, rotation := range instructions {
		modifier := rotation.distance
		if rotation.direction == 'L' {
			modifier = -modifier
		}

		newDialPosition := (dial + modifier) % DIAL_CLICKS
		if newDialPosition < 0 { // if rem is negative, wrap around (code mod != math mod)
			newDialPosition += DIAL_CLICKS
		}

		dial = newDialPosition
		if dial == 0 {
			count++
		}
	}

	return count
}

func part2() int {
	instructions := getInput()

	dial := 50
	count := 0
	for _, rotation := range instructions {
		// gave up on remainder stuff, must simpler to simulate rotations
		for range rotation.distance {
			if rotation.direction == 'L' {
				dial -= 1
			} else {
				dial += 1
			}

			dial = dial % DIAL_CLICKS

			if dial == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
