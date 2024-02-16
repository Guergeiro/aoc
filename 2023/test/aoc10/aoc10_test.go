package aoc10

import (
	"testing"

	"github.com/guergeiro/aoc/2023/cmd/aoc10"
	"github.com/guergeiro/aoc/2023/internal"
)

type input struct {
	file     string
	expected int
}

func TestPart1(t *testing.T) {
	examples := []input{
		// {
		// 	file:     "aoc10_example_1.txt",
		// 	expected: 4,
		// },
		// {
		// 	file:     "aoc10_example_2.txt",
		// 	expected: 8,
		// },
		// {
		// 	file:     "aoc10_example_3.txt",
		// 	expected: 4,
		// },
		{
			file:     "aoc10_example_4.txt",
			expected: 8,
		},
	}

	for _, example := range examples {
		arr, err := reader.ReadFileLines(example.file)
		if err != nil {
			t.Fail()
		}
		actual := aoc10.Part1(arr)

		if example.expected != actual {
			t.Errorf("expectected %d, actual %d", example.expected, actual)
		}
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 2

// 	arr, err := reader.ReadFileLines("aoc10_example.txt")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	actual := aoc10.Part2(arr)

// 	if expected != actual {
// 		t.Errorf("expectected %d, actual %d", expected, actual)
// 	}
// }
