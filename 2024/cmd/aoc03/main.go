package main

import (
	"fmt"

	reader "github.com/guergeiro/aoc/2024/internal"
	"github.com/guergeiro/aoc/2024/pkg/aoc03"
)

func main() {
	if arr, err := reader.ReadFileLines("assets/aoc03_input.txt"); err == nil {
		fmt.Println(aoc03.Part1(arr))
		fmt.Println(aoc03.Part2(arr))
	}
}
