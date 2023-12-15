package aoc07

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc07"
	"github.com/guergeiro/aoc/2023/internal"
)

func TestPart1Example(t *testing.T) {
	expected := 6440

	arr, err := reader.ReadFileLines("aoc07_example.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc07.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart1Extra1(t *testing.T) {
	expected := 1343

	arr, err := reader.ReadFileLines("aoc07_extra1.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc07.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart1Extra2(t *testing.T) {
	expected := 2237

	arr, err := reader.ReadFileLines("aoc07_extra2.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc07.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

func TestPart1Extra3(t *testing.T) {
	expected := 6592

	arr, err := reader.ReadFileLines("aoc07_extra3.txt")
	if err != nil {
		t.Fail()
	}
	actual := aoc07.Part1(arr)

	if expected != actual {
		t.Errorf("expectected %d, actual %d", expected, actual)
	}
}

// func TestPart2Example(t *testing.T) {
// 	expected := 5905

// 	arr, err := reader.ReadFileLines("aoc07_example.txt")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	actual := aoc07.Part2(arr)

// 	if expected != actual {
// 		t.Errorf("expectected %d, actual %d", expected, actual)
// 	}
// }

// func TestPart2Extra2(t *testing.T) {
// 	expected := 2297

// 	arr, err := reader.ReadFileLines("aoc07_extra2.txt")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	actual := aoc07.Part1(arr)

// 	if expected != actual {
// 		t.Errorf("expectected %d, actual %d", expected, actual)
// 	}
// }

// func TestPart2Extra3(t *testing.T) {
// 	expected := 6839

// 	arr, err := reader.ReadFileLines("aoc07_extra2.txt")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	actual := aoc07.Part1(arr)

// 	if expected != actual {
// 		t.Errorf("expectected %d, actual %d", expected, actual)
// 	}
// }
