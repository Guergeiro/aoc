package aoc12

import (
	"fmt"
	"sort"

	reader "github.com/guergeiro/aoc/2022/_reader"
)

type point struct {
	elevation  rune
	identifier rune
	distance   int
	higher     []*point
	lower      []*point
}

func (p point) isHigher(p2 point) bool {
	return p2.elevation-p.elevation > 0
}

func (p point) canClimb(p2 point) bool {
	if p2.elevation-p.elevation == 1 {
		return true
	}
	return false
}

func (p *point) calculate(currentDistance int) {
	if p.distance != -1 {
		if p.distance <= currentDistance {
			// Already went through this
			return
		}
	}
	p.distance = currentDistance
	for _, h := range p.higher {
		h.calculate(currentDistance + 1)
	}
	for _, l := range p.lower {
		l.calculate(currentDistance + 1)
	}
}

func createPoint(r rune) *point {
	return &point{
		elevation: r,
		higher:    []*point{},
		lower:     []*point{},
	}
}

func add(p1 *point, p2 *point) {
	if p1.isHigher(*p2) == true {
		if p1.canClimb(*p2) == true {
			p1.higher = append(p1.higher, p2)
		}
	} else {
		p1.lower = append(p1.lower, p2)
	}
}

func parsePoints(input []string) (*point, *point, [][]*point) {
	start := &point{}
	end := &point{}

	points := [][]*point{}

	for _, line := range input {
		r := []rune(line)
		current := []*point{}
		for _, p := range r {
			newPoint := &point{
				identifier: p,
				distance:   -1,
			}

			switch string(p) {
			case "S":
				newPoint.elevation = rune('a')
				start = newPoint
			case "E":
				newPoint.elevation = rune('z')
				end = newPoint
			default:
				newPoint.elevation = p
			}
			current = append(current, newPoint)
		}
		points = append(points, current)
	}

	for i := 0; i < len(points); i += 1 {
		for j := 0; j < len(points[i]); j += 1 {
			if i-1 >= 0 {
				add(points[i][j], points[i-1][j])
			}
			if j-1 >= 0 {
				add(points[i][j], points[i][j-1])
			}
			if i+1 < len(points) {
				add(points[i][j], points[i+1][j])
			}
			if j+1 < len(points[i]) {
				add(points[i][j], points[i][j+1])
			}
		}
	}

	return start, end, points
}

func handle1(input []string) int {
	start, end, _ := parsePoints(input)

	start.calculate(0)
	return end.distance
}
func handle2(input []string) int {
	_, end, points := parsePoints(input)

	allA := []*point{}

	for _, line := range points {
		for _, point := range line {
			if point.elevation == rune('a') {
				allA = append(allA, point)
			}
		}
	}

	possibleResults := []int{}

	for _, a := range allA {
		a.calculate(0)
		possibleResults = append(possibleResults, end.distance)
	}

	sort.Slice(possibleResults, func(i, j int) bool {
		return possibleResults[i] < possibleResults[j]
	})

	return possibleResults[0]
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc12/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc12/input_question.txt")
	fmt.Println(handle1(arrExample))
	fmt.Println(handle1(arrQuestion))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc12/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc12/input_question.txt")
	fmt.Println(handle2(arrExample))
	fmt.Println(handle2(arrQuestion))
}
