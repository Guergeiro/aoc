package aoc05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc05_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	almanac := newAlmanac(lines)

	newAlmanac := almanac.
		processSeeds(almanac.seed2soil).
		processSeeds(almanac.soil2fertilizer).
		processSeeds(almanac.fertilizer2water).
		processSeeds(almanac.water2light).
		processSeeds(almanac.light2temparature).
		processSeeds(almanac.temperature2humidity).
		processSeeds(almanac.humidity2location)


	lowest := newAlmanac.seeds[0]
  for _, seed := range newAlmanac.seeds {
    if seed < lowest {
      lowest = seed
    }
  }

	return lowest
}

func Part2(lines []string) int {
	sum := 0
	return sum
}

type almanac struct {
	seeds                []int
	seed2soil            []converter
	soil2fertilizer      []converter
	fertilizer2water     []converter
	water2light          []converter
	light2temparature    []converter
	temperature2humidity []converter
	humidity2location    []converter
}

type converter struct {
	source      int
	destination int
	length      int
}

func (a almanac) processSeeds(converters []converter) almanac {
	newSeeds := []int{}
	for _, seed := range a.seeds {
		for _, conv := range converters {
			if newSeed := conv.process(seed); newSeed != seed {
				seed = newSeed
				break
			}
		}
		newSeeds = append(newSeeds, seed)
	}

	a.seeds = newSeeds
	return a
}

func (c converter) process(seed int) int {
	if seed < c.source {
		return seed
	}
	if seed > c.source+c.length {
		return seed
	}
	return seed - c.source + c.destination
}

func newAlmanac(lines []string) almanac {
	lines = strings.Split(strings.Join(lines, "\n"), "\n\n")

	almanac := almanac{}

	seeds := []int{}
	numberExp := regexp.MustCompile(`\d+`)
	seedsStr := numberExp.FindAllString(lines[0], -1)
	for _, seedStr := range seedsStr {
		if number, err := strconv.Atoi(seedStr); err == nil {
			seeds = append(seeds, number)
		}
	}
	almanac.seeds = seeds
	// skip 1 line because we use it for seeds
	lines = lines[1:]

	for _, line := range lines {
		mapLine := strings.Split(line, ":\n")
		converters := newConverters(mapLine[1])
		if strings.HasPrefix(mapLine[0], "seed") {
			almanac.seed2soil = converters
		} else if strings.HasPrefix(mapLine[0], "soil") {
			almanac.soil2fertilizer = converters
		} else if strings.HasPrefix(mapLine[0], "fertilizer") {
			almanac.fertilizer2water = converters
		} else if strings.HasPrefix(mapLine[0], "water") {
			almanac.water2light = converters
		} else if strings.HasPrefix(mapLine[0], "light") {
			almanac.light2temparature = converters
		} else if strings.HasPrefix(mapLine[0], "temperature") {
			almanac.temperature2humidity = converters
		} else if strings.HasPrefix(mapLine[0], "humidity") {
			almanac.humidity2location = converters
		}
	}
	return almanac
}

func newConverters(line string) []converter {
	converters := []converter{}

	lines := strings.Split(line, "\n")
	for _, l := range lines {
		converters = append(converters, newConverter(l))
	}

	return converters
}

func newConverter(line string) converter {
	numberExp := regexp.MustCompile(`\d+`)
	numbers := []int{}
	numbersStr := numberExp.FindAllString(line, -1)
	for _, numberStr := range numbersStr {
		if number, err := strconv.Atoi(numberStr); err == nil {
			numbers = append(numbers, number)
		}
	}
	return converter{
		destination: numbers[0],
		source:      numbers[1],
		length:      numbers[2],
	}
}
