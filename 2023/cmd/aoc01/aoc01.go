package aoc01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc01_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	numbers := []int{}
	for _, line := range lines {
		numbers = append(numbers, getNumbers(line, false))
	}
	return sum(numbers)
}

func Part2(lines []string) int {
	numbers := []int{}
	for _, line := range lines {
		numbers = append(numbers, getNumbers(line, true))
	}
	return sum(numbers)
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func getNumbers(line string, expandAlpha bool) int {
	digits := []rune{}
	strArr := []rune(line)
	for idx, c := range strArr {
		if unicode.IsNumber(c) {
			digits = append(digits, c)
		}
		if expandAlpha {
			digits = append(digits, getAlpha(line[idx:])...)
		}
	}
	numberStr := fmt.Sprintf(
		"%s%s",
		string(digits[0]),
		string(digits[len(digits)-1]),
	)
	if number, err := strconv.Atoi(numberStr); err == nil {
		return number
	}
	return 0
}

func getAlpha(curline string) []rune {
	needles := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	digits := []rune{}
	for d, key := range needles {
		if found := strings.HasPrefix(curline, key); found {
			digits = append(digits, rune(d+1+48))
		}
	}
	return digits
}
