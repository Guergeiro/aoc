package aoc02

import (
	"testing"

	reader "github.com/guergeiro/aoc/2025/internal"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	assert.Equal(t, 1227775554, Part1(arr))
}

func TestPart2(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	assert.Equal(t, 4174379265, Part2(arr))
}
