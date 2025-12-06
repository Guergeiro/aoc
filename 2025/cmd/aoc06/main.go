package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/guergeiro/aoc/2025/pkg/aoc06"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc06_input.txt"); err == nil {
		fmt.Println(aoc06.Part1(arr))
		fmt.Println(aoc06.Part2(arr))
	}
}
