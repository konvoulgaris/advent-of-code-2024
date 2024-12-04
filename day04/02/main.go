package main

import (
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

func countPattern(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	isMASPattern := func(x, y, dx, dy int) bool {
		if x+2*dx < 0 || x+2*dx >= rows || y+2*dy < 0 || y+2*dy >= cols {
			return false
		}

		if grid[x][y] == 'M' && grid[x+dx][y+dy] == 'A' && grid[x+2*dx][y+2*dy] == 'S' {
			return true
		}

		if grid[x+2*dx][y+2*dy] == 'M' && grid[x+dx][y+dy] == 'A' && grid[x][y] == 'S' {
			return true
		}

		return false
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if (isMASPattern(i-1, j-1, 1, 1) && isMASPattern(i-1, j+1, 1, -1)) ||
				(isMASPattern(i+1, j-1, -1, 1) && isMASPattern(i+1, j+1, -1, -1)) {
				count++
			}
		}
	}

	return count
}

func main() {
	grid := utils.GetInputFileContentAs2DArray("day04/02/input.txt")
	count := countPattern(grid)
	println(count)

	start := time.Now()

	for i := 0; i < 1000; i++ {
		countPattern(grid)
	}

	println(time.Since(start), "ns")
}
