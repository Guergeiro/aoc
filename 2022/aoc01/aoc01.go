package aoc01

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/guergeiro/aoc/2022/_array"
	"github.com/guergeiro/aoc/2022/_reader"
)

func Part1() {
	arrExample := reader.ReadFileStrings("aoc01/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc01/input_question.txt")

	topNCaloriesExample := getTopNCalories(1, getPerElfCalories(arrExample))
	topNCaloriesQuestion := getTopNCalories(1, getPerElfCalories(arrQuestion))

	fmt.Println(array.Reduce(topNCaloriesExample, func(acc int, cur int) int {
		return acc + cur
	}, 0))
	fmt.Println(array.Reduce(topNCaloriesQuestion, func(acc int, cur int) int {
		return acc + cur
	}, 0))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc01/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc01/input_question.txt")

	topNCaloriesExample := getTopNCalories(3, getPerElfCalories(arrExample))
	topNCaloriesQuestion := getTopNCalories(3, getPerElfCalories(arrQuestion))

	fmt.Println(array.Reduce(topNCaloriesExample, func(acc int, cur int) int {
		return acc + cur
	}, 0))
	fmt.Println(array.Reduce(topNCaloriesQuestion, func(acc int, cur int) int {
		return acc + cur
	}, 0))
}

func getPerElfCalories(arrInput []string) []int {
	perElfCalories := []int{0}
	idx := 0
	for _, strValue := range arrInput {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			idx += 1
			perElfCalories = append(perElfCalories, 0)
			continue
		}
		perElfCalories[idx] += value
	}
	return perElfCalories
}

func getTopNCalories(amount int, perElfCalories []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(perElfCalories)))

	return perElfCalories[:amount]
}
