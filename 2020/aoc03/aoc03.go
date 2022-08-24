package aoc03

import (
	"github.com/guergeiro/aoc/2020/utils"
)

func Part1()(int, int) {
  arrExample := utils.ReadFileStrings("aoc03/input_example.txt")
  arrQuestion := utils.ReadFileStrings("aoc03/input_question.txt")

  _, treesExample := calculate(arrExample, [2]int{0, 1},[2]int {0, 3})
  _, treesQuestion := calculate(arrQuestion, [2]int{0, 1}, [2]int {0, 3})
  return treesExample, treesQuestion
}

func Part2()(int, int) {
  arrExample := utils.ReadFileStrings("aoc03/input_example.txt")
  arrQuestion := utils.ReadFileStrings("aoc03/input_question.txt")

  checks := [][2][2]int {
    {{0,1}, {0,1}},
    {{0,1}, {0,3}},
    {{0,1}, {0,5}},
    {{0,1}, {0,7}},
    {{0,2}, {0,1}},
  }

  exampleOutput := []int{}
  questionOutput := []int{}
  for _, check := range checks {
    _, treesExample := calculate(arrExample, check[0],check[1])
    exampleOutput = append(exampleOutput, treesExample)
    _, treesQuestion := calculate(arrQuestion, check[0], check[1])
    questionOutput = append(questionOutput, treesQuestion)
  }

  out1 := utils.Reduce(exampleOutput, func(acc int, curr int)int {
    return acc * curr
  }, 1)
  out2 := utils.Reduce(questionOutput, func(acc int, curr int)int {
    return acc * curr
  }, 1)

  return out1, out2
}

func calculate(arr []string, rowinfo [2]int, colinfo [2]int)(int, int) {
  squares := 0
  trees := 0
  col := colinfo[0]
  for row:=rowinfo[0]; row < len(arr); row+=rowinfo[1] {
    if string(arr[row][col % len(arr[row])]) == "#" {
        trees += 1
    } else {
      squares += 1
    }
    col += colinfo[1]
  }
  return squares, trees
}
