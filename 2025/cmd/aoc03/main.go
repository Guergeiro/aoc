package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/guergeiro/aoc/2025/pkg/aoc03"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc03_input.txt"); err == nil {
		fmt.Println(aoc03.Part1(arr))
		fmt.Println(aoc03.Part2(arr))
	}
}
