package aoc03

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"deedles.dev/xiter"
)

var mulExp = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
var numberExp = regexp.MustCompile(`\d{1,3}`)
var removeDonts = regexp.MustCompile(`do\(\)(.*?)don\'t\(\)`)

func Part1(arr []string) int {
	line := strings.Join(arr, "")
	return calculate(line)
}

func Part2(arr []string) int {
	inputLine := strings.Join(arr, "")

	line := strings.Join(
		[]string{
			"do()",
			inputLine,
			"don't()",
		},
		"",
	)

	removals := removeDonts.FindAllString(line, -1)
	cleanline := strings.Join(removals, "")
	return calculate(cleanline)
}

func calculate(line string) int {
	valid := mulExp.FindAllString(line, -1)
	digits := xiter.Map(
		slices.Values(valid),
		func(mulStr string) int {
			pair := numberExp.FindAllString(mulStr, -1)
			left, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}
			return left * right
		},
	)
	return xiter.Fold(
		digits,
		func(acc int, curr int) int {
			return acc + curr
		},
	)
}
