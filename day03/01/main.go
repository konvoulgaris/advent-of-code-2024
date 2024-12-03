package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

func main() {
	input := utils.GetInputFileContent("day03/01/input.txt")
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)

	sum := 0

	for _, match := range matches {
		inner := strings.TrimPrefix(match, "mul(")
		inner = strings.TrimSuffix(inner, ")")
		parts := strings.Split(inner, ",")

		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])

			if err1 == nil && err2 == nil {
				sum += num1 * num2
			}
		}
	}

	println(sum)
}
