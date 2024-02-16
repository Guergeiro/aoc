package aoc10

import (
	"fmt"
	"slices"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc10_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	m := newMaze(lines)

	fmt.Println(m.pipes[toIndex(2, 0)])
	fmt.Println(m.pipes[toIndex(3, 0)])
	fmt.Println(m.pipes[toIndex(4, 0)])
	fmt.Println(m.pipes[toIndex(4, 1)])
	fmt.Println(m.pipes[toIndex(3, 1)])
	fmt.Println(m.pipes[toIndex(3, 2)])
	fmt.Println(m.pipes[toIndex(3, 3)])
	fmt.Println(m.pipes[toIndex(3, 4)])

	return 0
}

func Part2(lines []string) int {
	sum := 0
	return sum
}

var num_cols, num_rows int = 0, 0

type maze struct {
	pipes map[int]*pipe
	start *pipe
}

type pipe struct {
	row  int
	col  int
	char rune
}

func (m maze) loop() []*pipe {
  loop := []*pipe{}

  return loop
}

func (p *pipe) next(pipes map[int]*pipe, from *pipe) *pipe {
  if top, err := p.top(pipes); err == nil && top != from {
  }
  if right, err := p.right(pipes); err == nil && right != from {
  }
  if bottom, err := p.bottom(pipes); err == nil && bottom != from {
  }
  if left, err := p.left(pipes); err == nil && left != from {
  }
}

func (p *pipe) top(pipes map[int]*pipe) (*pipe, error) {
  if top, exists := pipes[toIndex(p.row - 1, p.col)]; exists {
    if isTop(p.char, top.char) {
      return top, nil
    }
  }
  return nil, fmt.Errorf("No top")
}
func (p *pipe) right(pipes map[int]*pipe) (*pipe, error) {
  if right, exists := pipes[toIndex(p.row, p.col + 1)]; exists {
    if isLeft(p.char, right.char) {
      return right, nil
    }
  }
  return nil, fmt.Errorf("No top")
}
func (p *pipe) bottom(pipes map[int]*pipe) (*pipe, error) {
  if bottom, exists := pipes[toIndex(p.row + 1, p.col)]; exists {
    if isTop(p.char, bottom.char) {
      return bottom, nil
    }
  }
  return nil, fmt.Errorf("No top")
}
func (p *pipe) left(pipes map[int]*pipe) (*pipe, error) {
  if left, exists := pipes[toIndex(p.row, p.col - 1)]; exists {
    if isLeft(p.char, left.char) {
      return left, nil
    }
  }
  return nil, fmt.Errorf("No top")
}

func newMaze(lines []string) maze {
	num_rows = len(lines)
	num_cols = len(lines[0])

	var startPipe *pipe
	pipes := map[int]*pipe{}
	for row, line := range lines {
		for col, char := range []rune(line) {
			index := toIndex(row, col)
			pipes[index] = &pipe{
				row:  row,
				col:  col,
				char: char,
			}
			if char == 'S' {
				startPipe = pipes[index]
			}
		}
	}

	m := maze{
		start: startPipe,
		pipes: pipes,
	}

	return m
}

func isTop(value rune, next rune) bool {
	possibilities := []rune{}
	switch value {
	case '|':
		possibilities = append(possibilities, '|', '7', 'F')
		break
	case 'L':
		possibilities = append(possibilities, '|', '7', 'F')
		break
	case 'J':
		possibilities = append(possibilities, '|', '7', 'F')
		break
	case 'S':
		possibilities = append(possibilities, '|', '7', 'F')
		break
	}
	return slices.Contains(possibilities, next)
}
func isBottom(value rune, next rune) bool {
	possibilities := []rune{}
	switch value {
	case '|':
		possibilities = append(possibilities, '|', 'L', 'J')
		break
	case 'F':
		possibilities = append(possibilities, '|', 'L', 'J')
		break
	case '7':
		possibilities = append(possibilities, '|', 'L', 'J')
		break
	case 'S':
		possibilities = append(possibilities, '|', 'L', 'J')
		break
	}
	return slices.Contains(possibilities, next)
}
func isRight(value rune, next rune) bool {
	possibilities := []rune{}
	switch value {
	case '-':
		possibilities = append(possibilities, '-', 'J', '7')
		break
	case 'L':
		possibilities = append(possibilities, '-', 'J', '7')
		break
	case 'F':
		possibilities = append(possibilities, '-', 'J', '7')
		break
	case 'S':
		possibilities = append(possibilities, '-', 'J', '7')
		break
	}
	return slices.Contains(possibilities, next)
}
func isLeft(value rune, next rune) bool {
	possibilities := []rune{}
	switch value {
	case '-':
		possibilities = append(possibilities, '-', 'L', 'F')
		break
	case 'J':
		possibilities = append(possibilities, '-', 'L', 'F')
		break
	case '7':
		possibilities = append(possibilities, '-', 'L', 'F')
		break
	case 'S':
		possibilities = append(possibilities, '-', 'L', 'F')
		break
	}
	return slices.Contains(possibilities, next)
}

func (p *pipe) toIndex() int {
	return toIndex(p.row, p.col)
}

func toIndex(row, col int) int {
	return row*num_cols + col
}

func toCoordinate(index int) (int, int) {
	row := index / num_cols
	col := index % num_cols
	return row, col
}
