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
	return findMarker(input, 4)
}

func part2(input string) int {
	return findMarker(input, 14)
}

func findMarker(input string, markerLen int) int {
	for c := 0; c < len(input); c++ {
		if c+markerLen > len(input) {
			return -1
		}
		if isUnique(input[c : c+markerLen]) {
			return c + markerLen
		}
	}
	return -1
}

func isUnique(input string) bool {
	for i := 0; i < len(input); i++ {
		if strings.IndexByte(input[i+1:], input[i]) != -1 {
			return false
		}
	}
	return true
}
