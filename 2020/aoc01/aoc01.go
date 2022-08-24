package aoc01

import (
	"github.com/guergeiro/aoc/2020/utils"
)

func Part1()(int, int) {
  arrExample := utils.ReadFileInts("aoc01/input_example.txt")
  arrQuestion := utils.ReadFileInts("aoc01/input_question.txt")

  return calculate(arrExample), calculate(arrQuestion)
}

func calculate(arr []int)int {
  const year = 2020

  for _, value := range arr {
    complement := year - value
    if existsInArray(arr, complement) {
      return complement * value
    }
  }

  return 0
}

func Part2()(int, int) {
  arrExample := utils.ReadFileInts("aoc01/input_example.txt")
  arrQuestion := utils.ReadFileInts("aoc01/input_question.txt")

  return calculate2(arrExample), calculate2(arrQuestion)
}

func calculate2(arr []int)int {
  const year = 2020

  for _, i := range arr {
    for _, j := range arr {
      complement := year - i - j
      if existsInArray(arr, complement) {
        return complement * i * j
      }
    }
  }

  return 0
}

func existsInArray(arr []int, val int)bool {
  for _,value := range arr {
    if val == value {
      return true
    }
  }

  return false
}

