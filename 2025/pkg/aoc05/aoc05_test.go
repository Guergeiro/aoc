package aoc05

import (
	"testing"

	"github.com/guergeiro/aoc/2025/internal/reader"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	assert.Equal(t, 3, Part1(arr))
}

func TestPart2(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	assert.Equal(t, 14, Part2(arr))
}
