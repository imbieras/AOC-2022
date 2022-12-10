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
		fmt.Print("Output:\n", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	cycles := []int{20, 60, 100, 140, 180, 220}
	x, t, ans := 1, 0, 0
	for _, line := range parsed {
		switch line[0] {
		case "noop":
			t++
			ans = Check(t, x, cycles, ans)
		case "addx":
			for i := 0; i < 2; i++ {
				t++
				ans = Check(t, x, cycles, ans)
			}
			value, _ := strconv.Atoi(line[1])
			x += value
		}
	}
	return ans
}

func part2(input string) string {
	parsed := parseInput(input)
	x, t := 1, 0
	grid := make([][]string, 6)
	for i := 0; i < 6; i++ {
		grid[i] = make([]string, 40)
	}
	for _, line := range parsed {
		switch line[0] {
		case "noop":
			t++
			PutSymbol(t, x, grid)
		case "addx":
			for i := 0; i < 2; i++ {
				t++
				PutSymbol(t, x, grid)
			}
			value, _ := strconv.Atoi(line[1])
			x += value
		}
	}
	ans := ""
	for _, row := range grid {
		ans += strings.Join(row, "") + "\n"
	}
	return ans
}

func Check(t int, x int, cycles []int, ans int) int {
	for _, cycle := range cycles {
		if t == cycle {
			ans += t * x
		}
	}
	return ans
}

func PutSymbol(t int, x int, grid [][]string) {
	t--
	if Abs(x-(t%40)) <= 1 {
		grid[t/40][t%40] = "\xe2\xac\x9c"
	} else {
		grid[t/40][t%40] = "\xe2\xac\x9b"
	}
}

type Numbers interface {
	float32 | float64 | int | int8 | int16 | int32 | int64
}

func Abs[T Numbers](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	ans := make([][]string, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		ans = append(ans, parts)
	}
	return ans
}
