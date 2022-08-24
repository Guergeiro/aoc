package aoc02

import (
	"strconv"
	"strings"
	"github.com/guergeiro/aoc/2020/utils"
)

type password struct {
  value string
  lowerBound int
  higherBound int
  requiredChar string
  isValid1 bool
  isValid2 bool
}

func Part1()(int, int) {
  arrExample := utils.ReadFileStrings("aoc02/input_example.txt")
  arrQuestion := utils.ReadFileStrings("aoc02/input_question.txt")

  return calculate1(arrExample), calculate1(arrQuestion)
}

func Part2()(int, int) {
  arrExample := utils.ReadFileStrings("aoc02/input_example.txt")
  arrQuestion := utils.ReadFileStrings("aoc02/input_question.txt")

  return calculate2(arrExample), calculate2(arrQuestion)
}

func calculate1(lines []string)int {
  passwords := []password{}

  for _, line := range lines {
    passwords = append(passwords, splitPassword(line))
  }

  sum := 0

  for _, password := range passwords {
    if password.isValid1 {
      sum += 1
    }
  }
  return sum
}

func calculate2(lines []string)int {
  passwords := []password{}

  for _, line := range lines {
    passwords = append(passwords, splitPassword(line))
  }

  sum := 0

  for _, password := range passwords {
    if password.isValid2 {
      sum += 1
    }
  }
  return sum
}

func splitPassword(line string)password {
  x := strings.Split(line, ": ")
  pw := x[1]
  y := strings.Split(x[0], " ")
  char := y[1]

  pwStruct := password{
    value: pw,
    requiredChar: char,
  }
  bounds := strings.Split(y[0], "-")
  lower, err := strconv.Atoi(bounds[0])
  if err == nil {
    pwStruct.lowerBound = lower
  }

  higher, err := strconv.Atoi(bounds[1])
  if err == nil {
    pwStruct.higherBound = higher
  }
  pwStruct.isValid1 = isValidPassword1(pwStruct)
  pwStruct.isValid2 = isValidPassword2(pwStruct)
  return pwStruct
}

func isValidPassword1(pw password)bool {
  occurrences := strings.Count(pw.value, pw.requiredChar)
  return occurrences >= pw.lowerBound && occurrences <= pw.higherBound
}

func isValidPassword2(pw password)bool {
  if string(pw.value[pw.lowerBound - 1]) != pw.requiredChar {
    return string(pw.value[pw.higherBound - 1]) == pw.requiredChar
  }
  if string(pw.value[pw.higherBound - 1]) != pw.requiredChar {
    return string(pw.value[pw.lowerBound - 1]) == pw.requiredChar
  }
  return false
}
