package aoc01

import (
	"container/ring"
	"strconv"
)

func Part1(arr []string) int {
	r := newRing()

	sum := 0

	for _, line := range arr {
		direction, amount := getMoveAmount(line)
		r = r.Move(direction * amount)
		if r.Value == 0 {
			sum += 1
		}
	}

	return sum
}

func Part2(arr []string) int {
	r := newRing()

	sum := 0

	for _, line := range arr {
		direction, amount := getMoveAmount(line)
		for i := 0; i < amount; i += 1 {
			r = r.Move(direction)
			if r.Value == 0 {
				sum += 1
			}
		}
	}

	return sum
}

func newRing() *ring.Ring {
	r := ring.New(100)
	// Initialize ring
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	// Make it start at 50
	return r.Move(r.Len() / 2)
}

func getMoveAmount(line string) (int, int) {
	prefix := line[0:1]
	suffix, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, 0
	}
	if prefix == "L" {
		return -1, suffix
	} else {
		return 1, suffix
	}
}
