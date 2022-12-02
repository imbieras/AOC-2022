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
		me := string(l[2])
		opp := string(l[0])
		switch opp {
		case "A":
			switch me {
			case "X":
				sum += 4 //3 + 1
			case "Y":
				sum += 8 //6 + 2
			case "Z":
				sum += 3 //3 + 0
			}
		case "B":
			switch me {
			case "X":
				sum += 1 //1 + 0
			case "Y":
				sum += 5 //3 + 2
			case "Z":
				sum += 9 //6 + 3
			}
		case "C":
			switch me {
			case "X":
				sum += 7 //6 + 1
			case "Y":
				sum += 2 //2 + 0
			case "Z":
				sum += 6 //3 + 3
			}
		}
	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	sum := 0
	for _, l := range parsed {
		me := string(l[2])
		opp := string(l[0])
		switch opp {
		case "A":
			switch me {
			case "X":
				sum += 3 //0 + 3
			case "Y":
				sum += 4 //3 + 1
			case "Z":
				sum += 8 //6 + 2
			}
		case "B":
			switch me {
			case "X":
				sum += 1 //0 + 1
			case "Y":
				sum += 5 //3 + 2
			case "Z":
				sum += 9 //6 + 3
			}
		case "C":
			switch me {
			case "X":
				sum += 2 //2 + 0
			case "Y":
				sum += 6 //3 + 3
			case "Z":
				sum += 7 //6 + 1
			}
		}
	}
	return sum
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
