package aoc02

import (
	"strconv"
	"strings"
)

func Part1(arr []string) int {
	intervals := parseIntervals(arr[0])
	sum := 0
	for _, interval := range intervals {
		sum += getInvalidIdsSum(interval[0], interval[1], func(str string)bool {
			return len(str)%2 == 0 && isInvalidId(str, len(str)/2, len(str)/2)
		})
	}
	return sum
}

func Part2(arr []string) int {
	intervals := parseIntervals(arr[0])
	sum := 0
	for _, interval := range intervals {
		sum += getInvalidIdsSum(interval[0], interval[1], func(str string)bool {
			return isInvalidId(str, 1, len(str)/2)
		})
	}
	return sum
}

func parseIntervals(line string) [][]string {
	rawIntervals := strings.Split(line, ",")

	intervals := [][]string{}
	for _, rawInterval := range rawIntervals {
		interval := strings.Split(rawInterval, "-")
		intervals = append(intervals, interval)
	}
	return intervals
}

func getInvalidIdsSum(left string, right string, filter func(string)bool) int {
	lower, err := strconv.Atoi(left)
	if err != nil {
		return 0
	}
	higher, err := strconv.Atoi(right)
	if err != nil {
		return 0
	}
	sum := 0
	for i := lower; i <= higher; i += 1 {
		str := strconv.Itoa(i)
		if filter(str) {
			sum += i
		}
	}
	return sum
}

func isInvalidId(str string, start, end int) bool {
	for i := start; i <= end; i++ {
		suffix := str[:i]
		repeated := strings.Repeat(suffix, len(str)/i)
		if repeated == str {
			return true
		}
	}
	return false
}
