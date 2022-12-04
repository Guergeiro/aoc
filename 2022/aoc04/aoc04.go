package aoc04

import (
	"fmt"
	"strconv"
	"strings"

	array "github.com/guergeiro/aoc/2022/_array"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

type assignment struct {
	firstPair  [2]int
	secondPair [2]int
}

func (a assignment) lowerBounds() [2]int {
	return [2]int{
		a.firstPair[0], a.secondPair[0],
	}
}

func (a assignment) upperBounds() [2]int {
	return [2]int{
		a.firstPair[1], a.secondPair[1],
	}
}

func (a assignment) fullyContainsAssignment() bool {
	lowerBounds := a.lowerBounds()
	upperBounds := a.upperBounds()
	if lowerBounds[0] >= lowerBounds[1] && upperBounds[0] <= upperBounds[1] {
		return true
	}
	if lowerBounds[1] >= lowerBounds[0] && upperBounds[1] <= upperBounds[0] {
		return true
	}
	return false
}

func (a assignment) overlapAssignment() bool {
	lowerBounds := a.lowerBounds()
	upperBounds := a.upperBounds()

	if lowerBounds[0] <= upperBounds[1] && lowerBounds[1] <= upperBounds[0] {
		return true
	}

	return false
}

func createPair(splittedPair string) [2]int {
	pair := strings.SplitN(splittedPair, "-", 2)
	convertedPair := [2]int{
		0, 0,
	}
	for idx, strValue := range pair {
		if value, ok := strconv.Atoi(strValue); ok == nil {
			convertedPair[idx] = value
		}
	}
	return convertedPair
}

func createAssignment(row string) assignment {
	splittedPairs := strings.SplitN(row, ",", 2)
	a := assignment{
		firstPair:  createPair(splittedPairs[0]),
		secondPair: createPair(splittedPairs[1]),
	}
	return a
}

func createAssignments(input []string) []assignment {
	assignments := []assignment{}
	for _, row := range input {
		assignments = append(assignments, createAssignment(row))
	}
	return assignments
}

func fullyContainedAmount(assignments []assignment) int {
	return array.Reduce(assignments, func(acc int, cur assignment) int {
		if cur.fullyContainsAssignment() {
			return acc + 1
		}
		return acc
	}, 0)
}

func overlapAmount(assignments []assignment) int {
	return array.Reduce(assignments, func(acc int, cur assignment) int {
		if cur.overlapAssignment() {
			return acc + 1
		}
		return acc
	}, 0)
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc04/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc04/input_question.txt")
	fmt.Println(fullyContainedAmount(createAssignments(arrExample)))
	fmt.Println(fullyContainedAmount(createAssignments(arrQuestion)))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc04/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc04/input_question.txt")
	fmt.Println(overlapAmount(createAssignments(arrExample)))
	fmt.Println(overlapAmount(createAssignments(arrQuestion)))
}
