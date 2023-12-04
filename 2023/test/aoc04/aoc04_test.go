package aoc04

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc04"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 13

	arr, err := reader.ReadFileLines("aoc04_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc04.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 30

	arr, err := reader.ReadFileLines("aoc04_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc04.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
