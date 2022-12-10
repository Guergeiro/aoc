package aoc10

import (
	"fmt"
	"strconv"
	"strings"

	array "github.com/guergeiro/aoc/2022/_array"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

func markAsLit(sprite int, cycle int, crt *[6][40]string) {
	col, row := int(cycle/40), cycle%40
	if sprite == row || sprite == row+1 || sprite == row-1 {
		crt[col][row] = "#"
	}
}

func calculateDeterministicCpu(instructions []int, crt *[6][40]string) []int {
	cpu := []int{}

	currentStrength := 1
	for cycle, instruction := range instructions {
		currentStrength += instruction
		cpu = append(cpu, currentStrength)

		markAsLit(currentStrength, cycle, crt)
	}
	return cpu
}

func parseInstructions(input []string) []int {
	instructions := []int{
		0,
	}

	for _, line := range input {
		if line == "noop" {
			instructions = append(instructions, 0)
			continue
		}

		splittedLine := strings.SplitN(line, " ", 2)
		if amount, err := strconv.Atoi(splittedLine[1]); err == nil {

			instructions = append(instructions, []int{
				0,
				amount,
			}...)
		}

	}
	return instructions
}

func inspect(cpu []int, inspectPoints []int) []int {
	output := []int{}
	for _, point := range inspectPoints {
		output = append(output, point*cpu[point-1])
	}

	return output
}

func handler(input []string) int {
	instructions := parseInstructions(input)

	crt := [6][40]string{}

	for y, col := range crt {
		for x := range col {
			crt[y][x] = "."
		}
	}

	cpu := calculateDeterministicCpu(instructions, &crt)
	inspectionResults := inspect(cpu, []int{
		20,
		60,
		100,
		140,
		180,
		220,
	})

	for _, col := range crt {
		fmt.Println(col)
	}

	return array.Reduce(inspectionResults, func(acc, curr int) int {
		return acc + curr
	}, 0)
}

func Part() {
	arrExample := reader.ReadFileStrings("aoc10/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc10/input_question.txt")
	fmt.Println(handler(arrExample))
	fmt.Println(handler(arrQuestion))
}
