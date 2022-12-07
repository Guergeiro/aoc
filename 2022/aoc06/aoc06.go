package aoc06

import (
	"fmt"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

func hasDup(s []rune) bool {
	unique := map[string]int{}
	for _, v := range s {
		if _, ok := unique[string(v)]; ok == false {
			unique[string(v)] = 1
		} else {
			return true
		}
	}
	return false
}

func getDifferentIndex(input []string, amount int) []int {
	indexes := []int{}
	for _, row := range input {
		runeSlice := []rune(row)
		for idx := 0; idx < len(runeSlice); idx += 1 {
			slice := runeSlice[idx : idx+amount]
			if hasDup(slice) == false {
				indexes = append(indexes, idx+amount)
				break
			}
		}
	}
	return indexes
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc06/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc06/input_question.txt")
	fmt.Println(getDifferentIndex(arrExample, 4))
	fmt.Println(getDifferentIndex(arrQuestion, 4))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc06/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc06/input_question.txt")
	fmt.Println(getDifferentIndex(arrExample, 14))
	fmt.Println(getDifferentIndex(arrQuestion, 14))
}
