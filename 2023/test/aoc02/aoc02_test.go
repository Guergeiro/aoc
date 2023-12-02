package aoc02

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc02"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 8

	arr, err := reader.ReadFileLines("aoc02_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc02.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 2286

	arr, err := reader.ReadFileLines("aoc02_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc02.Part2(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}
