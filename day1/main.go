package main

import (
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i] > parsed[j]
	})
	return parsed[0]
}

func part2(input string) int {
	parsed := parseInput(input)
	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i] > parsed[j]
	})
	return sum(parsed[:3])
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func parseInput(input string) (ans []int) {
	temp := make([]int, 0)
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			ans = append(ans, sum(temp))
			temp = nil
		} else {
			if s, err := strconv.Atoi(l); err == nil {
				temp = append(temp, s)
			}
		}
	}
	return ans
}
