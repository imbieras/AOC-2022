package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
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
	parsed := parseInput(input)
	sum := 0
	for _, l := range parsed {
		left, right := l[:len(l)/2], l[len(l)/2:]
		for _, c := range left {
			if strings.ContainsRune(right, c) {
				sum += priority(c)
				break
			}
		}
	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	sum := 0
	chunks := chunkSlice(parsed, 3)
	for _, chunk := range chunks {
		for _, c := range chunk[0] {
			if strings.ContainsRune(chunk[1], c) && strings.ContainsRune(chunk[2], c) {
				sum += priority(c)
				break
			}
		}
	}
	return sum
}

func priority(r rune) int {
	val := int(r)
	if val < 97 {
		return val - 38
	}
	return val - 96
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

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
