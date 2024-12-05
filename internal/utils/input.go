package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetInputFileContent(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func GetInputFileLines(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func GetInputFileContentAs2DArray(path string) [][]rune {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func GetInputValueCommaListAsIntArray(lines []string) [][]int {
	var input [][]int

	for _, update := range lines {
		split := strings.Split(update, ",")
		var intList []int

		for _, str := range split {
			num, err := strconv.Atoi(strings.TrimSpace(str))

			if err != nil {
				panic(err)
			}

			intList = append(intList, num)
		}
		input = append(input, intList)
	}

	return input
}
