package aoc09

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc09"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 114

	arr, err := reader.ReadFileLines("aoc09_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc09.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 2

	arr, err := reader.ReadFileLines("aoc09_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc09.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
