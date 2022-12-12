package aoc11

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	array "github.com/guergeiro/aoc/2022/_array"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

type operation struct {
	opType  string
	target  string
	divider int
}

type test struct {
	divAmout int
	truthy   *monkey
	falsy    *monkey
}

type monkey struct {
	id               int
	items            []int
	operation        operation
	test             test
	inspectionAmount int
}

func (o operation) apply(value int) int {
	v, err := strconv.Atoi(o.target)
	if err != nil {
		v = value
	}
	switch o.opType {
	case "*":
		return v * value
	case "+":
		return v + value
	}

	return v
}

func (t test) assert(value int) *monkey {
	if value%t.divAmout == 0 {
		return t.truthy
	}
	return t.falsy
}

func (m *monkey) inspectNormal() {
	for _, item := range m.items {
		newItem := m.operation.apply(item)
		newItem = newItem / 3
		nextMonkey := m.test.assert(newItem)
		nextMonkey.items = append(nextMonkey.items, newItem)
		m.inspectionAmount += 1
	}
	m.items = append([]int{})
}

func (m *monkey) inspectModulo(allDivisors int) {
	for _, item := range m.items {
		newItem := m.operation.apply(item)
		newItem = newItem % allDivisors
		nextMonkey := m.test.assert(newItem)
		nextMonkey.items = append(nextMonkey.items, newItem)
		m.inspectionAmount += 1
	}
	m.items = append([]int{})
}

func createOperation(items []string) operation {
	return operation{
		opType: items[0],
		target: items[1],
	}
}

func createItem(items []string) []int {
	return array.Reduce(items, func(acc []int, cur string) []int {
		if value, err := strconv.Atoi(cur); err == nil {
			acc = append(acc, value)
		}
		return acc
	}, []int{})
}

func findMonkey(monkeys []*monkey, id int) (*monkey, error) {
	for _, monkey := range monkeys {
		if monkey.id == id {
			return monkey, nil
		}
	}
	return nil, errors.New("No monkey found")
}

func createTest(monkeys []*monkey, items []string) test {
	//19
	test := test{}
	for i, s := range items {
		s = strings.Trim(s, " ")
		switch i {
		case 0:
			r := []rune(s)
			r = r[19:]
			str := string(r)
			if value, err := strconv.Atoi(str); err == nil {
				test.divAmout = value
			}
		case 1:
			r := []rune(s)
			r = r[25:]
			str := string(r)
			if value, err := strconv.Atoi(str); err == nil {
				if monkey, ok := findMonkey(monkeys, value); ok == nil {
					test.truthy = monkey
				}
			}
		case 2:
			r := []rune(s)
			r = r[26:]
			str := string(r)
			if value, err := strconv.Atoi(str); err == nil {
				if monkey, ok := findMonkey(monkeys, value); ok == nil {
					test.falsy = monkey
				}
			}
		}
	}

	return test
}

func createMonkey(monkeyLines []string) *monkey {
	monkey := monkey{
		inspectionAmount: 0,
	}
	for i, s := range monkeyLines {
		s = strings.Trim(s, " ")
		switch i {
		case 0:
			r := []rune(s)
			monkey.id = int(r[7] - '0')
		case 1:
			r := []rune(s)
			r = r[16:]
			str := string(r)
			items := strings.Split(str, ", ")
			monkey.items = createItem(items)
		case 2:
			r := []rune(s)
			r = r[21:]
			str := string(r)
			monkey.operation = createOperation(strings.SplitN(str, " ", 2))
		}
	}
	return &monkey
}

func parseMonkeys(input []string) []*monkey {
	monkeys := []*monkey{}

	for idx := 0; idx < len(input); idx += 7 {
		// Create initial monkeys
		slice := input[idx : idx+3]
		monkeys = append(monkeys, createMonkey(slice))
	}

	for idx, monkeyId := 0, 0; idx < len(input); idx += 7 {
		// Populate test next monkeys
		slice := input[idx+3 : idx+6]
		test := createTest(monkeys, slice)
		monkeys[monkeyId].test = test
		monkeyId += 1
	}

	return monkeys
}

func tick(monkeys []*monkey, amount int, modulo bool) {
	allDivisors := array.Reduce(monkeys, func(acc int, curr *monkey) int {
		return acc * curr.test.divAmout
	}, 1)
	for i := 0; i < amount; i += 1 {
		for _, monkey := range monkeys {
			if modulo == false {
				monkey.inspectNormal()
			} else {
				monkey.inspectModulo(allDivisors)
			}
		}
	}
}

func calculateMultiplicationTopInspections(monkeys []*monkey, top int) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionAmount > monkeys[j].inspectionAmount
	})
	tops := monkeys[:top]
	return array.Reduce(tops, func(acc int, curr *monkey) int {
		return acc * curr.inspectionAmount
	}, 1)
}

func handler(input []string, tickAmount int, modulo bool) int {
	monkeys := parseMonkeys(input)
	tick(monkeys, tickAmount, modulo)
	return calculateMultiplicationTopInspections(monkeys, 2)
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc11/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc11/input_question.txt")
	fmt.Println(handler(arrExample, 20, false))
	fmt.Println(handler(arrQuestion, 20, false))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc11/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc11/input_question.txt")
	fmt.Println(handler(arrExample, 10000, true))
	fmt.Println(handler(arrQuestion, 10000, true))
}
