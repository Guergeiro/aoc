package aoc03

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc03"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 4361

	arr, err := reader.ReadFileLines("aoc03_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc03.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 467835

	arr, err := reader.ReadFileLines("aoc03_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc03.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
