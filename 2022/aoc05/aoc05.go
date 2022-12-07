package aoc05

import (
	"fmt"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2022/_reader"
)

type move struct {
	amount int
	from   int
	to     int
}

func parseMove(row string) (move, error) {
	splittedRows := strings.Split(row, " ")

	outMove := move{}

	if amount, ok := strconv.Atoi(splittedRows[1]); ok != nil {
		return move{}, ok
	} else {
		outMove.amount = amount
	}
	if from, ok := strconv.Atoi(splittedRows[3]); ok != nil {
		return move{}, ok
	} else {
		outMove.from = from
	}
	if to, ok := strconv.Atoi(splittedRows[5]); ok != nil {
		return move{}, ok
	} else {
		outMove.to = to
	}
	return outMove, nil
}

type stack struct {
	position int
	value    []rune
}

func (s stack) getTip() string {
	r := s.value[len(s.value)-1]
	return string(r)
}

func (s *stack) push(values []rune) {
	s.value = append(s.value, values...)
}

func (s *stack) pop(amount int) []rune {
	crates := s.value[len(s.value)-amount:]
	s.value = append(s.value[:len(s.value)-amount])
	return crates
}

func parseStack(row string) (stack, error) {
	outStack := stack{}

	splittedStr := strings.SplitN(row, " ", 2)

	if position, ok := strconv.Atoi(splittedStr[0]); ok != nil {
		return stack{}, ok
	} else {
		outStack.position = position
	}

	outStack.value = []rune(splittedStr[1])

	return outStack, nil
}

func extract(input []string) ([]stack, []move) {
	stacks := []stack{}
	moves := []move{}
	inMovesSection := false
	for _, row := range input {
		if row == "" {
			inMovesSection = true
			continue
		}
		if inMovesSection == false {
			if s, ok := parseStack(row); ok == nil {
				stacks = append(stacks, s)
			}
			continue
		} else {
			if m, ok := parseMove(row); ok == nil {
				moves = append(moves, m)
			}
		}
	}
	return stacks, moves
}

func executeMove(stacks []stack, move move, canMoveStackOfCrates bool) {
	from := &stacks[move.from-1]
	to := &stacks[move.to-1]
	if canMoveStackOfCrates == false {
		for count := 0; count < move.amount; count += 1 {
			crates := from.pop(1)
			to.push(crates)
		}
	} else {
		crates := from.pop(move.amount)
		to.push(crates)
	}
}

func handler(input []string, canMoveStackOfCrates bool) string {
	stacks, moves := extract(input)
	for _, move := range moves {
		executeMove(stacks, move, canMoveStackOfCrates)
	}

	output := []string{}
	for _, stack := range stacks {
		output = append(output, stack.getTip())
	}
	return strings.Join(output, "")
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc05/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc05/input_question.txt")
	fmt.Println(handler(arrExample, false))
	fmt.Println(handler(arrQuestion, false))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc05/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc05/input_question.txt")
	fmt.Println(handler(arrExample, true))
	fmt.Println(handler(arrQuestion, true))
}
