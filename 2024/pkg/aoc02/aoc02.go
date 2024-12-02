package aoc02

import (
	"regexp"
	"slices"
	"strconv"
)

var numberExp = regexp.MustCompile(`\d+`)

func Part1(arr []string) int {
	reports := newReports(arr)

	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe += 1
		}
	}

	return safe
}

func Part2(arr []string) int {
	reports := newReports(arr)

	safe := 0

	for _, report := range reports {
		if isSafeToleration(report) {
			safe += 1
		}
	}

	return safe
}

func isSafe(levels []int) bool {
	if findIndexFailsAscending(levels) == -1 {
		return true
	}
	if findIndexFailsDescending(levels) == -1 {
		return true
	}
	return false
}

func isSafeToleration(levels []int) bool {
	if index := findIndexFailsAscending(levels); index != -1 {
		withoutJ := slices.Delete(append(levels[:0:0], levels...), index, index+1)
		if findIndexFailsAscending(withoutJ) == -1 {
			return true
		}
		withoutI := slices.Delete(append(levels[:0:0], levels...), index-1, index)
		if findIndexFailsAscending(withoutI) == -1 {
			return true
		}
	} else {
		return true
	}
	if index := findIndexFailsDescending(levels); index != -1 {
		withoutJ := slices.Delete(append(levels[:0:0], levels...), index, index+1)
		if findIndexFailsDescending(withoutJ) == -1 {
			return true
		}
		withoutI := slices.Delete(append(levels[:0:0], levels...), index-1, index)
		if findIndexFailsDescending(withoutI) == -1 {
			return true
		}
	} else {
		return true
	}
	return false
}

func findIndexFailsAscending(levels []int) int {
	i := 0
	j := 1
	for j < len(levels) {
		diff := levels[j] - levels[i]
		if diff > 3 {
			return j
		}
		if diff < 1 {
			return j
		}
		i += 1
		j += 1
	}
	return -1
}

func findIndexFailsDescending(levels []int) int {
	i := 0
	j := 1
	for j < len(levels) {
		diff := levels[i] - levels[j]
		if diff > 3 {
			return j
		}
		if diff < 1 {
			return j
		}
		i += 1
		j += 1
	}
	return -1
}

func newReports(lines []string) [][]int {
	reports := [][]int{}
	for _, line := range lines {
		reports = append(reports, newReport(line))
	}
	return reports
}

func newReport(line string) []int {
	digits := numberExp.FindAllString(line, -1)

	levels := []int{}

	for _, digit := range digits {

		value, err := strconv.Atoi(digit)
		if err != nil {
			continue
		}
		levels = append(levels, value)
	}

	return levels
}
