package main

import (
	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

type Point struct {
	x, y int
}

type Region struct {
	plots     map[Point]bool
	letter    rune
	area      int
	perimeter int
}

var directions = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func findRegions(garden [][]rune) []Region {
	visited := make(map[Point]bool)
	var regions []Region

	for y, row := range garden {
		for x := range row {
			if !visited[Point{x, y}] {
				region := exploreRegion(garden, x, y, visited)

				if region != nil {
					regions = append(regions, *region)
				}
			}
		}
	}

	return regions
}

func exploreRegion(garden [][]rune, startX, startY int, visited map[Point]bool) *Region {
	startPoint := Point{startX, startY}

	if visited[startPoint] {
		return nil
	}

	letter := garden[startY][startX]
	plots := make(map[Point]bool)
	queue := []Point{startPoint}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] || garden[current.y][current.x] != letter {
			continue
		}

		visited[current] = true
		plots[current] = true

		addAdjacentPoints(garden, current, letter, visited, &queue)
	}

	perimeter := calculatePerimeter(garden, plots, letter)

	return &Region{
		plots:     plots,
		letter:    letter,
		area:      len(plots),
		perimeter: perimeter,
	}
}

func isValidPoint(garden [][]rune, x, y int, letter rune, visited map[Point]bool) bool {
	return x >= 0 && x < len(garden[0]) &&
		y >= 0 && y < len(garden) &&
		garden[y][x] == letter &&
		!visited[Point{x, y}]
}

func addAdjacentPoints(garden [][]rune, current Point, letter rune, visited map[Point]bool, queue *[]Point) {
	for _, dir := range directions {
		newX, newY := current.x+dir.x, current.y+dir.y

		if isValidPoint(garden, newX, newY, letter, visited) {
			*queue = append(*queue, Point{newX, newY})
		}
	}
}

func isEdgeOrDifferentLetter(garden [][]rune, x, y int, letter rune) bool {
	return x < 0 || x >= len(garden[0]) ||
		y < 0 || y >= len(garden) ||
		garden[y][x] != letter
}

func calculatePerimeter(garden [][]rune, plots map[Point]bool, letter rune) int {
	perimeter := 0

	for plot := range plots {
		for _, dir := range directions {
			newX, newY := plot.x+dir.x, plot.y+dir.y

			if isEdgeOrDifferentLetter(garden, newX, newY, letter) {
				perimeter++
			}
		}
	}

	return perimeter
}

func solvePart1(garden [][]rune) int {
	totalCost := 0

	for _, region := range findRegions(garden) {
		totalCost += region.area * region.perimeter
	}

	return totalCost
}

func main() {
	garden := utils.GetInputFileContentAs2DArray("day12/input.txt")
	println("Part 1:", solvePart1(garden))
}
