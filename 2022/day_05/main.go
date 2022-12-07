package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"regexp"
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
	stackCount, stackLines, instructions := parseInput(data)

	stacks := parseStacks(stackCount, stackLines)

	applyInstructions9000(stacks, instructions)

	res := ""
	for _, stack := range stacks {
		top, err := stack.Pop()
		if err != nil {
			continue
		}
		res += top

	}

	return res
}

func part2(data string) string {
	stackCount, stackLines, instructions := parseInput(data)

	stacks := parseStacks(stackCount, stackLines)

	applyInstructions9001(stacks, instructions)

	res := ""
	for _, stack := range stacks {
		top, err := stack.Pop()
		if err != nil {
			continue
		}
		res += top
	}

	return res
}

func parseInput(data string) (stackCount int, stackLines, instructions []string) {
	stackLines = []string{}
	stackCount = 0

	lines := strings.Split(data, "\n")
	// parse stackLines
	for _, line := range lines {
		if strings.HasPrefix(line, " 1") {
			cols := strings.Split(line, "   ")
			stackCount = len(cols)
			break
		}
		stackLines = append(stackLines, line)
	}
	instructions = lines[len(stackLines)+2:]

	return stackCount, stackLines, instructions
}

func parseStacks(stackCount int, lines []string) (stacks []*Stack) {
	stacks = make([]*Stack, stackCount)

	for i := range stacks {
		stacks[i] = NewStack()
	}

	for lineIndex := len(lines) - 1; lineIndex >= 0; lineIndex-- {
		line := lines[lineIndex]
		i := 0
		for colIndex := 0; colIndex < stackCount; colIndex++ {
			item := line[i : i+3]
			if item[0] == '[' {
				stacks[colIndex].Push(string(item[1]))
			}
			i += 4
		}
	}

	return stacks
}

func applyInstructions9000(stacks []*Stack, instructions []string) {
	r := regexp.MustCompile(`move\ (?P<count>\d+)\ from\ (?P<from>\d+)\ to\ (?P<to>\d+)`)
	for _, inst := range instructions {
		m := r.FindStringSubmatch(inst)[1:]
		count, err := strconv.Atoi(m[0])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		from -= 1
		to, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		to -= 1

		for i := 0; i < count; i++ {
			node, err := stacks[from].Pop()
			if err != nil {
				panic(err)
			}
			stacks[to].Push(node)
		}
	}
}

func applyInstructions9001(stacks []*Stack, instructions []string) {
	for _, inst := range instructions {
		r := regexp.MustCompile(`move\ (?P<count>\d+)\ from\ (?P<from>\d+)\ to\ (?P<to>\d+)`)
		m := r.FindStringSubmatch(inst)[1:]
		count, err := strconv.Atoi(m[0])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		from -= 1
		to, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		to -= 1

		tmp := []string{}
		for i := 0; i < count; i++ {
			node, err := stacks[from].Pop()
			if err != nil {
				panic(err)
			}
			tmp = append(tmp, node)
		}
		for i := range tmp {
			stacks[to].Push(tmp[len(tmp)-i-1])
		}
	}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes string
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(item string) {
	s.nodes += item
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() (item string, err error) {
	if len(s.nodes) == 0 {
		return item, errors.New("empty stack")
	}
	item = string(s.nodes[len(s.nodes)-1])
	s.nodes = s.nodes[:len(s.nodes)-1]
	return item, nil
}

// Peek returns a node from the stack in last to first order.
func (s *Stack) Top() (item string, err error) {
	if len(s.nodes) == 0 {
		return item, errors.New("empty stack")
	}
	item = string(s.nodes[len(s.nodes)-1])
	return item, nil
}
