package main

import (
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

var directions = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func findGuard(grid []string) (int, int) {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (grid)[i][j] == '^' {
				return i, j
			}
		}
	}

	panic("Guard not found!")
}

func turnRight(direction *int) {
	*direction = (*direction + 1) % 4
}

func moveForwardAndRemember(x, y *int, nx, ny int, visited *map[[2]int]bool) {
	*x, *y = nx, ny
	(*visited)[[2]int{*x, *y}] = true
}

func solvePart1(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])

	direction := 0 // Starting direction is Up

	x, y := findGuard(grid)

	visited := map[[2]int]bool{
		{x, y}: true,
	}

	for {
		nx, ny := x+directions[direction][0], y+directions[direction][1]

		if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
			break
		}

		if grid[nx][ny] == '#' {
			turnRight(&direction)
		} else {
			moveForwardAndRemember(&x, &y, nx, ny, &visited)
		}

		if len(visited) > rows*cols {
			break
		}
	}

	return len(visited)
}

func main() {
	grid := utils.GetInputFileLines("day06/01/input.txt")
	println("Part 1", solvePart1(grid))

	start := time.Now()

	for i := 0; i < 10000; i++ {
		solvePart1(grid)
	}

	println(time.Since(start)/time.Duration(10000), "ns")
}
