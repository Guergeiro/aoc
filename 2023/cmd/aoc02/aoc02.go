package aoc02

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc02_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		if game, err := newGame(line); err == nil {
			if game.isValid(func(p play) bool {
				if p.red > 12 {
					return false
				}
				if p.green > 13 {
					return false
				}
				if p.blue > 14 {
					return false
				}
				return true
			}) {
				sum += game.id
			}
		}
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		if game, err := newGame(line); err == nil {
			sum += (game.highestRed() * game.highestGreen() * game.highestBlue())
		}
	}
	return sum
}

type game struct {
	id    int
	plays []play
}

type play struct {
	red   int
	green int
	blue  int
}

func (g *game) isValid(f func(play) bool) bool {
	for _, play := range g.plays {
		if f(play) == false {
			return false
		}
	}
	return true
}

func (g *game) highestRed() int {
	return g.highestField("red")
}

func (g *game) highestGreen() int {
	return g.highestField("green")
}

func (g *game) highestBlue() int {
	return g.highestField("blue")
}

func (g *game) highestField(field string) int {
	highest := 0
	for _, play := range g.plays {
		playReflection := reflect.ValueOf(play)
		colorReflection := reflect.Indirect(playReflection).FieldByName(field)
		digit := int(colorReflection.Int())
		if digit == -1 {
			continue
		}
		if digit > highest {
			highest = digit
		}
	}
	return highest
}

func newGame(line string) (*game, error) {
	splittedLine := strings.Split(line, ": ")
	strId := strings.Split(splittedLine[0], " ")[1]
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, err
	}

	strPlays := strings.Split(splittedLine[1], "; ")
	game := &game{
		id:    id,
		plays: dividePlays(strPlays),
	}

	return game, nil
}

func dividePlays(strPlays []string) []play {
	plays := []play{}
	for _, strPlay := range strPlays {
		playline := strings.Split(strPlay, ", ")
		plays = append(plays, newPlay(playline))
	}
	return plays
}

func newPlay(playline []string) play {
	red, green, blue := -1, -1, -1
	for _, play := range playline {
		splittedPlay := strings.Split(play, " ")
		if digit, err := strconv.Atoi(splittedPlay[0]); err == nil {
			switch splittedPlay[1] {
			case "red":
				red = digit
				break
			case "green":
				green = digit
				break
			case "blue":
				blue = digit
				break
			}
		}
	}

	return play{
		red:   red,
		green: green,
		blue:  blue,
	}
}
