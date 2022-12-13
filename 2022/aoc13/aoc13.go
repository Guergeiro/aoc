package aoc13

import (
	"encoding/json"
	"fmt"
	"sort"

	reader "github.com/guergeiro/aoc/2022/_reader"
)

type pair struct {
	left  []any
	right []any
}

func compare(left, right any) int {
	if lValue, lOk := left.(float64); lOk {
		if rValue, rOk := right.(float64); rOk {
			if lValue < rValue {
				return -1
			} else if lValue > rValue {
				return 1
			}
		}
	}

	if lValue, lOk := left.([]any); lOk {
		if rValue, rOk := right.(float64); rOk {
			return compare(lValue, []any{rValue})
		}
	}

	if lValue, lOk := left.(float64); lOk {
		if rValue, rOk := right.([]any); rOk {
			return compare([]any{lValue}, rValue)
		}
	}

	if lValue, lOk := left.([]any); lOk {
		if rValue, rOk := right.([]any); rOk {
			if len(lValue) == 0 && len(rValue) > 0 {
				return -1
			}
			if len(rValue) == 0 && len(lValue) > 0 {
				return 1
			}
			i := 0
			for ; i < len(lValue) && i < len(rValue); i += 1 {
				res := compare(lValue[i], rValue[i])
				if res != 0 {
					return res
				}
			}
			if len(lValue) < len(rValue) {
				return -1
			}
		}
	}

	return 0
}

func (p pair) isInRightOrder() bool {
	return compare(p.left, p.right) == -1
}

func getPairs(input []string) []pair {
	pairs := []pair{}
	for i := 0; i < len(input); i += 3 {
		s := input[i : i+2]

		p := pair{}
		if err := json.Unmarshal([]byte(s[0]), &p.left); err != nil {
			p.left = []any{}
		}
		if err := json.Unmarshal([]byte(s[1]), &p.right); err != nil {
			p.right = []any{}
		}

		pairs = append(pairs, p)
	}
	return pairs
}

func handle1(input []string) int {
	pairs := getPairs(input)
	sum := 0
	for i, pair := range pairs {
		if pair.isInRightOrder() {
			sum += i + 1
		}
	}
	return sum
}

func handle2(input []string) int {
	input = append(input, "", "[[2]]", "[[6]]")
	pairs := getPairs(input)
	allArrays := []any{}
	for _, pair := range pairs {
		allArrays = append(allArrays, pair.left, pair.right)
	}
	sort.Slice(allArrays, func(i, j int) bool {
		return compare(allArrays[i], allArrays[j]) == -1
	})
	idx := [2]int{0, 0}
	for i, pair := range allArrays {
		if out, err := json.Marshal(pair); err == nil {
			if "[[2]]" == string(out) {
				idx[0] = i + 1
			}
			if "[[6]]" == string(out) {
				idx[1] = i + 1
			}
		}
	}
	return idx[0] * idx[1]
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc13/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc13/input_question.txt")
	fmt.Println(handle1(arrExample))
	fmt.Println(handle1(arrQuestion))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc13/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc13/input_question.txt")
	fmt.Println(handle2(arrExample))
	fmt.Println(handle2(arrQuestion))
}
