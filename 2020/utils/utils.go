package utils

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
  "path/filepath"
)

func readFile(path string) []string {
  absPath, err := filepath.Abs(path)
  if (err != nil) {
    fmt.Println(err)
  }

  file, err := os.Open(absPath)

  if (err != nil) {
    fmt.Println(err)
  }
  defer file.Close()

  output := []string{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    output = append(output, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
  return output
}

func ReadFileStrings(path string) []string {
  return readFile(path);
}

func ReadFileInts(path string) []int {
  lines := readFile(path)

  output := []int{}

  for _, line := range lines {
    value, err := strconv.Atoi(line)
    if err != nil {
      continue
    }
    output = append(output, value)
  }

  return output
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}
