package aoc09

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc09_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	sum := 0
	for _, h := range createHistories(lines) {
		sum += h.calcNextValue()
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, h := range createHistories(lines) {
		sum += h.reverse().calcNextValue()
	}
	return sum
}

type history []int

func (h history) allZeros() bool {
	for _, value := range h {
		if value != 0 {
			return false
		}
	}
	return true
}

func (h history) reverse() history {
	slices.Reverse(h)
	return h
}

func (values history) calcNextValue() int {
	diffs := history{}

	i := 1
	for i < len(values) {
		diffs = append(diffs, values[i]-values[i-1])
		i += 1
	}

	if diffs.allZeros() {
		return values[len(values)-1]
	}
	return values[len(values)-1] + diffs.calcNextValue()
}

func createHistories(lines []string) []history {
	histories := []history{}
	for _, line := range lines {
		histories = append(histories, createHistory(line))
	}
	return histories
}

func createHistory(line string) history {
	splitted := strings.Split(line, " ")
	value := history{}
	for _, split := range splitted {
		if number, err := strconv.Atoi(split); err == nil {
			value = append(value, number)
		}
	}
	return value
}
