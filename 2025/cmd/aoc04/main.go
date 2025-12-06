package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/guergeiro/aoc/2025/pkg/aoc04"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc04_input.txt"); err == nil {
		fmt.Println(aoc04.Part1(arr))
		fmt.Println(aoc04.Part2(arr))
	}
}
