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
	monkeys := parseInput(input)
	for i := 0; i < 20; i++ {
		monkeys.Round(func(item int) int {
			return item / 3
		})
	}
	return monkeyBusiness(monkeys)
}

func part2(input string) int {
	monkeys := parseInput(input)
	nums := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		nums[i] = monkey.test.value
	}
	relief := Multiply(nums...)
	for i := 0; i < 10000; i++ {
		monkeys.Round(func(item int) int {
			return item % relief
		})
	}
	return monkeyBusiness(monkeys)
}

type Monkey struct {
	id        int
	items     []int
	operation Operation
	test      Test
	activity  int
}

type Operation struct {
	left, operator, right string
}

type Test struct {
	value, true, false int
}

type Monkeys []*Monkey

func (m *Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) Inspect(item int) int {
	nums := make([]int, 2)
	if m.operation.left == "old" {
		nums[0] = item
	} else {
		tmp, _ := strconv.Atoi(m.operation.left)
		nums[0] = tmp
	}
	if m.operation.right == "old" {
		nums[1] = item
	} else {
		tmp, _ := strconv.Atoi(m.operation.right)
		nums[1] = tmp
	}
	switch m.operation.operator {
	case "+":
		return nums[0] + nums[1]
	case "*":
		return nums[0] * nums[1]
	}
	panic("unhandled operator")
}

func (m *Monkey) Test(item int) bool {
	return item%m.test.value == 0
}

func (m Monkeys) Round(relief func(int) int) {
	for _, monkey := range m {
		for len(monkey.items) > 0 {
			item := monkey.items[0]
			monkey.items = monkey.items[1:]
			item = monkey.Inspect(item)
			item = relief(item)
			if monkey.Test(item) {
				m[monkey.test.true].AddItem(item)
			} else {
				m[monkey.test.false].AddItem(item)
			}
			monkey.activity++
		}
	}
}

func monkeyBusiness(m Monkeys) int {
	var max1, max2 int
	for _, monkey := range m {
		if monkey.activity > max1 {
			max2 = max1
			max1 = monkey.activity
		} else if monkey.activity > max2 {
			max2 = monkey.activity
		}
	}
	return max1 * max2
}

type Numbers interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func Multiply[T Numbers](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	total := nums[0]
	for _, num := range nums[1:] {
		total *= num
	}
	return total
}

func parseInput(input string) Monkeys {
	blocks := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n\n")
	monkeys := make(Monkeys, len(blocks))
	for i, block := range blocks {
		monkey := &Monkey{id: i, items: make([]int, 0)}
		monkeys[i] = monkey
		lines := strings.Split(block, "\n")
		fields := strings.Fields(lines[1])
		for _, field := range fields[2:] {
			field = strings.Trim(field, ",")
			item, _ := strconv.Atoi(field)
			monkey.items = append(monkey.items, item)
		}
		fields = strings.Fields(lines[2])
		monkey.operation = Operation{left: fields[3], operator: fields[4], right: fields[5]}
		fields = strings.Fields(lines[3])
		tmp, _ := strconv.Atoi(fields[3])
		monkey.test = Test{
			value: tmp,
		}
		fields = strings.Fields(lines[4])
		trueLine, _ := strconv.Atoi(fields[len(fields)-1])
		monkey.test.true = trueLine
		fields = strings.Fields(lines[5])
		falseLine, _ := strconv.Atoi(fields[len(fields)-1])
		monkey.test.false = falseLine
	}
	return monkeys
}
