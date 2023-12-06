package aoc06

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc06"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 288

	arr, err := reader.ReadFileLines("aoc06_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc06.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 71503

	arr, err := reader.ReadFileLines("aoc06_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc06.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
