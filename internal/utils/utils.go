package utils

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func ReadFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileLines
}

func StringToInt(x string) int {
	xInt, err := strconv.Atoi(x)

	if err != nil {
		log.Fatal(err)
	}

	return xInt
}

func Abs(x int, y int) int {
	return int(math.Abs(float64(x - y)))
}

func Max(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// cba to implement set data structure, maybe grab something already made for actual AOC 2025?
func AppendIfMissing[T comparable](slice []T, item T) []T {
	if slices.Contains(slice, item) {
		return slice
	}

	return append(slice, item)
}
