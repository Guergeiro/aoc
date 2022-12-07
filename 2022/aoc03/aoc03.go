package aoc03

import (
	"fmt"
	"math"

	array "github.com/guergeiro/aoc/2022/_array"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

func Part1() {
	arrExample := reader.ReadFileStrings("aoc03/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc03/input_question.txt")
	fmt.Println(handlePart1(arrExample))
	fmt.Println(handlePart1(arrQuestion))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc03/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc03/input_question.txt")
	fmt.Println(handlePart2(arrExample))
	fmt.Println(handlePart2(arrQuestion))
}

func handlePart1(input []string) int {
	ruckstacks := convertStrToRune(input)
	sum := 0
	for _, ruckstack := range ruckstacks {
		middleIndex := int(math.Floor(float64(len(ruckstack)) / 2))
		firstHalf := ruckstack[:middleIndex]
		secondHalf := ruckstack[middleIndex:]
		repeatedElement := findRepeatedElement(firstHalf, [][]rune{
			secondHalf,
		})
		sum += getValue(repeatedElement)
	}
	return sum
}

func handlePart2(input []string) int {
	ruckstacks := convertStrToRune(input)
	groups := createGroup(ruckstacks, 3)
	sum := 0
	for _, ruckstacks := range groups {
		firstHalf := ruckstacks[0]
		repeatedElement := findRepeatedElement(firstHalf, ruckstacks[1:])
		sum += getValue(repeatedElement)
	}
	return sum
}
func createGroup(runestacks [][]rune, perGroup int) [][][]rune {
	groups := [][][]rune{}

	for idx := 0; idx < len(runestacks); idx += perGroup {
		batch := runestacks[idx:int(math.Min(float64(idx+perGroup), float64(len(runestacks))))]
		groups = append(groups, batch)
	}

	return groups
}

func getValue(str rune) int {
	if str >= 65 && str <= 90 {
		return 27 + int(str) - 65
	}
	if str >= 97 && str <= 122 {
		return 1 + int(str) - 97
	}
	return 0
}

func convertStrToRune(input []string) [][]rune {
	ruckstacks := [][]rune{}
	for _, value := range input {
		strAsRune := []rune(value)
		ruckstacks = append(ruckstacks, strAsRune)
	}
	return ruckstacks
}

func findRepeatedElement(initialRuckstack []rune, ruckstacks [][]rune) rune {
	for _, individualRune := range initialRuckstack {
		if array.ContainsInAll(ruckstacks, individualRune) {
			return individualRune
		}
	}
	return 0
}
