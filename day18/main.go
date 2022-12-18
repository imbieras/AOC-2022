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
	cubes := parseInput(input)
	ans := 0
	for _, cube := range cubes {
		for _, dx := range []int{-1, 1} {
			if !checkIfInSlice(Cube{cube.x + dx, cube.y, cube.z}, cubes) {
				ans++
			}
		}
		for _, dy := range []int{-1, 1} {
			if !checkIfInSlice(Cube{cube.x, cube.y + dy, cube.z}, cubes) {
				ans++
			}
		}
		for _, dz := range []int{-1, 1} {
			if !checkIfInSlice(Cube{cube.x, cube.y, cube.z + dz}, cubes) {
				ans++
			}
		}
	}
	return ans
}

func part2(input string) int {
	cubes := parseInput(input)
	ans := 0
	out := make(map[[3]int]bool)
	in := make(map[[3]int]bool)
	for _, cube := range cubes {
		for _, dx := range []int{-1, 1} {
			if reachesOutside(Cube{cube.x + dx, cube.y, cube.z}, cubes, in, out) {
				ans++
			}
		}
		for _, dy := range []int{-1, 1} {
			if reachesOutside(Cube{cube.x, cube.y + dy, cube.z}, cubes, in, out) {
				ans++
			}
		}
		for _, dz := range []int{-1, 1} {
			if reachesOutside(Cube{cube.x, cube.y, cube.z + dz}, cubes, in, out) {
				ans++
			}
		}
	}
	return ans
}

type Cube struct {
	x, y, z int
}

func checkIfInSlice(cube Cube, cubes []Cube) bool {
	for _, p := range cubes {
		if p == cube {
			return true
		}
	}
	return false
}

func checkIfInMap(cube Cube, cubes map[[3]int]bool) bool {
	return cubes[[3]int{cube.x, cube.y, cube.z}]
}

func reachesOutside(cube Cube, cubes []Cube, in, out map[[3]int]bool) bool {
	if checkIfInMap(cube, out) {
		return true
	}
	if checkIfInMap(cube, in) {
		return false
	}
	visited := make(map[[3]int]bool)
	queue := [][3]int{{cube.x, cube.y, cube.z}}
	for len(queue) > 0 {
		x, y, z := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if checkIfInSlice(Cube{x, y, z}, cubes) {
			continue
		}
		if visited[[3]int{x, y, z}] {
			continue
		}
		visited[[3]int{x, y, z}] = true
		if len(visited) > 1500 {
			for cube := range visited {
				out[cube] = true
			}
			return true
		}
		for _, dx := range []int{-1, 1} {
			queue = append(queue, [3]int{x + dx, y, z})
		}
		for _, dy := range []int{-1, 1} {
			queue = append(queue, [3]int{x, y + dy, z})
		}
		for _, dz := range []int{-1, 1} {
			queue = append(queue, [3]int{x, y, z + dz})
		}
	}
	for cube := range visited {
		in[cube] = true
	}
	return false
}

func parseInput(input string) []Cube {
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	ans := make([]Cube, 0)
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		ans = append(ans, Cube{x, y, z})
	}
	return ans
}
