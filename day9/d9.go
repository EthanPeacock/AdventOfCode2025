package main

import (
	"aoc/internal/utils"
	"fmt"
	"slices"
	"strings"
)

type Location struct {
	x int
	y int
}

func getInput() []Location {
	rawInput := utils.ReadFileLines("input.txt")

	var locations []Location
	for _, tile := range rawInput {
		coords := strings.Split(tile, ",")
		x := utils.StringToInt(coords[0])
		y := utils.StringToInt(coords[1])
		locations = append(locations, Location{x, y})
	}

	return locations
}

func part1() int {
	tiles := getInput()

	biggestArea := 0
	for i, tileA := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tileB := tiles[j]
			dx := utils.Abs(tileA.x, tileB.x) + 1
			dy := utils.Abs(tileA.y, tileB.y) + 1
			area := dx * dy
			if area > biggestArea {
				biggestArea = area
			}
		}
	}

	return biggestArea
}

type Line struct {
	start Location
	end   Location
}

type Rectangle struct {
	start Location
	end   Location
	area  int
}

func IsPointInside(corner Location, edges []Line) bool {
	inside := false

	for _, edge := range edges {
		if CheckPointOnEdge(corner, edge) {
			return true // rest dont matter because it must be inside
		}

		// basically XOR, make sure corner.y is within vertical range of edge
		// this also ignores horizontal lines
		if (edge.start.y > corner.y) != (edge.end.y > corner.y) {
			if corner.x < edge.start.x {
				inside = !inside
			}
		}
	}

	return inside
}

func CheckPointOnEdge(p Location, edge Line) bool {
	minX := utils.Min(edge.start.x, edge.end.x)
	maxX := utils.Max(edge.start.x, edge.end.x)
	minY := utils.Min(edge.start.y, edge.end.y)
	maxY := utils.Max(edge.start.y, edge.end.y)

	if p.x < minX || p.x > maxX || p.y < minY || p.y > maxY {
		return false
	}

	// collinearity cross product
	crossProduct := (p.y-edge.start.y)*(edge.end.x-edge.start.x) - (p.x-edge.start.x)*(edge.end.y-edge.start.y)

	// if result is 0, the 3 points are on same line
	return crossProduct == 0
}

func IsRectangleIntersected(rectEdges []Line, polygonEdges []Line) bool {
	for _, rectEdge := range rectEdges {
		for _, polygonEdge := range polygonEdges {
			rectEdgeVertical := rectEdge.start.x == rectEdge.end.x
			polygonEdgeVertical := polygonEdge.start.x == polygonEdge.end.x

			var lineA Line
			var lineB Line

			if rectEdgeVertical && !polygonEdgeVertical {
				lineA = rectEdge
				lineB = polygonEdge
			} else if !rectEdgeVertical && polygonEdgeVertical {
				lineA = polygonEdge
				lineB = rectEdge
			}

			if rectEdgeVertical != polygonEdgeVertical {
				// check lineA.x is between lineB range
				xInRange := (lineA.start.x > utils.Min(lineB.start.x, lineB.end.x)) &&
					(lineA.start.x < utils.Max(lineB.start.x, lineB.end.x))

				// check lineB.y is between lineA range
				yInRange := (lineB.start.y > utils.Min(lineA.start.y, lineA.end.y)) &&
					(lineB.start.y < utils.Max(lineA.start.y, lineA.end.y))

				if xInRange && yInRange {
					return true
				}
			}
		}
	}

	return false
}

func part2() int {
	tiles := getInput()

	var edges []Line
	var rectangles []Rectangle
	for i, redTile1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			redTile2 := tiles[j]

			if redTile1.y == redTile2.y { // in same row, different col
				edges = append(edges, Line{redTile1, redTile2})
			}
			if redTile1.x == redTile2.x { // in same col, different row
				edges = append(edges, Line{redTile1, redTile2})
			}

			dx := utils.Abs(redTile1.x, redTile2.x) + 1
			dy := utils.Abs(redTile1.y, redTile2.y) + 1
			area := dx * dy
			rectangles = append(rectangles, Rectangle{redTile1, redTile2, area})
		}
	}

	slices.SortFunc(rectangles, func(a, b Rectangle) int {
		if a.area < b.area {
			return 1
		}
		if a.area > b.area {
			return -1
		}
		return 0
	})

	biggestArea := 0
	for _, rect := range rectangles {
		corner1 := rect.start
		corner2 := rect.end
		corner3 := Location{rect.start.x, rect.start.y + (rect.end.y - rect.start.y)}
		corner4 := Location{rect.start.x + (rect.end.x - rect.start.x), rect.start.y}

		corner1Within := IsPointInside(corner1, edges)
		corner2Within := IsPointInside(corner2, edges)
		corner3Within := IsPointInside(corner3, edges)
		corner4Within := IsPointInside(corner4, edges)

		if !corner1Within || !corner2Within || !corner3Within || !corner4Within {
			continue
		}

		rectangleEdges := []Line{
			{corner1, corner3},
			{corner3, corner2},
			{corner2, corner4},
			{corner4, corner1},
		}

		if !IsRectangleIntersected(rectangleEdges, edges) {
			biggestArea = rect.area
			break
		}
	}

	return biggestArea
}

func main() {
	p1Result := part1()
	p2Result := part2()

	fmt.Printf("Part 1: %d \n", p1Result)
	fmt.Printf("Part 2: %d \n", p2Result)
}
