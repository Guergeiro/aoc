package aoc06

import (
	"strconv"
)

func Part1(arr []string) int {
	homework := buildHomework(arr, func(numberRows []string, r rangePos) []int {
		numbers := []int{}
		for _, row := range numberRows {
			num := extractNumber(row, r.start, r.end)
			if num >= 0 {
				numbers = append(numbers, num)
			}
		}
		return numbers
	})
	sum := 0
	for _, p := range homework {
		sum += p.Solve()
	}
	return sum
}

func Part2(arr []string) int {
	homework := buildHomework(arr, func(numberRows []string, r rangePos) []int {
		numbers := []int{}
		for col := r.end - 1; col >= r.start; col-- {
			num := extractNumberFromColumn(numberRows, col)
			if num >= 0 {
				numbers = append(numbers, num)
			}
		}
		return numbers
	})
	sum := 0
	for _, p := range homework {
		sum += p.Solve()
	}
	return sum
}

type problem struct {
	numbers []int
	operand string
}

type rangePos struct {
	start int
	end   int
}

func buildHomework(arr []string, extractor func([]string, rangePos) []int) []problem {
	operandRow := arr[len(arr)-1]
	numberRows := arr[:len(arr)-1]

	// Find the maximum length to ensure all rows are considered equally
	maxLen := len(operandRow)
	for _, row := range numberRows {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	// Pad operand row if needed
	for len(operandRow) < maxLen {
		operandRow += " "
	}

	ranges := findOperandRanges(operandRow)
	problems := []problem{}

	for _, r := range ranges {
		operand := string(operandRow[r.start])
		numbers := extractor(numberRows, r)

		problems = append(problems, problem{
			numbers: numbers,
			operand: operand,
		})
	}

	return problems
}

// findOperandRanges returns the ranges for each operand group
func findOperandRanges(operandRow string) []rangePos {
	ranges := []rangePos{}
	start := 0

	for i := 1; i < len(operandRow); i++ {
		if operandRow[i] == '*' || operandRow[i] == '+' {
			// Found next operand, save previous range (excluding separator space)
			ranges = append(ranges, rangePos{start: start, end: i - 1})
			start = i
		}
	}

	// Don't forget the last range
	ranges = append(ranges, rangePos{start: start, end: len(operandRow)})

	return ranges
}

// extractNumber extracts a number from a row within the given range
func extractNumber(row string, start, end int) int {
	numStr := ""
	for i := start; i < end && i < len(row); i++ {
		ch := row[i]
		if ch >= '0' && ch <= '9' {
			numStr += string(ch)
		}
	}

	if numStr == "" {
		return -1
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}
	return num
}

// extractNumberFromColumn extracts a number from a specific column position across all rows
func extractNumberFromColumn(rows []string, col int) int {
	numStr := ""
	for _, row := range rows {
		if col < len(row) {
			ch := row[col]
			if ch >= '0' && ch <= '9' {
				numStr += string(ch)
			}
		}
	}

	if numStr == "" {
		return -1
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}
	return num
}

func (p problem) initialValue() int {
	if p.operand == "*" {
		return 1
	}
	return 0
}

func (p problem) Solve() int {
	total := p.initialValue()

	for _, number := range p.numbers {
		if p.operand == "*" {
			total *= number
		} else {
			total += number
		}
	}

	return total
}
