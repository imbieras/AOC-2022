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

func part1(input string) string {
	stacks, instrl := parseInput(input)
	for _, instr := range instrl {
		count, from, to := parseInstr(instr)
		fromStack := stacks[from]
		toStack := stacks[to]
		crates, fromStack := PopN(fromStack, count)
		Reverse(crates)
		toStack = append(toStack, crates...)
		stacks[from] = fromStack
		stacks[to] = toStack
	}
	var sb strings.Builder
	for _, stack := range stacks {
		sb.WriteString(stack[len(stack)-1])
	}
	return sb.String()
}

func part2(input string) string {
	stacks, instrl := parseInput(input)
	for _, instr := range instrl {
		count, from, to := parseInstr(instr)
		fromStack := stacks[from]
		toStack := stacks[to]
		crates, fromStack := PopN(fromStack, count)
		toStack = append(toStack, crates...)
		stacks[from] = fromStack
		stacks[to] = toStack
	}
	var sb strings.Builder
	for _, stack := range stacks {
		sb.WriteString(stack[len(stack)-1])
	}
	return sb.String()
}

func chunkSlice[T any](slice []T, size int) [][]T {
	length := len(slice)
	chunks := make([][]T, 0)
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func PopN[T any](slice []T, n int) ([]T, []T) {
	return slice[len(slice)-n:], slice[:len(slice)-n]
}

func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type Numbers interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func Between[T Numbers](i, min, max T) bool {
	return min <= i && i <= max
}

func parseInput(input string) ([][]string, []string) {
	part := strings.Split(input, "\n\n")
	stackLine := strings.Split(part[0], "\n")
	instrLine := strings.Split(part[1], "\n")
	Reverse(stackLine)
	crates := chunkSlice([]rune(stackLine[0]), 4)
	stacks := make([][]string, len(crates))
	for i := range stacks {
		stacks[i] = make([]string, 0)
	}
	for _, l := range stackLine[1:] {
		groups := chunkSlice([]rune(l), 4)
		for i, group := range groups {
			if Between(group[1], 65, 90) {
				stacks[i] = append(stacks[i], string(rune(group[1])))
			}
		}
	}
	return stacks, instrLine
}

func parseInstr(instr string) (int, int, int) {
	fields := strings.Fields(instr)
	count, _ := strconv.Atoi(fields[1])
	from, _ := strconv.Atoi(fields[3])
	from--
	to, _ := strconv.Atoi(fields[5])
	to--
	return count, from, to
}
