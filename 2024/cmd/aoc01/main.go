package main

import (
	"fmt"

	reader "github.com/guergeiro/aoc/2024/internal"
	"github.com/guergeiro/aoc/2024/pkg/aoc01"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc01_input.txt"); err == nil {
		fmt.Println(aoc01.Part1(arr))
		fmt.Println(aoc01.Part2(arr))
	}
}
