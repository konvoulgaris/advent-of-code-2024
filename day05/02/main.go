package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

func getRulesAndUpdates(lines []string) ([]string, []string) {
	var rules []string
	var updates []string
	isUpdates := false

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			isUpdates = true
			continue
		}
		if isUpdates {
			updates = append(updates, line)
		} else {
			rules = append(rules, line)
		}
	}

	return rules, updates
}

func getRuleMap(rules []string) *map[int][]int {
	ruleMap := make(map[int][]int)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		ruleMap[x] = append(ruleMap[x], y)
	}

	return &ruleMap
}

func isValidUpdate(ruleMap *map[int][]int, update []int) bool {
	indexMap := make(map[int]int)

	for i, page := range update {
		indexMap[page] = i
	}

	for x, dependents := range *ruleMap {
		if iX, existsX := indexMap[x]; existsX {
			for _, y := range dependents {
				if iY, existsY := indexMap[y]; existsY && iX >= iY {
					return false
				}
			}
		}
	}

	return true
}

func fixUpdate(ruleMap *map[int][]int, update []int) []int {
	inDegree := make(map[int]int)
	graph := make(map[int][]int)

	pageSet := make(map[int]bool)
	for _, page := range update {
		pageSet[page] = true
	}

	for x, dependents := range *ruleMap {
		if !pageSet[x] {
			continue
		}
		for _, y := range dependents {
			if !pageSet[y] {
				continue
			}
			graph[x] = append(graph[x], y)
			inDegree[y]++
		}
	}

	queue := []int{}

	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	sortedUpdate := []int{}
	visited := make(map[int]bool)

	// Topological sort
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		sortedUpdate = append(sortedUpdate, current)
		visited[current] = true

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If topological sort is incomplete, add remaining pages
	for _, page := range update {
		if !visited[page] {
			sortedUpdate = append(sortedUpdate, page)
		}
	}

	return sortedUpdate
}

func calculateMiddlePageSum(ruleMap *map[int][]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if !isValidUpdate(ruleMap, update) {
			correctedUpdate := fixUpdate(ruleMap, update)
			middle := correctedUpdate[len(correctedUpdate)/2]
			sum += middle
		}
	}

	return sum
}

func main() {
	lines := utils.GetInputFileLines("day05/02/input.txt")
	rules, updates := getRulesAndUpdates(lines)
	ruleMap := getRuleMap(rules)
	updatesList := utils.GetInputValueCommaListAsIntArray(updates)
	sum := calculateMiddlePageSum(ruleMap, updatesList)

	println(sum)

	start := time.Now()

	for i := 0; i < 10000; i++ {
		calculateMiddlePageSum(ruleMap, updatesList)
	}

	println(time.Since(start)/time.Duration(10000), "ns")
}
