package aoc07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc07_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	hands := newHands(
		lines,
		func(card rune) int {
			values := map[rune]int{}
			keys := []rune{
				'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A',
			}
			for i, key := range keys {
				values[key] = i
			}
			return values[card]
		})

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}
	return sum
}

func Part2(lines []string) int {
	return 0
}

type hand struct {
	cards []rune
	bid   int
}

func (h hand) strength() int {
	amount := map[rune]int{}
	for _, c := range h.cards {
		if value, exists := amount[c]; exists {
			amount[c] = value + 1
		} else {
			amount[c] = 1
		}
	}
	max, length := 0, 0
	for _, value := range amount {
		if value > max {
			max = value
		}
		length += 1
	}
	return max - length
}

func newHands(
	lines []string,
	cardValue func(rune) int,
) []hand {
	hands := []hand{}
	for _, line := range lines {
		if hand, err := newHand(line); err == nil {
			hands = append(hands, hand)
		}
	}
	slices.SortFunc(hands, func(h1, h2 hand) int {
		h1Strength := h1.strength()
		h2Strength := h2.strength()
		if h1Strength > h2Strength {
			return 1
		} else if h1Strength < h2Strength {
			return -1
		}

		for i, c := range h1.cards {
			h1CardValue := cardValue(c)
			h2CardValue := cardValue(h2.cards[i])
			if h1CardValue > h2CardValue {
				return 1
			} else if h1CardValue < h2CardValue {
				return -1
			}
		}

		return 0
	})
	return hands
}

func newHand(line string) (hand, error) {
	splitted := strings.Split(line, " ")
	cards := splitted[0]
	bid, err := strconv.Atoi(splitted[1])
	if err != nil {
		return hand{}, err
	}

	return hand{
		cards: []rune(cards),
		bid:   bid,
	}, nil
}
