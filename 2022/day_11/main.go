package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)

	if part == 1 {
		fmt.Println(part1(inputData))
	} else {
		fmt.Println(part2(inputData))
	}
}

func part1(data string) string {
	monkeys := []*Monkey{}

	for _, monkeyBlock := range strings.Split(data, "\n\n") {
		monkeys = append(monkeys, NewMonkey(strings.Split(monkeyBlock, "\n")))
	}

	monkeyParsedItemsCounter := make([]int, len(monkeys))

	for round := 0; round < 20; round++ {
		for monkeyIndex, monkey := range monkeys {
			// get all items to process
			items := []int{}
			items = append(items, monkey.Items...)
			monkey.Items = []int{}
			monkeyParsedItemsCounter[monkeyIndex] += len(items)
			// process items
			for _, item := range items {
				// operation
				item = monkey.ApplyOperation(item)
				// divide by 3
				item /= 3
				// throw item
				nextMonkeyIndex := monkey.TestNextMonkey(item)
				monkeys[nextMonkeyIndex].Items = append(monkeys[nextMonkeyIndex].Items, item)
			}
		}
	}

	sort.Ints(monkeyParsedItemsCounter)

	res := monkeyParsedItemsCounter[len(monkeyParsedItemsCounter)-1] * monkeyParsedItemsCounter[len(monkeyParsedItemsCounter)-2]

	return strconv.Itoa(res)
}

func part2(data string) string {
	monkeys := []*Monkey{}

	for _, monkeyBlock := range strings.Split(data, "\n\n") {
		monkeys = append(monkeys, NewMonkey(strings.Split(monkeyBlock, "\n")))
	}

	// all operations are periodical over
	// least common multiple
	leastCommonMultiple := 1
	for _, monkey := range monkeys {
		leastCommonMultiple *= monkey.Mod
	}

	monkeyParsedItemsCounter := make([]int, len(monkeys))

	for round := 0; round < 10000; round++ {
		for monkeyIndex, monkey := range monkeys {
			// get all items to process
			items := []int{}
			items = append(items, monkey.Items...)
			monkey.Items = []int{}
			monkeyParsedItemsCounter[monkeyIndex] += len(items)
			// process items
			for _, item := range items {
				// operation
				item = monkey.ApplyOperation(item)
				// Chinese Remainder Theorem
				// x % 3 = 2
				// x = 2, 5, 8, 11, 14, 17, 20, [23]
				// x % 5 = 3
				// x = 3, 8, 13, 18, [23]
				// x % 7 = 2
				// x = 2, 9, 16, [23]
				// lcm = 3*5*7 = 105
				// ===> x = 23 + k*105
				// 23 = (lcm*k + 23) % 3 = (lcm*k + 23) % 5 = (lcm*k + 23) % 7
				// modulo by lcm does not change modulo by prime divisors
				item %= leastCommonMultiple
				// throw item
				nextMonkeyIndex := monkey.TestNextMonkey(item)
				monkeys[nextMonkeyIndex].Items = append(monkeys[nextMonkeyIndex].Items, item)
			}
		}
	}

	sort.Ints(monkeyParsedItemsCounter)

	res := monkeyParsedItemsCounter[len(monkeyParsedItemsCounter)-1] * monkeyParsedItemsCounter[len(monkeyParsedItemsCounter)-2]

	return strconv.Itoa(res)
}

type Monkey struct {
	Items              []int
	CurrentItemIndexes []int
	ApplyOperation     func(old int) int
	Mod                int
	TestNextMonkey     func(old int) (nextMonkeyIndex int)
}

func NewMonkey(rawMonkey []string) (monkey *Monkey) {
	// Monkey x:
	monkey = &Monkey{}
	//starting items
	monkey.parseItems(rawMonkey[1])
	// operation
	monkey.parseOperation(rawMonkey[2])
	// test
	monkey.parseTest(rawMonkey[3:])

	return monkey
}

func (monkey *Monkey) parseItems(itemsRaw string) {
	// Starting items: 79, 98
	monkey.Items = []int{}
	nums := strings.Split(strings.Split(itemsRaw, ":")[1], ",")
	for _, numString := range nums {
		n, err := strconv.Atoi(strings.TrimSpace(numString))
		if err != nil {
			panic(err)
		}
		monkey.Items = append(monkey.Items, n)
	}
}

func (monkey *Monkey) parseOperation(operationRaw string) {
	// Operation: new = old * 19
	operands := strings.Split(strings.Split(operationRaw, "=")[1], " ")
	monkey.ApplyOperation = func(old int) int {
		inputStr := []string{operands[1], operands[3]}
		inputInt := []int{}
		for _, variable := range inputStr {
			if variable == "old" {
				inputInt = append(inputInt, old)
				continue
			}
			num, err := strconv.Atoi(strings.TrimSpace(variable))
			if err != nil {
				panic(err)
			}
			inputInt = append(inputInt, num)
		}

		operator := operands[2]

		if operator == "+" {
			return inputInt[0] + inputInt[1]
		}
		if operator == "*" {
			return inputInt[0] * inputInt[1]
		}

		panic(fmt.Errorf("Not allowed operation: %s", operator))
	}
}

func (monkey *Monkey) parseTest(testRaw []string) {
	// Test: divisible by 23
	//   If true: throw to monkey 2
	//   If false: throw to monkey 3
	divisibleByStr := strings.Split(testRaw[0], "by")[1]
	divisibleBy, err := strconv.Atoi(strings.TrimSpace(divisibleByStr))
	if err != nil {
		panic(err)
	}
	nextIfTrueStr := strings.Split(testRaw[1], "monkey")[1]
	nextIfTrue, err := strconv.Atoi(strings.TrimSpace(nextIfTrueStr))
	if err != nil {
		panic(err)
	}
	nextIfFalseStr := strings.Split(testRaw[2], "monkey")[1]
	nextIfFalse, err := strconv.Atoi(strings.TrimSpace(nextIfFalseStr))
	if err != nil {
		panic(err)
	}
	monkey.Mod = divisibleBy
	monkey.TestNextMonkey = func(old int) (nextMonkeyIndex int) {
		if old%monkey.Mod == 0 {
			return nextIfTrue
		}

		return nextIfFalse
	}
}
