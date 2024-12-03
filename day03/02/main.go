package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

func main() {
	input := utils.GetInputFileContent("day03/02/input.txt")
	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	mulEnabled := true
	sum := 0

	instructions := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(input, -1)

	for _, instruction := range instructions {
		if mulRe.MatchString(instruction) {
			if mulEnabled {
				inner := strings.TrimPrefix(instruction, "mul(")
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
		} else if doRe.MatchString(instruction) {
			mulEnabled = true
		} else if dontRe.MatchString(instruction) {
			mulEnabled = false
		}
	}

	println(sum)
}
