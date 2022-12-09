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
	h, t := Point{}, make([]Point, 1)
	tailLen := len(t)
	visited := make(map[Point]struct{})
	visited[t[tailLen-1]] = struct{}{}
	for _, line := range parsed {
		steps, _ := strconv.Atoi(line[1])
		for i := 0; i < steps; i++ {
			h = h.Add(getDirection(line[0]))
			t[0] = Move(h, t[0])
			visited[t[tailLen-1]] = struct{}{}
		}
	}
	return len(visited)
}

func part2(input string) int {
	parsed := parseInput(input)
	h, t := Point{}, make([]Point, 9)
	tailLen := len(t)
	visited := make(map[Point]struct{})
	visited[t[tailLen-1]] = struct{}{}
	for _, line := range parsed {
		steps, _ := strconv.Atoi(line[1])
		for i := 0; i < steps; i++ {
			h = h.Add(getDirection(line[0]))
			for i, tPart := range t {
				if i == 0 {
					t[i] = Move(h, tPart)
				} else {
					t[i] = Move(t[i-1], tPart)
				}
			}
			visited[t[tailLen-1]] = struct{}{}
		}
	}
	return len(visited)
}

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	p.x += other.x
	p.y += other.y
	return p
}

func getDirection(dir string) Point {
	switch dir {
	case "U":
		return Point{0, 1}
	case "D":
		return Point{0, -1}
	case "L":
		return Point{-1, 0}
	case "R":
		return Point{1, 0}
	}
	return Point{}
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

func Sign[T Numbers](x T) T {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func Move(h, t Point) Point {
	distY, distX := h.y-t.y, h.x-t.x
	if Abs(distY) >= 2 || Abs(distX) >= 2 {
		t.y += Sign(distY)
		t.x += Sign(distX)
	}
	return t
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
