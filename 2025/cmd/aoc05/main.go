package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/guergeiro/aoc/2025/pkg/aoc05"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc05_input.txt"); err == nil {
		fmt.Println(aoc05.Part1(arr))
		fmt.Println(aoc05.Part2(arr))
	}
}
