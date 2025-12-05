package aoc03

import (
	"slices"
	"strconv"
	"strings"

	"github.com/guergeiro/aoc/2025/internal/stack"
)

func Part1(arr []string) int {
	batteries := parseBatteries(arr)

	joltage := 0

	for _, battery := range batteries {
		joltage += getMax(battery, 2)
	}

	return joltage
}

func Part2(arr []string) int {
	batteries := parseBatteries(arr)

	joltage := 0

	for _, battery := range batteries {
		joltage += getMax(battery, 12)
	}

	return joltage
}

func parseBatteries(rawBatteries []string) [][]string {
	batteries := [][]string{}
	for _, line := range rawBatteries {
		rawBattery := strings.Split(line, "")
		batteries = append(batteries, rawBattery)
	}
	return batteries
}

func getMax(s []string, maxLen int) int {
	st := stack.NewStack([]string{})
	n := len(s)

	for i := range n {
		char := s[i]

		for st.Size() > 0 {
			top, _ := st.Peek()
			if char <= top || st.Size()-1+n-i < maxLen {
				break
			}
			st.Pop()
		}

		if st.Size() < maxLen {
			st.Push(char)
		}
	}

	result := st.PopAll()
	slices.Reverse(result)

	num, err := strconv.Atoi(strings.Join(result, ""))
	if err != nil {
		return 0
	}
	return num
}
