package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"sort"
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
	blocks := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n\n")
	pairs := make([]*Pair, len(blocks))
	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		var left, right []any
		json.Unmarshal([]byte(lines[0]), &left)
		json.Unmarshal([]byte(lines[1]), &right)
		pairs[i] = &Pair{
			left:  left,
			right: right,
		}
	}

	rightOrder := make([]int, 0)

	for i, pair := range pairs {
		if Compare(pair.left, pair.right) == 1 {
			rightOrder = append(rightOrder, i+1)
		}
	}

	return Sum(rightOrder...)
}

func part2(input string) int {
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	packets := make([]any, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var packet any
		json.Unmarshal([]byte(line), &packet)
		packets = append(packets, packet)
	}
	startDiv := []any{2.0}
	endDiv := []any{6.0}
	packets = append(packets, []any{startDiv})
	packets = append(packets, []any{endDiv})

	sort.SliceStable(packets, func(i, j int) bool {
		return Compare(packets[i], packets[j]) == 1
	})

	start, end := 0, 0
	for i, packet := range packets {
		if fmt.Sprintf("%v", packet) == "[[2]]" {
			start = i + 1
		}
		if fmt.Sprintf("%v", packet) == "[[6]]" {
			end = i + 1
		}
	}
	return start * end
}

type Pair struct {
	left, right []any
}

func Compare(left, right any) int {
	if left == nil || right == nil {
		return -1
	}
	if leftVal, leftOk := left.(float64); leftOk {
		if rightVal, rightOk := right.(float64); rightOk {
			if leftVal < rightVal {
				return 1
			} else if leftVal > rightVal {
				return 0
			}
			return -1
		}
		return Compare([]any{left}, right)
	}
	if _, rightOk := right.(float64); rightOk {
		return Compare(left, []any{right})
	}
	leftVal := left.([]any)
	rightVal := right.([]any)
	for i := 0; i < Min(len(leftVal), len(rightVal)); i++ {
		if Compare(leftVal[i], rightVal[i]) == 1 {
			return 1
		} else if Compare(leftVal[i], rightVal[i]) == 0 {
			return 0
		}
	}
	if len(leftVal) < len(rightVal) {
		return 1
	}
	if len(leftVal) > len(rightVal) {
		return 0
	}
	return -1
}

type Numbers interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func Sum[T Numbers](numbers ...T) T {
	sum := numbers[0]
	for _, n := range numbers[1:] {
		sum += n
	}
	return sum
}

func Min[T Numbers](a, b T) T {
	if a < b {
		return a
	}
	return b
}
