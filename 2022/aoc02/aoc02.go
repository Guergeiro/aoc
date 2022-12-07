package aoc02

import (
	"fmt"
	array "github.com/guergeiro/aoc/2022/_array"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

func Part1() {
	arrExample := reader.ReadFileStrings("aoc02/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc02/input_question.txt")

	fmt.Println(getSum(convertToCleanGames(arrExample)))
	fmt.Println(getSum(convertToCleanGames(arrQuestion)))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc02/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc02/input_question.txt")

	fmt.Println(getSum(extractGames(arrExample)))
	fmt.Println(getSum(extractGames(arrQuestion)))
}

func getSum(games []string) int {
	possiblePlays := getPossiblePlays()
	sum := array.Reduce(games, func(acc int, cur string) int {
		value := possiblePlays[cur]
		return acc + value
	}, 0)
	return sum
}

func convertToCleanGames(input []string) []string {
	games := []string{}
	for _, value := range input {
		games = append(games, value)
	}
	return games
}

func extractGames(input []string) []string {
	games := []string{}
	possibleMatchingPlays := getMatchingPlays()
	for _, value := range input {
		games = append(games, possibleMatchingPlays[value])
	}
	return games
}

func getPossiblePlays() map[string]int {
	possiblePlays := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	return possiblePlays
}

func getMatchingPlays() map[string]string {
	matchingPlays := map[string]string{
		"A X": "A Z",
		"A Y": "A X",
		"A Z": "A Y",
		"B X": "B X",
		"B Y": "B Y",
		"B Z": "B Z",
		"C X": "C Y",
		"C Y": "C Z",
		"C Z": "C X",
	}
	return matchingPlays
}
