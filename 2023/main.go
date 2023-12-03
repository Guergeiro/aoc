package main

import (
	"fmt"

	"github.com/guergeiro/aoc/2023/cmd/aoc01"
	"github.com/guergeiro/aoc/2023/cmd/aoc02"
	"github.com/guergeiro/aoc/2023/cmd/aoc03"
)

func main() {
	fmt.Println("Day 01")
	aoc01.Solve()
	fmt.Println("")

	fmt.Println("Day 02")
	aoc02.Solve()
	fmt.Println("")

	fmt.Println("Day 03")
	aoc03.Solve()
	fmt.Println("")
}
