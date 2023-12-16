package aoc08

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc08"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 2

	arr, err := reader.ReadFileLines("aoc08_part1_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc08.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 6

	arr, err := reader.ReadFileLines("aoc08_part2_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc08.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
