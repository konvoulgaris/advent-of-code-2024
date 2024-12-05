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

func calculateMiddlePageSum(ruleMap *map[int][]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if isValidUpdate(ruleMap, update) {
			middle := update[len(update)/2]
			sum += middle
		}
	}

	return sum
}

func main() {
	lines := utils.GetInputFileLines("day05/01/input.txt")
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
