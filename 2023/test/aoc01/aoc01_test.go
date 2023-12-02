package aoc01

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc01"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 142

	arr, err := reader.ReadFileLines("aoc01_part1_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc01.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 281

	arr, err := reader.ReadFileLines("aoc01_part2_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc01.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
