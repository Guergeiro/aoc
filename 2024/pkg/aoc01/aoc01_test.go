package aoc01

import (
	"slices"
	"testing"

	reader "github.com/guergeiro/aoc/2024/internal"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	left, right, err := newLists(arr)
	assert.NoError(t, err)

	assert.Len(t, left, 6)
	assert.Len(t, right, 6)

	sortedLeft := slices.Sorted(slices.Values(left))
	sortedRight := slices.Sorted(slices.Values(right))

	listDifferences := getListsDifferences(sortedLeft, sortedRight)

	sum := reduceList(listDifferences)

	assert.Equal(t, 11, sum)
}

func TestPart2(t *testing.T) {
	arr, err := reader.ReadFileLines("example1.txt")

	assert.NoError(t, err)

	left, right, err := newLists(arr)
	assert.NoError(t, err)

	assert.Len(t, left, 6)
	assert.Len(t, right, 6)

	counted := convertToRepeatedElements(right)

	similarities := getSimilarities(left, counted)

	sum := reduceList(similarities)

	assert.Equal(t, 31, sum)
}
