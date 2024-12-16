package main

import (
	"container/heap"
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

type Face int

const (
	NORTH Face = iota
	EAST
	SOUTH
	WEST
)

func (f Face) moveCounterclockwise() Face {
	return Face((f - 1 + 4) % 4)
}

func (f Face) moveClockwise() Face {
	return Face((f + 1) % 4)
}

type State struct {
	row, col int
	facing   Face
	score    int
	path     [][]int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var directions = map[Face][2]int{
	NORTH: {-1, 0},
	EAST:  {0, 1},
	SOUTH: {1, 0},
	WEST:  {0, -1},
}

func findPoint(grid []string, point byte) (int, int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == point {
				return row, col
			}
		}
	}
	panic("Point not found")
}

func moveForward(curr *State, grid []string, pq *PriorityQueue) {
	dr, dc := directions[curr.facing][0], directions[curr.facing][1]
	newRow, newCol := curr.row+dr, curr.col+dc
	if newRow >= 0 && newRow < len(grid) &&
		newCol >= 0 && newCol < len(grid[0]) &&
		grid[newRow][newCol] != '#' {
		newPath := make([][]int, len(curr.path))
		copy(newPath, curr.path)
		heap.Push(pq, &State{
			row:    newRow,
			col:    newCol,
			facing: curr.facing,
			score:  curr.score + 1,
			path:   append(newPath, []int{curr.row, curr.col}),
		})
	}
}

func rotateClockwise(curr *State, pq *PriorityQueue) {
	clockwiseFacing := curr.facing.moveClockwise()
	heap.Push(pq, &State{
		row:    curr.row,
		col:    curr.col,
		facing: clockwiseFacing,
		score:  curr.score + 1000,
		path:   append(curr.path, []int{curr.row, curr.col}),
	})
}

func rotateCounterclockwise(curr *State, pq *PriorityQueue) {
	counterclockwiseFacing := curr.facing.moveCounterclockwise()
	heap.Push(pq, &State{
		row:    curr.row,
		col:    curr.col,
		facing: counterclockwiseFacing,
		score:  curr.score + 1000,
		path:   append(curr.path, []int{curr.row, curr.col}),
	})
}

func solvePartOne(grid []string) int {
	startRow, startCol := findPoint(grid, 'S')
	endRow, endCol := findPoint(grid, 'E')

	pq := &PriorityQueue{}
	heap.Init(pq)

	bestScores := make(map[[3]int]int)
	initialPath := [][]int{}
	heap.Push(pq, &State{
		row:    startRow,
		col:    startCol,
		facing: EAST,
		score:  0,
		path:   initialPath,
	})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*State)
		key := [3]int{curr.row, curr.col, int(curr.facing)}

		if prevScore, exists := bestScores[key]; exists && prevScore <= curr.score {
			continue
		}

		bestScores[key] = curr.score

		if curr.row == endRow && curr.col == endCol {
			return curr.score
		}

		moveForward(curr, grid, pq)
		rotateClockwise(curr, pq)
		rotateCounterclockwise(curr, pq)
	}

	return -1
}

func main() {
	maze := utils.GetInputFileLines("day16/input.txt")
	println("Part 1:", solvePartOne(maze))

	start := time.Now()

	for i := 0; i < 1000; i++ {
		println(i)
		solvePartOne(maze)
	}

	println("Part 1 Timed:", time.Since(start)/time.Duration(1000), "ns")
}
