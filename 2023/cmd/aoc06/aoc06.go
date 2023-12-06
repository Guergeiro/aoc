package aoc06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc06_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	races := newRaces(func() ([]string, []string) {
		numberExp := regexp.MustCompile(`\d+`)
		timesStr := numberExp.FindAllString(lines[0], -1)
		distancesStr := numberExp.FindAllString(lines[1], -1)
		return timesStr, distancesStr
	})
	sum := 1
	for _, race := range races {
		records := 0
		results := race.race(*newBoat())
		for _, result := range results {
			if race.isRecord(result) {
				records += 1
			}
		}

		sum *= records
	}
	return sum
}

func Part2(lines []string) int {
	races := newRaces(func() ([]string, []string) {
		numberExp := regexp.MustCompile(`\d+`)
		timesStr := []string{strings.Join(numberExp.FindAllString(lines[0], -1), "")}
		distancesStr := []string{strings.Join(numberExp.FindAllString(lines[1], -1), "")}
		return timesStr, distancesStr
	})
	sum := 1
	for _, race := range races {
		records := 0
		results := race.race(*newBoat())
		for _, result := range results {
			if race.isRecord(result) {
				records += 1
			}
		}

		sum *= records
	}
	return sum
}

type race struct {
	time   int
	record int
}

func newRaces(numberParser func() ([]string, []string)) []race {
	races := []race{}

	timesStr, distancesStr := numberParser()
	for i, timeStr := range timesStr {
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			continue
		}
		distance, err := strconv.Atoi(distancesStr[i])
		if err != nil {
			continue
		}
		races = append(races, newRace(time, distance))
	}

	return races
}

func newRace(time int, distance int) race {
	return race{
		time:   time,
		record: distance,
	}
}

func (r race) race(b boat) []int {
	results := []int{}

	t := 1
	for t < r.time {
		travel := b.hold(t).release(r.time - t)
		results = append(results, travel)
		t += 1
	}
	return results
}

func (r race) isRecord(distance int) bool {
	return distance > r.record
}

type boat struct {
	speed int
}

func newBoat() *boat {
	return &boat{}
}

func (b *boat) hold(time int) *boat {
	b.speed = time
	return b
}
func (b *boat) release(time int) int {
	return time * b.speed
}
