package aoc05

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc05"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	expected := 35

	arr, err := reader.ReadFileLines("aoc05_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc05.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 46

// 	arr, err := reader.ReadFileLines("aoc05_example.txt")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	actual := aoc05.Part2(arr)

// 	if expected != actual {
// 		t.Errorf("expectected %d, actual %d", expected, actual)
// 	}
// }
