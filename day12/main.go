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
	grid := parseInput(input)

	queue := [][3]int{}
	visited := map[[2]int]bool{}
	er, ec := 0, 0

	for r, rows := range grid {
		for c, cell := range rows {
			if cell == "S" {
				queue = append(queue, [3]int{0, r, c})
				visited[[2]int{r, c}] = true
				grid[r][c] = "a"
			}
			if cell == "E" {
				er, ec = r, c
				grid[r][c] = "z"
			}
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newR, newC := front[1]+dir[0], front[2]+dir[1]
			if newR < 0 || newC < 0 || newR >= len(grid) || newC >= len(grid[0]) {
				continue
			}
			if visited[[2]int{newR, newC}] {
				continue
			}
			if Distance(grid[front[1]][front[2]], grid[newR][newC]) > 1 {
				continue
			}
			if newR == er && newC == ec {
				return front[0] + 1
			}
			visited[[2]int{newR, newC}] = true
			queue = append(queue, [3]int{front[0] + 1, newR, newC})
		}
	}
	panic("no path found")
}

func part2(input string) int {
	grid := parseInput(input)

	queue := [][3]int{}
	visited := map[[2]int]bool{}

	for r, rows := range grid {
		for c, cell := range rows {
			if cell == "S" {
				grid[r][c] = "a"
			}
			if cell == "E" {
				queue = append(queue, [3]int{0, r, c})
				visited[[2]int{r, c}] = true
				grid[r][c] = "z"
			}
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newR, newC := front[1]+dir[0], front[2]+dir[1]
			if newR < 0 || newC < 0 || newR >= len(grid) || newC >= len(grid[0]) {
				continue
			}
			if visited[[2]int{newR, newC}] {
				continue
			}
			if Distance(grid[front[1]][front[2]], grid[newR][newC]) < -1 {
				continue
			}
			if grid[newR][newC] == "a" {
				return front[0] + 1
			}
			visited[[2]int{newR, newC}] = true
			queue = append(queue, [3]int{front[0] + 1, newR, newC})
		}
	}
	panic("no path found")
}

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func Distance(x, y string) int {
	xRune := []rune(x)[0]
	yRune := []rune(y)[0]
	return int(yRune - xRune)
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
