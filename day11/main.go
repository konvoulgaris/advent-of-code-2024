package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

type StoneCounter map[string]int

func removeLeadingZeros(s string) string {
	if s == "0" {
		return s
	}

	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}

	if s == "" {
		return "0"
	}

	return s
}

func transformStones(stoneCounts StoneCounter) StoneCounter {
	newStoneCounts := make(StoneCounter)

	for stoneValue, count := range stoneCounts {
		if stoneValue == "0" {
			newStoneCounts["1"] += count
			continue
		}

		if len(stoneValue)%2 == 0 {
			midpoint := len(stoneValue) / 2
			leftStone := removeLeadingZeros(stoneValue[:midpoint])
			rightStone := removeLeadingZeros(stoneValue[midpoint:])

			newStoneCounts[leftStone] += count
			newStoneCounts[rightStone] += count
			continue
		}

		num, _ := strconv.Atoi(stoneValue)
		newValue := strconv.Itoa(num * 2024)
		newStoneCounts[newValue] += count
	}

	return newStoneCounts
}

func solve(times int, stoneCounts StoneCounter) int {
	for blink := 0; blink < times; blink++ {
		stoneCounts = transformStones(stoneCounts)
	}

	totalStones := 0

	for _, count := range stoneCounts {
		totalStones += count
	}

	return totalStones
}

func main() {
	input := utils.GetInputFileContent("day11/input.txt")
	numStrings := strings.Fields(input)
	stoneCounts := make(StoneCounter)

	for _, numStr := range numStrings {
		stoneCounts[numStr]++
	}

	println("Part 1 (25 Blinks):", solve(25, stoneCounts))
	println("Part 1 (75 Blinks):", solve(75, stoneCounts))

	start := time.Now()

	for i := 0; i < 10000; i++ {
		solve(25, stoneCounts)
	}

	println("Part 1 Timed:", time.Since(start)/time.Duration(10000), "ns")

	start = time.Now()

	for i := 0; i < 10000; i++ {
		solve(75, stoneCounts)
	}

	println("Part 2 Timed:", time.Since(start)/time.Duration(10000), "ns")
}
