package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
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
	lines := parseInput(input)
	points := make(map[Point]struct{})
	for _, corners := range lines {
		prev := Point{0, 0}
		for _, corner := range corners {
			if prev.x != 0 && prev.y != 0 {
				dx := corner.x - prev.x
				dy := corner.y - prev.y
				len := Max(Abs(dx), Abs(dy))
				for i := 0; i <= len; i++ {
					points[Point{prev.x + i*dx/len, prev.y + i*dy/len}] = struct{}{}
				}
			}
			prev = corner
		}
	}

	for i := 0; i < 1e7; i++ {
		sand := Point{500, 0}
		for {
			if sand.y > 300 {
				return i
			} else if !checkIfIn(Point{sand.x, sand.y + 1}, points) {
				sand.y++
			} else if !checkIfIn(Point{sand.x - 1, sand.y + 1}, points) {
				sand.x--
				sand.y++
			} else if !checkIfIn(Point{sand.x + 1, sand.y + 1}, points) {
				sand.x++
				sand.y++
			} else {
				break
			}
		}
		points[sand] = struct{}{}
	}
	return 0
}

func part2(input string) int {
	lines := parseInput(input)
	points := make(map[Point]struct{})
	for _, corners := range lines {
		prev := Point{0, 0}
		for _, corner := range corners {
			if prev.x != 0 && prev.y != 0 {
				dx := corner.x - prev.x
				dy := corner.y - prev.y
				len := Max(Abs(dx), Abs(dy))
				for i := 0; i <= len; i++ {
					points[Point{prev.x + i*dx/len, prev.y + i*dy/len}] = struct{}{}
				}
			}
			prev = corner
		}
	}

	floorY := maxPoint(points).y + 2

	for i, end := 0, 500+floorY+1; i < end; i++ { // ;P
		points[Point{i, floorY}] = struct{}{}
	}

	for i := 0; i < 1e7; i++ {
		sand := Point{500, 0}
		for {
			if !checkIfIn(Point{sand.x, sand.y + 1}, points) {
				sand.y++
			} else if !checkIfIn(Point{sand.x - 1, sand.y + 1}, points) {
				sand.x--
				sand.y++
			} else if !checkIfIn(Point{sand.x + 1, sand.y + 1}, points) {
				sand.x++
				sand.y++
			} else {
				break
			}
		}
		if (sand == Point{500, 0}) {
			return i + 1
		}
		points[sand] = struct{}{}
	}
	return 0
}

type Point struct {
	x, y int
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

func Max[T Numbers](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func maxPoint(points map[Point]struct{}) Point {
	max := Point{0, 0}
	for point := range points {
		if point.x > max.x {
			max.x = point.x
		}
		if point.y > max.y {
			max.y = point.y
		}
	}
	return max
}

func checkIfIn(point Point, points map[Point]struct{}) bool {
	_, ok := points[point]
	return ok
}

func parseInput(input string) [][]Point {
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	ans := make([][]Point, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, " -> ")
		for _, coord := range coords {
			coord := strings.Split(coord, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			ans[i] = append(ans[i], Point{x, y})
		}
	}
	return ans
}
