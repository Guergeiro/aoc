package aoc04

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc04_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	cards := parseCards(lines)
	sum := 0
	for _, card := range cards {
		sum += card.getCardWinnings()
	}
	return sum
}

func Part2(lines []string) int {
	cards := parseCards(lines)
	cardsAmounts := map[int]int{}
	for _, card := range cards {
		cardsAmounts[card.id] = 1
	}
	for _, card := range cards {
		wins := card.winnerNumbers()
		for j := range wins {

			originalAmount, exists := cardsAmounts[card.id+j+1]
			if exists == false {
				continue
			}
			toAdd, exists := cardsAmounts[card.id]
			if exists == false {
				continue
			}
			cardsAmounts[card.id+j+1] = originalAmount + toAdd
		}
	}
	sum := 0
	for _, card := range cardsAmounts {
		sum += card
	}
	return sum
}

type card struct {
	winning map[int]bool
	owned   map[int]bool
	id      int
}

func parseCards(lines []string) []*card {
	cards := []*card{}

	for _, line := range lines {
		if card, err := newCard(line); err == nil {
			cards = append(cards, card)
		}
	}

	return cards
}

func newCard(line string) (*card, error) {
	splits := strings.Split(line, ": ")

	gameIdExp := regexp.MustCompile(`\d+`)
	gameIdStr := gameIdExp.FindString(splits[0])

	games := strings.Split(splits[1], " | ")
	winningGameStr, ownGameStr := games[0], games[1]
	numberExp := regexp.MustCompile(`\d+`)
	winningGames := numberExp.FindAllString(winningGameStr, -1)
	ownGames := numberExp.FindAllString(ownGameStr, -1)

	winning := map[int]bool{}
	for _, win := range winningGames {
		if number, err := strconv.Atoi(win); err == nil {
			winning[number] = true
		}
	}

	owned := map[int]bool{}
	for _, own := range ownGames {
		if number, err := strconv.Atoi(own); err == nil {
			owned[number] = true
		}
	}

	id, err := strconv.Atoi(gameIdStr)
	if err != nil {
		return nil, err
	}
	return &card{
		id:      id,
		winning: winning,
		owned:   owned,
	}, nil
}

func (c *card) getCardWinnings() int {
	return calculate(len(c.winnerNumbers()))
}

func calculate(wins int) int {
	return int(math.Pow(2, float64(wins-1)))
}

func (c *card) winnerNumbers() []int {
	winnerNumbers := []int{}
	for own := range c.owned {
		if _, exists := c.winning[own]; exists {
			winnerNumbers = append(winnerNumbers, own)
		}
	}
	return winnerNumbers
}
