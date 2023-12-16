package aoc08

import (
	"fmt"
	"regexp"
	"strings"

	reader "github.com/guergeiro/aoc/2023/internal"
)

func Solve() {
	if arr, err := reader.ReadFileLines("assets/aoc08_input.txt"); err == nil {
		fmt.Println(Part1(arr))
		fmt.Println(Part2(arr))
	}
}

func Part1(lines []string) int {
	tree := newTree(lines, "AAA")
	return tree.stepsUntilEnd("ZZZ")[0]
}

func Part2(lines []string) int {
	tree := newTree(lines, "A")
	steps := tree.stepsUntilEnd("Z")
	return lcmAll(steps[0], steps[1:]...)
}

type tree struct {
	heads        []*node
	instructions []rune
}

type node struct {
	left  *node
	right *node
	value string
}

func (t tree) stepsUntilEnd(value string) []int {
	steps := []int{}
	for _, head := range t.heads {
		amountSteps := head.find(value, t.instructions, 0)
		steps = append(steps, amountSteps)
	}
	return steps
}

func (n *node) find(value string, instructions []rune, idx int) int {
	if strings.HasSuffix(n.value, value) {
		return 0
	}

	if idx >= len(instructions) {
		idx = idx % len(instructions)
	}

	if instructions[idx] == 'L' {
		return 1 + n.left.find(value, instructions, idx+1)
	}

	return 1 + n.right.find(value, instructions, idx+1)
}

func newTree(lines []string, firstElementSuffix string) tree {
	instructions := lines[0]
	lines = lines[2:]

	existingNodes := map[string]*node{}

	exp := regexp.MustCompile(`[A-Za-z0-9]+`)
	for _, line := range lines {
		matches := exp.FindAllString(line, -1)

		curStr := matches[0]
		leftStr := matches[1]
		rightStr := matches[2]

		leftNode, exists := existingNodes[leftStr]
		if exists == false {
			leftNode = newNode(leftStr, nil, nil)
			existingNodes[leftStr] = leftNode
		}
		rightNode, exists := existingNodes[rightStr]
		if exists == false {
			rightNode = newNode(rightStr, nil, nil)
			existingNodes[rightStr] = rightNode
		}
		currentNode, exists := existingNodes[curStr]
		if exists == false {
			currentNode = newNode(curStr, leftNode, rightNode)
			existingNodes[curStr] = currentNode
		}
		currentNode.left = leftNode
		currentNode.right = rightNode
	}

	heads := []*node{}
	for _, n := range existingNodes {
		if strings.HasSuffix(n.value, firstElementSuffix) {
			heads = append(heads, n)
		}
	}

	return tree{
		instructions: []rune(instructions),
		heads:        heads,
	}
}

func newNode(value string, left *node, right *node) *node {
	return &node{
		value: value,
		left:  left,
		right: right,
	}
}

// gcd,lcm,lcmAll inspired by:
// https://stackoverflow.com/questions/31302054/how-to-find-the-least-common-multiple-of-a-range-of-numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
func lcmAll(a int, bs ...int) int {
	result := a
	for _, b := range bs {
		result = lcm(result, b)
	}

	return result
}
