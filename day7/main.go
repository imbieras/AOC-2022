package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"strconv"
	"strings"
)

var (
	inputFile = util.ReadFile("./input.txt")
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(inputFile)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(inputFile)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	root := execHistory(parseInput(input))
	sizes := Filter(root, func(e *Entry) bool {
		return e.Size() < 100000
	})
	return Sum(sizes)
}

func part2(input string) int {
	root := execHistory(parseInput(input))
	needed := 30000000 - (70000000 - root.Size())
	sizes := Filter(root, func(e *Entry) bool {
		return e.Size() >= needed
	})
	return Min(sizes)
}

type Entry struct {
	name     string
	dir      bool
	parent   *Entry
	children []*Entry
	size     int
}

func makeDir(name string) *Entry {
	return &Entry{
		name:     name,
		dir:      true,
		children: make([]*Entry, 0),
	}
}

func makeFile(name string, size int) *Entry {
	return &Entry{
		name: name,
		size: size,
	}
}

func (e *Entry) Add(child *Entry) {
	child.parent = e
	e.children = append(e.children, child)
}

func (e *Entry) Size() int {
	if e.dir {
		sum := 0
		for _, child := range e.children {
			sum += child.Size()
		}
		return sum
	}
	return e.size
}

func (e *Entry) findChild(name string) *Entry {
	for _, child := range e.children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func parseInput(input string) [][]string {
	blocks := make([][]string, 0)
	block := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$") {
			blocks = append(blocks, block)
			block = []string{line}
			continue
		}
		block = append(block, line)
	}
	blocks = append(blocks, block)
	return blocks
}

func execHistory(history [][]string) *Entry {
	root := makeDir("/")
	currentDir := root
	for _, block := range history[1:] {
		cmd, output := block[0], block[1:]
		if strings.HasPrefix(cmd, "$ cd") {
			fields := strings.Fields(cmd)
			if fields[2] == ".." {
				currentDir = currentDir.parent
			} else {
				if child := currentDir.findChild(fields[2]); child != nil {
					currentDir = child
				}
			}
		}

		if strings.HasPrefix(cmd, "$ ls") {
			for _, line := range output {
				fields := strings.Fields(line)
				if fields[0] == "dir" {
					currentDir.Add(makeDir(fields[1]))
				} else {
					value, _ := strconv.Atoi(fields[0])
					currentDir.Add(makeFile(fields[1], value))
				}
			}
		}
	}
	return root
}

func Filter(e *Entry, fn func(*Entry) bool) []int {
	sizes := make([]int, 0)
	if e.dir {
		if fn(e) {
			sizes = append(sizes, e.Size())
		}
		for _, child := range e.children {
			sizes = append(sizes, Filter(child, fn)...)
		}
	}
	return sizes
}

type Numbers interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func Sum[T Numbers](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Min[T Numbers](numbers []T) T {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}
