package aoc07

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	array "github.com/guergeiro/aoc/2022/_array"
	reader "github.com/guergeiro/aoc/2022/_reader"
)

type filesystem interface {
	getName() string
	getFullName() string
	getSize() int
}

type utils interface {
	sizeBiggerThan(amount int) bool
	isRoot() bool
}

type directory struct {
	parent   filesystem
	children []filesystem
	name     string
}

type file struct {
	parent filesystem
	name   string
	size   int
}

func (f *file) getName() string {
	return f.name
}

func (f *file) getFullName() string {
	fullname := []string{}
	if f.parent != nil {
		fullname = append(fullname, f.parent.getFullName())
	}
	fullname = append(fullname, f.getName())
	return strings.Join(fullname, "/")
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) sizeBiggerThan(amount int) bool {
	if f.getSize() > amount {
		return true
	}
	return false
}

func (d *directory) getName() string {
	return d.name
}

func (d *directory) getFullName() string {
	fullname := []string{}
	if d.parent != nil {
		fullname = append(fullname, d.parent.getFullName())
	}
	fullname = append(fullname, d.getName())
	return strings.Join(fullname, "/")
}

func (d *directory) getSize() int {
	sum := 0
	for _, child := range d.children {
		sum += child.getSize()
	}
	return sum
}

func (d *directory) sizeBiggerThan(amount int) bool {
	if d.getSize() > amount {
		return true
	}
	return false
}

func (d *directory) isRoot() bool {
	return d.parent == nil
}

func (cur *directory) addChild(line string) {
	if splittedLine := strings.SplitN(line, " ", 2); splittedLine[0] == "dir" {
		// Directory
		newDir := createDir(splittedLine[1])
		newDir.parent = cur
		cur.children = append(cur.children, newDir)
	} else {
		// File
		newFile := createFile(splittedLine[1], splittedLine[0])
		newFile.parent = cur
		cur.children = append(cur.children, newFile)
	}
}

func (cur *directory) changeDir(newDir string) (*directory, error) {
	if newDir == ".." {
		return cur.parent.(*directory), nil
	}
	for _, child := range cur.children {
		if child.getName() == newDir {
			return child.(*directory), nil
		}
	}

	return nil, errors.New("No directory found")
}

func createFile(name string, sizeStr string) *file {
	file := &file{
		name: name,
		size: 0,
	}
	if value, ok := strconv.Atoi(sizeStr); ok == nil {
		file.size = value
	}
	return file
}

func createDir(name string) *directory {
	return &directory{
		name:     name,
		children: []filesystem{},
		parent:   nil,
	}
}

func parseCommand(cur *directory, line string) *directory {
	if line == "ls" {
		// we don't care...
		return cur
	}
	// change directory
	cmd := strings.SplitN(line, " ", 2)

	if newDir, err := cur.changeDir(cmd[1]); err == nil {
		return newDir
	}
	return nil
}

func createTree(input []string) *directory {
	root := createDir("/")
	cur := root
	shiftedInput := input[1:]
	for _, line := range shiftedInput {
		if strings.HasPrefix(line, "$ ") == false {
			// Listing dir
			cur.addChild(line)
		} else {
			extractedCmd := strings.SplitN(line, "$ ", 2)
			cur = parseCommand(cur, extractedCmd[1])
		}
	}
	return root
}

func (cur *directory) findDirSizesLessThan(amount int) []directory {
	found := []directory{}
	if cur.isRoot() && cur.sizeBiggerThan(amount) == false {
		found = append(found, *cur)
	}
	for _, child := range cur.children {
		if dir, isDirectory := child.(*directory); isDirectory {
			if dir.sizeBiggerThan(amount) == false {
				found = append(found, *dir)
			}
			found = append(found, dir.findDirSizesLessThan(amount)...)
		}
	}

	return found
}

func (cur *directory) findDirSizesBiggerThan(amount int) []directory {
	found := []directory{}
	if cur.isRoot() && cur.sizeBiggerThan(amount) {
		found = append(found, *cur)
	}
	for _, child := range cur.children {
		if dir, isDirectory := child.(*directory); isDirectory {
			if dir.sizeBiggerThan(amount) {
				found = append(found, *dir)
			}
			found = append(found, dir.findDirSizesBiggerThan(amount)...)
		}
	}

	return found
}

func handler1(input []string) int {
	root := createTree(input)
	dirs := root.findDirSizesLessThan(100000)
	sum := array.Reduce(dirs, func(acc int, cur directory) int {
		return acc + cur.getSize()
	}, 0)
	return sum
}

func handler2(input []string, totalSpace int, updateSpace int) int {
	root := createTree(input)

	unusedSpace := totalSpace - root.getSize()
	if unusedSpace >= updateSpace {
		return 0
	}
	requiredSpace := updateSpace - unusedSpace
	dirs := root.findDirSizesBiggerThan(requiredSpace)
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].getSize() < dirs[j].getSize()
	})
	return dirs[0].getSize()
}

func Part1() {
	arrExample := reader.ReadFileStrings("aoc07/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc07/input_question.txt")
	fmt.Println(handler1(arrExample))
	fmt.Println(handler1(arrQuestion))
}

func Part2() {
	arrExample := reader.ReadFileStrings("aoc07/input_example.txt")
	arrQuestion := reader.ReadFileStrings("aoc07/input_question.txt")
	fmt.Println(handler2(arrExample, 70000000, 30000000))
	fmt.Println(handler2(arrQuestion, 70000000, 30000000))
}
