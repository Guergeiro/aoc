package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/guergeiro/aoc/2025/pkg/aoc02"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc02_input.txt"); err == nil {
		fmt.Println(aoc02.Part1(arr))
		fmt.Println(aoc02.Part2(arr))
	}
}
