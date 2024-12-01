package aoc01

import (
	"math"
	"regexp"
	"slices"
	"strconv"
)

var numberExp = regexp.MustCompile(`\d+`)

func Part1(arr []string) int {
	left, right, err := newLists(arr)
	if err != nil {
		return 0
	}

	sortedLeft := slices.Sorted(slices.Values(left))
	sortedRight := slices.Sorted(slices.Values(right))

	listDifferences := getListsDifferences(sortedLeft, sortedRight)
	sum := reduceList(listDifferences)
	return sum
}

func Part2(arr []string) int {
	left, right, err := newLists(arr)
	if err != nil {
		return 0
	}

	counted := convertToRepeatedElements(right)

	similarities := getSimilarities(left, counted)

	sum := reduceList(similarities)
	return sum
}

func newLists(arr []string) ([]int, []int, error) {
	left := []int{}
	right := []int{}
	for _, line := range arr {
		digits := numberExp.FindAllString(line, -1)

		leftDigit, err := strconv.Atoi(digits[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		rightDigit, err := strconv.Atoi(digits[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		left = append(left, leftDigit)
		right = append(right, rightDigit)
	}
	return left, right, nil
}

func getListsDifferences(l1 []int, l2 []int) []int {
	out := []int{}

	for i := range l1 {
		diff := math.Abs(float64(l1[i] - l2[i]))
		out = append(out, int(diff))
	}

	return out
}

func convertToRepeatedElements(l []int) map[int]int {
	m := map[int]int{}

	for _, element := range l {
		count := m[element]
		m[element] = count + 1
	}

	return m
}

func getSimilarities(l []int, m map[int]int) []int {
	out := []int{}

	for _, element := range l {
		value := element * m[element]
		out = append(out, value)
	}
	return out
}

func reduceList(l []int) int {
	out := 0
	for _, v := range l {
		out += v
	}
	return out
}
