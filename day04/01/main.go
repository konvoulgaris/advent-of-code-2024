package main

import (
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

const word = "XMAS"

var directions = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func countWord(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLength := len(word)
	wordRunes := []rune(word)
	count := 0

	isWordFound := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLength; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != wordRunes[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, d := range directions {
				if isWordFound(i, j, d[0], d[1]) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	grid := utils.GetInputFileContentAs2DArray("day04/01/input.txt")
	count := countWord(grid)
	println(count)

	start := time.Now()

	for i := 0; i < 1000; i++ {
		countWord(grid)
	}

	println(time.Since(start), "ns")
}
