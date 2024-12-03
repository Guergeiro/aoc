package aoc03

import (
	"testing"

	reader "github.com/guergeiro/aoc/2024/internal"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)
	assert.Equal(t, 161, Part1(arr))
}

func TestPart2(t *testing.T) {
	arr, err := reader.ReadFileLines("example2.txt")

	assert.NoError(t, err)
	assert.Equal(t, 48, Part2(arr))
}
func TestPart2Edge(t *testing.T) {
	arr, err := reader.ReadFileLines("example3.txt")

	assert.NoError(t, err)
	assert.Equal(t, 8, Part2(arr))
}