package aoc04

import (
	"testing"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	assert.Equal(t, 13, Part1(arr))
}

func TestPart2(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	assert.Equal(t, 43, Part2(arr))
}
