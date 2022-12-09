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
	grid := parseInput(input)
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			isVisible := false
			for _, dir := range directions {
				tmpRow, tmpCol := row, col
				ok := true
				for {
					tmpRow, tmpCol = tmpRow+dir[0], tmpCol+dir[1]
					if !(0 <= tmpRow && tmpRow < rows && 0 <= tmpCol && tmpCol < cols) {
						break
					}
					if grid[tmpRow][tmpCol] >= grid[row][col] {
						ok = false
					}
				}
				if ok {
					isVisible = true
				}
			}
			if isVisible {
				count++
			}
		}
	}
	return count
}

func part2(input string) int {
	grid := parseInput(input)
	rows := len(grid)
	cols := len(grid[0])
	maxScore := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			score := 1
			for _, dir := range directions {
				dist := 1
				tmpRow, tmpCol := row+dir[0], col+dir[1]
				for {
					if !(0 <= tmpRow && tmpRow < rows && 0 <= tmpCol && tmpCol < cols) {
						dist -= 1
						break
					}
					if grid[tmpRow][tmpCol] >= grid[row][col] {
						break
					}
					dist++
					tmpRow, tmpCol = tmpRow+dir[0], tmpCol+dir[1]
				}
				score *= dist
			}
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	ans := make([][]int, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		chars := strings.Split(line, "")
		for _, char := range chars {
			intVal, _ := strconv.Atoi(char)
			ans[i] = append(ans[i], intVal)
		}
	}
	return ans
}
