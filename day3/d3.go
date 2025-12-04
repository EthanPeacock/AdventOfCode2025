package main

import (
	"aoc/internal/utils"
	"fmt"
	"strings"
)

func getInput() [][]int {
	rawBanks := utils.ReadFileLines("input.txt")

	var banks [][]int
	for _, bank := range rawBanks {
		stringBatteries := strings.Split(bank, "")
		var batteries []int
		for _, battery := range stringBatteries {
			batteries = append(batteries, utils.StringToInt(battery))
		}
		banks = append(banks, batteries)
	}

	return banks
}

func part1() int {
	banks := getInput()

	total := 0
	for _, bank := range banks {
		firstDigit := 0
		secondDigit := 0

		for index, battery := range bank {
			if battery > firstDigit && index < len(bank)-1 {
				firstDigit = battery
				secondDigit = 0
			} else if battery > secondDigit {
				secondDigit = battery
			}
		}

		joltage := (firstDigit * 10) + secondDigit
		total += joltage
	}

	return total
}

var BATTERIES_REQUIRED int = 12

func part2() int {
	banks := getInput()

	total := 0
	for _, bank := range banks {
		digits := make([]int, 0, BATTERIES_REQUIRED)

		for index, battery := range bank {
			remainingInBank := len(bank) - 1 - index

			for len(digits) > 0 && // previous digits exist
				digits[len(digits)-1] < battery && // previous digit is smaller than current battery
				(len(digits)+remainingInBank) >= BATTERIES_REQUIRED { // bank has enough left to fill required 12

				digits = digits[:len(digits)-1]
			}

			if len(digits) < BATTERIES_REQUIRED {
				digits = append(digits, battery)
			}
		}

		digit := 0
		mult := 1
		for d := BATTERIES_REQUIRED - 1; d >= 0; d-- {
			digit += digits[d] * mult
			mult *= 10
		}

		total += digit
	}

	return total
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
