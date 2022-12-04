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
	parsed := parseInput(input)
	count := 0
	for i := 0; i < len(parsed); i++ {
		p := parsed[i]
		overlap := Overlap(p[0], p[1], p[2], p[3])
		if overlap == 2 {
			count++
		}
	}
	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	count := 0
	for i := 0; i < len(parsed); i++ {
		p := parsed[i]
		overlap := Overlap(p[0], p[1], p[2], p[3])
		if overlap == 2 {
			count++
			continue
		}
		if overlap == 1 {
			count++
			continue
		}
	}
	return count
}

func Between(i, min, max int) bool {
	return min <= i && i <= max
}

func Overlap(ls, le, rs, re int) int {
	if Between(rs, ls, le) && Between(re, ls, le) {
		return 2 // fully contains
	}
	if Between(ls, rs, re) && Between(le, rs, re) {
		return 2 // fully contains
	}
	if Between(rs, ls, le) && re >= le {
		return 1 // left overlap
	}
	if Between(ls, rs, re) && le >= re {
		return 1 // right overlap
	}
	return 0 // no overlap
}

func parseInput(input string) [][]int {
	l := strings.Split(input, "\n")
	ans := make([][]int, len(l))
	for i, line := range l {
		parts := strings.Split(line, ",")
		ans[i] = make([]int, len(parts)*2)
		for j, part := range parts {
			rangeParts := strings.Split(part, "-")
			ans[i][j*2], _ = strconv.Atoi(rangeParts[0])
			ans[i][j*2+1], _ = strconv.Atoi(rangeParts[1])
		}
	}
	return ans
}
