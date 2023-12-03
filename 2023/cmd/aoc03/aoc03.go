package aoc03

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc03_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	schematic := newSchematic(lines)
	numbers := schematic.partsAdjacentNumbers()
	sum := 0
	for _, number := range numbers {
		sum += number.value
	}
	return sum
}

func Part2(lines []string) int {
	schematic := newSchematic(lines)
	numbers := schematic.gearsAdjacentNumbers()
	sum := 0
	for _, adjNumbers := range numbers {
		mult := 1
		for _, adjNumber := range adjNumbers {
			mult *= adjNumber.value
		}
		sum += mult
	}
	return sum
}

type schematic struct {
	parts   []part
	numbers []number
}

type number struct {
	start coordinate
	end   coordinate
	value int
}

type part struct {
	value string
	coordinate
}

type coordinate struct {
	row int
	col int
}

func newSchematic(lines []string) *schematic {
	numbers := []number{}
	parts := []part{}
	numberExp := regexp.MustCompile(`\d+`)
	partsExp := regexp.MustCompile(`[^0-9.]+`)
	for row, line := range lines {
		numberCoordMatches := numberExp.FindAllStringIndex(line, -1)
		numberMatches := numberExp.FindAllString(line, -1)
		for idx, match := range numberCoordMatches {
			value, err := strconv.Atoi(numberMatches[idx])
			if err != nil {
				continue
			}
			colStart, colEnd := match[0], match[1]-1
			numbers = append(numbers, number{
				start: coordinate{
					row: row,
					col: colStart,
				},
				end: coordinate{
					row: row,
					col: colEnd,
				},
				value: value,
			})
		}
		partsCoordMatches := partsExp.FindAllStringIndex(line, -1)
		partsMatches := partsExp.FindAllString(line, -1)
		for idx, match := range partsCoordMatches {
			col := match[0]
			parts = append(parts, part{
				value: partsMatches[idx],
				coordinate: coordinate{
					row: row,
					col: col,
				},
			})
		}
	}

	return &schematic{
		parts:   parts,
		numbers: numbers,
	}
}

func (c1 coordinate) isAdjacent(c2 coordinate) bool {
	deltaRow := int(math.Abs(float64(c1.row - c2.row)))
	deltaCol := int(math.Abs(float64(c1.col - c2.col)))

	return deltaRow <= 1 && deltaCol <= 1
}

func (p part) isAdjacentNumber(n number) bool {
	return p.coordinate.isAdjacent(n.start) || p.coordinate.isAdjacent(n.end)
}

func (s *schematic) partsAdjacentNumbers() []number {
	numbers := []number{}
	for _, part := range s.parts {
		for _, number := range s.numbers {
			if part.isAdjacentNumber(number) {
				numbers = append(numbers, number)
			}
		}
	}
	return numbers
}

func (s *schematic) gearsAdjacentNumbers() [][]number {
	numbers := [][]number{}
	for _, part := range s.parts {
		if part.value != "*" {
			continue
		}
		adjacentNumbers := []number{}
		for _, number := range s.numbers {
			if part.isAdjacentNumber(number) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}

		if len(adjacentNumbers) < 2 {
			continue
		}
		numbers = append(numbers, adjacentNumbers)
	}
	return numbers
}
