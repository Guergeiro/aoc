package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/guergeiro/aoc/2025/pkg/aoc01"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc01_input.txt"); err == nil {
		fmt.Println(aoc01.Part1(arr))
		fmt.Println(aoc01.Part2(arr))
	}
}
