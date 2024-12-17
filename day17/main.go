package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/konvoulgaris/advent-of-code-2024/internal/utils"
)

func parseValue(line string) int {
	parts := strings.Split(line, ":")
	value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic("Invalid register value")
	}
	return value
}

func parseProgram(line string) []int {
	parts := strings.Split(line, ":")
	programStr := strings.TrimSpace(parts[1])
	programParts := strings.Split(programStr, ",")
	program := make([]int, len(programParts))
	for i, part := range programParts {
		value, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			panic("Invalid program value")
		}
		program[i] = value
	}
	return program
}

func getLiteral(operand int) int {
	return operand
}

func getCombo(operand int, registerA, registerB, registerC int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	default:
		panic("Invalid combo operand")
	}
}

func solvePart1(input string) string {
	lines := strings.Split(input, "\n")
	var registerA, registerB, registerC int
	var program []int

	for _, line := range lines {
		if strings.HasPrefix(line, "Register A:") {
			registerA = parseValue(line)
		} else if strings.HasPrefix(line, "Register B:") {
			registerB = parseValue(line)
		} else if strings.HasPrefix(line, "Register C:") {
			registerC = parseValue(line)
		} else if strings.HasPrefix(line, "Program:") {
			program = parseProgram(line)
		}
	}

	instructionPointer := 0
	var output []string

	for instructionPointer < len(program) {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]

		switch opcode {
		case 0:
			denominator := 1 << getCombo(operand, registerA, registerB, registerC)
			if denominator != 0 {
				registerA /= denominator
			}
			instructionPointer += 2

		case 1:
			registerB ^= getLiteral(operand)
			instructionPointer += 2

		case 2:
			registerB = getCombo(operand, registerA, registerB, registerC) % 8
			instructionPointer += 2

		case 3:
			if registerA != 0 {
				instructionPointer = getLiteral(operand)
			} else {
				instructionPointer += 2
			}

		case 4:
			registerB ^= registerC
			instructionPointer += 2

		case 5:
			outputValue := getCombo(operand, registerA, registerB, registerC) % 8
			output = append(output, fmt.Sprintf("%d", outputValue))
			instructionPointer += 2

		case 6:
			denominator := 1 << getCombo(operand, registerA, registerB, registerC)
			if denominator != 0 {
				registerB = registerA / denominator
			}
			instructionPointer += 2

		case 7:
			denominator := 1 << getCombo(operand, registerA, registerB, registerC)
			if denominator != 0 {
				registerC = registerA / denominator
			}
			instructionPointer += 2

		default:
			panic("Unknown opcode")
		}
	}

	return strings.Join(output, ",")
}

func solvePart2(data string) int {
	parts := strings.Split(data, "\n\n")
	registersOriginal := make(map[string]int)
	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) == 0 {
			continue
		}
		part := strings.Split(line, ": ")
		value, _ := strconv.Atoi(part[1])
		key := strings.Fields(part[0])[1]
		registersOriginal[key] = value
	}

	program := make([]int, 0)
	parts[1] = strings.Split(parts[1], ": ")[1]
	for _, numStr := range strings.Split(strings.ReplaceAll(parts[1], "\n", ""), ",") {
		num, _ := strconv.Atoi(numStr)
		program = append(program, num)
	}

	var instructions = []func(int, map[string]int, int, *[]int) int{
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			registers["A"] = registers["A"] / (1 << combo(num, registers))
			return instr + 2
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			registers["B"] ^= num
			return instr + 2
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			registers["B"] = combo(num, registers) % 8
			return instr + 2
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			if registers["A"] == 0 {
				return instr + 2
			}
			return num
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			registers["B"] ^= registers["C"]
			return instr + 2
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			*outputs = append(*outputs, combo(num, registers)%8)
			return instr + 2
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			registers["B"] = registers["A"] / (1 << combo(num, registers))
			return instr + 2
		},
		func(num int, registers map[string]int, instr int, outputs *[]int) int {
			registers["C"] = registers["A"] / (1 << combo(num, registers))
			return instr + 2
		},
	}

	getOutput := func(a int) []int {
		outputs := []int{}
		registers := make(map[string]int)
		for k, v := range registersOriginal {
			registers[k] = v
		}
		registers["A"] = a
		instr := 0

		for instr >= 0 && instr < len(program) {
			opcode := program[instr]
			num := program[instr+1]
			instr = instructions[opcode](num, registers, instr, &outputs)
		}
		return outputs
	}

	valid := []int{0}
	for length := 1; length <= len(program); length++ {
		oldValid := valid
		valid = []int{}

		for _, num := range oldValid {
			for offset := 0; offset < 8; offset++ {
				newNum := 8*num + offset
				if equal(getOutput(newNum), program[len(program)-length:]) {
					valid = append(valid, newNum)
				}
			}
		}
	}

	answer := math.MaxInt
	for _, v := range valid {
		if v < answer {
			answer = v
		}
	}

	return answer
}

func combo(num int, registers map[string]int) int {
	switch num {
	case 0, 1, 2, 3:
		return num
	case 4:
		return registers["A"]
	case 5:
		return registers["B"]
	case 6:
		return registers["C"]
	default:
		panic("Invalid combo operand")
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	data := utils.GetInputFileContent("day17/input.txt")
	println("Part 1:", solvePart1(data))
	println("Part 2:", solvePart2(data))

	start := time.Now()

	for i := 0; i < 10000; i++ {
		solvePart1(data)
	}

	println("Part 1 Timed:", time.Since(start)/time.Duration(10000), "ns")

	start = time.Now()

	for i := 0; i < 10000; i++ {
		solvePart2(data)
	}

	println("Part 2 Timed:", time.Since(start)/time.Duration(10000), "ns")
}
