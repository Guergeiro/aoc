package aoc08

import (
	"fmt"

	reader "github.com/guergeiro/aoc/2022/_reader"
)

type tree struct {
	value rune
	row   int
	col   int
}

func (t tree) getLeftViewDistance(forest [][]tree) int {
	viewDistance := 0
	for j := t.col - 1; j >= 0; j -= 1 {
		viewDistance += 1
		if t.value <= forest[t.row][j].value {
			break
		}
	}
	return viewDistance
}

func (t tree) getRightViewDistance(forest [][]tree) int {
	viewDistance := 0
	for j := t.col + 1; j < len(forest[t.row]); j += 1 {
		viewDistance += 1
		if t.value <= forest[t.row][j].value {
			break
		}
	}
	return viewDistance
}

func (t tree) getUpperViewDistance(forest [][]tree) int {
	viewDistance := 0
	for i := t.row - 1; i >= 0; i -= 1 {
		viewDistance += 1
		if t.value <= forest[i][t.col].value {
			break
		}
	}
	return viewDistance
}

func (t tree) getLowerViewDistance(forest [][]tree) int {
	viewDistance := 0
	for i := t.row + 1; i < len(forest); i += 1 {
		viewDistance += 1
		if t.value <= forest[i][t.col].value {
			break
		}
	}
	return viewDistance
}

func (t tree) canSeeLeftEdge(forest [][]tree) bool {
	for j := t.col - 1; j >= 0; j -= 1 {
		if t.value <= forest[t.row][j].value {
			return false
		}
	}
	return true
}

func (t tree) canSeeRightEdge(forest [][]tree) bool {
	for j := t.col + 1; j < len(forest[t.row]); j += 1 {
		if t.value <= forest[t.row][j].value {
			return false
		}
	}
	return true
}

func (t tree) canSeeUpperEdge(forest [][]tree) bool {
	for i := t.row - 1; i >= 0; i -= 1 {
		if t.value <= forest[i][t.col].value {
			return false
		}
	}
	return true
}

func (t tree) canSeeLowerEdge(forest [][]tree) bool {
	for i := t.row + 1; i < len(forest); i += 1 {
		if t.value <= forest[i][t.col].value {
			return false
		}
	}
	return true
}

func (t tree) isVisible(forest [][]tree) bool {
	if t.canSeeLowerEdge(forest) {
		return true
	}
	if t.canSeeUpperEdge(forest) {
		return true
	}
	if t.canSeeLeftEdge(forest) {
		return true
	}
	if t.canSeeRightEdge(forest) {
		return true
	}
	return false
}

func parseForest(input []string) [][]tree {
	forest := [][]tree{}
	for i, line := range input {
		lineTree := []tree{}
		for j, t := range []rune(line) {
			lineTree = append(lineTree, tree{
				value: t,
				row:   i,
				col:   j,
			})
		}
		forest = append(forest, lineTree)
	}
	return forest
}

func countVisibleTrees(forest [][]tree) int {
	sum := len(forest)*2 + (len(forest[0])*2 - 4)

	for i := 1; i < len(forest)-1; i += 1 {
		for j := 1; j < len(forest[i])-1; j += 1 {
			if forest[i][j].isVisible(forest) {
				sum += 1
			}
		}
	}

	return sum
}

func getHighestScenic(forest [][]tree) int {
	max := 0

	for _, line := range forest {
		for _, tree := range line {
			topView := tree.getUpperViewDistance(forest)
			leftView := tree.getLeftViewDistance(forest)
			rightView := tree.getRightViewDistance(forest)
			bottomView := tree.getLowerViewDistance(forest)

			score := topView * bottomView * leftView * rightView
			if score > max {
				max = score
			}
		}
	}

	return max
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc08/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc08/input_question.txt")
	fmt.Println(countVisibleTrees(parseForest(arrExample)))
	fmt.Println(countVisibleTrees(parseForest(arrQuestion)))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc08/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc08/input_question.txt")
	fmt.Println(getHighestScenic(parseForest(arrExample)))
	fmt.Println(getHighestScenic(parseForest(arrQuestion)))
}
