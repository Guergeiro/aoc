package aoc09

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	reader "github.com/guergeiro/aoc/2022/_reader"
)

type location struct {
	xLocation int
	yLocation int
}

type node struct {
	location location
	visited  []location
	tails    []*node // head keeps track of tail
}

type move struct {
	direction string
	amount    int
}

func (l *location) changeLocation(xAmount, yAmount int) {
	l.xLocation += xAmount
	l.yLocation += yAmount
}

func (l location) equals(l2 location) bool {
	if l.xLocation != l2.xLocation {
		return false
	}
	if l.yLocation != l2.yLocation {
		return false
	}
	return true
}

func (l location) difference(l2 location) (int, int) {
	xDifference := l.xLocation - l2.xLocation
	yDifference := l.yLocation - l2.yLocation
	return xDifference, yDifference
}

func (n *node) appendVisited(l location) {
	for _, location := range n.visited {
		if l.equals(location) {
			// One location already exists
			return
		}
	}
	n.visited = append(n.visited, l)
}

func (head *node) moveTails() {
	previousTail := head
	for _, tail := range head.tails {
		xDiff, yDiff := previousTail.location.difference(tail.location)
		if math.Abs(float64(xDiff)) > 1 || math.Abs(float64(yDiff)) > 1 {
			xMove := 0
			yMove := 0
			if xDiff < 0 {
				xMove = -1
			} else if xDiff > 0 {
				xMove = 1
			}
			if yDiff < 0 {
				yMove = -1
			} else if yDiff > 0 {
				yMove = 1
			}
			tail.location.changeLocation(xMove, yMove)
			tail.appendVisited(tail.location)
		}
		previousTail = tail
	}
}

func (head *node) moveHead(move move) {
	for step := 0; step < move.amount; step += 1 {
		switch move.direction {
		case "R":
			head.location.changeLocation(1, 0)
		case "L":
			head.location.changeLocation(-1, 0)
		case "U":
			head.location.changeLocation(0, 1)
		case "D":
			head.location.changeLocation(0, -1)
		}
		head.appendVisited(head.location)
		head.moveTails()
	}
}

func parseAllMoves(input []string) []move {
	moves := []move{}
	for _, line := range input {
		splittedLine := strings.SplitN(line, " ", 2)
		if amount, err := strconv.Atoi(splittedLine[1]); err == nil {
			moves = append(moves, move{
				direction: splittedLine[0],
				amount:    amount,
			})
		}
	}
	return moves
}

func handler(input []string, tailAmount int) int {
	moves := parseAllMoves(input)
	defaultVisited := []location{
		{
			xLocation: 0,
			yLocation: 0,
		},
	}
	head := &node{
		location: location{
			xLocation: 0,
			yLocation: 0,
		},
		visited: defaultVisited,
		tails:   []*node{},
	}
	for i := 0; i < tailAmount; i += 1 {
		tail := &node{
			location: location{
				xLocation: 0,
				yLocation: 0,
			},
			visited: defaultVisited,
			tails:   nil,
		}
		head.tails = append(head.tails, tail)
	}
	for _, move := range moves {
		head.moveHead(move)
	}
	amountVisited := []int{}
	for _, tail := range head.tails {
		amountVisited = append(amountVisited, len(tail.visited))
	}
	return amountVisited[len(amountVisited)-1]
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc09/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc09/input_question.txt")
	fmt.Println(handler(arrExample, 1))
	fmt.Println(handler(arrQuestion, 1))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc09/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc09/input_question.txt")
	fmt.Println(handler(arrExample, 9))
	fmt.Println(handler(arrQuestion, 9))
}
