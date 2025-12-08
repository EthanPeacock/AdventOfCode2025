package main

import (
	"aoc/internal/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

type JunctionBox struct {
	x int
	y int
	z int
}

type BoxDistance struct {
	boxIndex        int
	compareBoxIndex int
	distance        float64
}

func getInput() []JunctionBox {
	rawInput := utils.ReadFileLines("input.txt")

	var junctionBoxes []JunctionBox
	for _, boxPosition := range rawInput {
		position := strings.Split(boxPosition, ",")
		box := JunctionBox{utils.StringToInt(position[0]), utils.StringToInt(position[1]), utils.StringToInt(position[2])}
		junctionBoxes = append(junctionBoxes, box)
	}

	return junctionBoxes
}

func calculateDistance(box1 JunctionBox, box2 JunctionBox) float64 {
	dx := box1.x - box2.x
	dy := box1.y - box2.y
	dz := box1.z - box2.z

	distance := math.Sqrt(float64(dx*dx + dy*dy + dz*dz))

	return distance
}

func isInACircuit(circuits [][]int, targetBox int) (bool, int) {
	for circuitIndex, circuit := range circuits {
		for _, box := range circuit {
			if box == targetBox {
				return true, circuitIndex
			}
		}
	}

	return false, -1
}

func part1() int {
	junctionBoxes := getInput()

	var distances []BoxDistance
	for i, box := range junctionBoxes {
		// just look ahead, previous will alreayd be done
		// e.g., A -> B, A -> C when you look at B or C you dont need to recompare with A etc
		for j := i + 1; j < len(junctionBoxes); j++ {
			if j == i {
				continue
			}
			compareBox := junctionBoxes[j]
			dist := calculateDistance(box, compareBox)
			distances = append(distances, BoxDistance{i, j, dist})
		}
	}

	slices.SortFunc(distances, func(a, b BoxDistance) int {
		if a.distance < b.distance {
			return -1
		}
		if a.distance > b.distance {
			return 1
		}
		return 0
	})

	iterationsLimit := 1000 // 10 for sample input, 1000 for actual
	var circuits [][]int
	for i, potentialConnection := range distances {
		if i == iterationsLimit {
			break
		}

		box := potentialConnection.boxIndex
		closeBox := potentialConnection.compareBoxIndex

		boxInCircuit, boxCircuitIndex := isInACircuit(circuits, box)
		closeBoxInCircuit, closeBoxCircuitIndex := isInACircuit(circuits, closeBox)

		if boxInCircuit && closeBoxInCircuit {
			if boxCircuitIndex == closeBoxCircuitIndex {
				continue // in same circuit already
			}
			circuits[boxCircuitIndex] = append(circuits[boxCircuitIndex], circuits[closeBoxCircuitIndex]...)
			circuits = slices.Delete(circuits, closeBoxCircuitIndex, closeBoxCircuitIndex+1)
		} else if boxInCircuit && !closeBoxInCircuit {
			circuits[boxCircuitIndex] = append(circuits[boxCircuitIndex], closeBox)
		} else if !boxInCircuit && closeBoxInCircuit {
			circuits[closeBoxCircuitIndex] = append(circuits[closeBoxCircuitIndex], box)
		} else {
			circuits = append(circuits, []int{box, closeBox})
		}
	}

	slices.SortFunc(circuits, func(a, b []int) int {
		if len(a) < len(b) {
			return 1
		}
		if len(a) > len(b) {
			return -1
		}
		return 0
	})

	total := 1
	for i, circuit := range circuits {
		total *= len(circuit)

		if i == 2 {
			break
		}
	}

	return total
}

func part2() int {
	junctionBoxes := getInput()

	var distances []BoxDistance
	for i, box := range junctionBoxes {
		// just look ahead, previous will alreayd be done
		// e.g., A -> B, A -> C when you look at B or C you dont need to recompare with A etc
		for j := i + 1; j < len(junctionBoxes); j++ {
			if j == i {
				continue
			}
			compareBox := junctionBoxes[j]
			dist := calculateDistance(box, compareBox)
			distances = append(distances, BoxDistance{i, j, dist})
		}
	}

	slices.SortFunc(distances, func(a, b BoxDistance) int {
		if a.distance < b.distance {
			return -1
		}
		if a.distance > b.distance {
			return 1
		}
		return 0
	})

	var circuits [][]int
	var lastBox JunctionBox
	var lastConnectingBox JunctionBox
	for _, potentialConnection := range distances {
		box := potentialConnection.boxIndex
		closeBox := potentialConnection.compareBoxIndex

		boxInCircuit, boxCircuitIndex := isInACircuit(circuits, box)
		closeBoxInCircuit, closeBoxCircuitIndex := isInACircuit(circuits, closeBox)

		if boxInCircuit && closeBoxInCircuit {
			if boxCircuitIndex == closeBoxCircuitIndex {
				continue // in same circuit already
			}
			circuits[boxCircuitIndex] = append(circuits[boxCircuitIndex], circuits[closeBoxCircuitIndex]...)
			circuits = slices.Delete(circuits, closeBoxCircuitIndex, closeBoxCircuitIndex+1)
		} else if boxInCircuit && !closeBoxInCircuit {
			circuits[boxCircuitIndex] = append(circuits[boxCircuitIndex], closeBox)
		} else if !boxInCircuit && closeBoxInCircuit {
			circuits[closeBoxCircuitIndex] = append(circuits[closeBoxCircuitIndex], box)
		} else {
			circuits = append(circuits, []int{box, closeBox})
		}

		lastBox = junctionBoxes[box]
		lastConnectingBox = junctionBoxes[closeBox]
	}

	return lastBox.x * lastConnectingBox.x
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
