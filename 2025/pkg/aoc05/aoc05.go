package aoc05

import (
	"regexp"
	"slices"
	"strconv"
)

func Part1(arr []string) int {
	inventory := newInventory(arr)
	return len(inventory.GetFreshIngredients())
}

func Part2(arr []string) int {
	inventory := newInventory(arr)
	return inventory.GetPossibleFreshIngredients()
}

type inventory struct {
	ids         [][]int
	ingredients []int
}

func (i inventory) GetPossibleFreshIngredients() int {
	possibleFreshIngredients := 0
	for _, limits := range i.ids {
		possibleFreshIngredients += (limits[1] + 1 - limits[0])
	}
	return possibleFreshIngredients
}

func (i inventory) GetFreshIngredients() []int {

	output := []int{}
	for _, ingredient := range i.ingredients {
		for _, limits := range i.ids {
			if limits[0] <= ingredient && limits[1] >= ingredient {
				output = append(output, ingredient)
			}
		}
	}

	return output
}

var numberExp = regexp.MustCompile(`\d+`)

func newInventory(arr []string) inventory {
	inventory := inventory{
		ids:         [][]int{},
		ingredients: []int{},
	}

	atIngredients := false
	for _, line := range arr {
		if line == "" {
			// change from parsing Ids to Ingredients
			atIngredients = true
			continue
		}

		if atIngredients == false {
			// parse Ids
			digits := numberExp.FindAllString(line, -1)
			number1, err := strconv.Atoi(digits[0])
			if err != nil {
				continue
			}
			number2, err := strconv.Atoi(digits[1])
			if err != nil {
				continue
			}
			sorted := slices.Sorted(slices.Values([]int{number1, number2}))
			inventory.ids = append(inventory.ids, sorted)
		} else {
			// parse Ingredients
			digits := numberExp.FindAllString(line, -1)
			number, err := strconv.Atoi(digits[0])
			if err != nil {
				continue
			}
			inventory.ingredients = append(inventory.ingredients, number)
		}

	}

	inventory.ids = mergeRanges(inventory.ids)

	return inventory
}

func mergeRanges(ranges [][]int) [][]int {
	// first let's sort by lower bound
	sortedRanges := slices.SortedFunc(
		slices.Values(ranges),
		func(a, b []int) int {
			return a[0] - b[0]
		},
	)

	merged := [][]int{sortedRanges[0]}

	for i := 1; i < len(sortedRanges); i++ {
		last := merged[len(merged)-1]
		current := sortedRanges[i]

		// Check if ranges overlap or are adjacent
		if current[0] <= last[1]+1 {
			// Merge by extending the end
			last[1] = max(last[1], current[1])
		} else {
			// No overlap, add as new range
			merged = append(merged, current)
		}
	}

	return merged
}
