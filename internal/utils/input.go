package utils

import (
	"bufio"
	"io"
	"os"
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
