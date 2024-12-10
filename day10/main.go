package main

import (
	"fmt"
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

type Point struct {
	x, y int
}

func findTrailheads(heightMap [][]int) []Point {
	trailheads := []Point{}
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			if heightMap[y][x] == 0 {
				trailheads = append(trailheads, Point{x, y})
			}
		}
	}
	return trailheads
}

func solvePart1(heightMap [][]int) int {
	trailheads := findTrailheads(heightMap)
	totalScore := 0
	for _, trailhead := range trailheads {
		totalScore += calculateTrailheadScore(heightMap, trailhead)
	}
	return totalScore
}

func calculateTrailheadScore(heightMap [][]int, start Point) int {
	visited := make(map[Point]bool)
	reachedPeaks := make(map[Point]bool)

	var dfs func(current Point, prevHeight int)
	dfs = func(current Point, prevHeight int) {
		if current.y < 0 || current.y >= len(heightMap) ||
			current.x < 0 || current.x >= len(heightMap[current.y]) ||
			visited[current] {
			return
		}

		currentHeight := heightMap[current.y][current.x]

		if prevHeight != -1 && currentHeight != prevHeight+1 {
			return
		}

		visited[current] = true

		if currentHeight == 9 {
			reachedPeaks[current] = true
		}

		directions := []Point{
			{current.x, current.y - 1},
			{current.x, current.y + 1},
			{current.x - 1, current.y},
			{current.x + 1, current.y},
		}

		for _, next := range directions {
			dfs(next, currentHeight)
		}
	}

	dfs(start, -1)

	return len(reachedPeaks)
}

func solvePart2(heightMap [][]int) int {
	trailheads := findTrailheads(heightMap)
	totalRating := 0
	for _, trailhead := range trailheads {
		totalRating += calculateTrailheadRating(heightMap, trailhead)
	}
	return totalRating
}

func calculateTrailheadRating(heightMap [][]int, start Point) int {
	ratingCount := 0

	var dfs func(current Point, prevHeight int)
	dfs = func(current Point, prevHeight int) {
		if current.y < 0 || current.y >= len(heightMap) ||
			current.x < 0 || current.x >= len(heightMap[current.y]) {
			return
		}

		currentHeight := heightMap[current.y][current.x]

		if prevHeight != -1 && currentHeight != prevHeight+1 {
			return
		}

		if currentHeight == 9 {
			ratingCount++
			return
		}

		directions := []Point{
			{current.x, current.y - 1},
			{current.x, current.y + 1},
			{current.x - 1, current.y},
			{current.x + 1, current.y},
		}

		for _, next := range directions {
			dfs(next, currentHeight)
		}
	}

	dfs(start, -1)

	return ratingCount
}

func main() {
	lines := utils.GetInputFileLines("day10/input.txt")
	heightMap := make([][]int, len(lines))

	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, ch := range line {
			heightMap[i][j] = int(ch - '0')
		}
	}

	fmt.Println("Part 1:", solvePart1(heightMap))
	fmt.Println("Part 2:", solvePart2(heightMap))

	start := time.Now()

	for i := 0; i < 10000; i++ {
		solvePart1(heightMap)
	}

	println("Part 1 Timed:", time.Since(start)/time.Duration(10000), "ns")

	start = time.Now()

	for i := 0; i < 10000; i++ {
		solvePart2(heightMap)
	}

	println("Part 2 Timed:", time.Since(start)/time.Duration(10000), "ns")
}
