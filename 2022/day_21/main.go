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
	program := NewProgram(data)
	res, err := program.Solve("root", "")
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(res)
}

const LIMIT = 1000000000000000

func part2(data string) string {
	program := NewProgram(data)
	parts := strings.Split(program.instructions["root"].operation, " ")
	a := strings.TrimSpace(parts[0])
	b := strings.TrimSpace(parts[2])

	results := map[string]int{}
	resA, err := program.Solve(a, "humn")
	if err != nil {
		fmt.Println("res a", err.Error())
	} else {
		fmt.Println("res a", resA)
		results[a] = resA
	}
	resB, err := program.Solve(b, "humn")
	if err != nil {
		fmt.Println("res b", err.Error())
	} else {
		fmt.Println("res b", resB)
		results[b] = resB
	}

	program.Reset()
	program.results["humn"] = 3247317268284
	for id, value := range results {
		program.results[id] = value
	}
	resA, err = program.Solve(a, "")
	if err != nil {
		panic(err)
	}
	resB, err = program.Solve(b, "")
	if err != nil {
		panic(err)
	}
	if resA != resB {
		panic("Wrong answer")
	}
	return "FAIL!"
}

type Instruction struct {
	id        string
	operation string
}

type Program struct {
	instructions map[string]Instruction
	results      map[string]int
}

func NewProgram(data string) Program {
	program := Program{
		instructions: map[string]Instruction{},
		results:      map[string]int{},
	}

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, ":")
		id := parts[0]
		operation := strings.TrimSpace(parts[1])
		program.instructions[id] = Instruction{id, operation}
		if num, err := strconv.Atoi(operation); err == nil {
			program.results[id] = num
		}
	}
	return program
}

var operationRe = regexp.MustCompile(`(\w{4}) ([\+\-\*/]){1} (\w{4})`)

func (p *Program) Solve(id string, breakId string) (int, error) {
	if id == breakId {
		return -1, errors.New(id)
	}
	if res, found := p.results[id]; found {
		return res, nil
	}
	groups := operationRe.FindStringSubmatch(p.instructions[id].operation)

	op := groups[2]

	a, errA := p.Solve(groups[1], breakId)
	b, errB := p.Solve(groups[3], breakId)

	if errA != nil || errB != nil {
		errMsg := "("
		if errA != nil {
			errMsg += errA.Error()
		} else {
			errMsg += strconv.Itoa(a)
		}

		errMsg += op
		if errB != nil {
			errMsg += errB.Error()
		} else {
			errMsg += strconv.Itoa(b)
		}
		errMsg += ")"

		return -1, errors.New(errMsg)
	}

	if op == "+" {
		return a + b, nil
	}
	if op == "-" {
		return a - b, nil
	}
	if op == "*" {
		return a * b, nil
	}
	return a / b, nil
}

func (p *Program) Reset() {
	p.results = map[string]int{}
	for id, operation := range p.instructions {
		if num, err := strconv.Atoi(operation.operation); err == nil {
			p.results[id] = num
		}
	}
}

func (p *Program) SolveHasHuman(id string) bool {
	if id == "humn" {
		return true
	}
	if _, found := p.results[id]; found {
		return false
	}
	groups := operationRe.FindStringSubmatch(p.instructions[id].operation)

	return p.SolveHasHuman(groups[1]) || p.SolveHasHuman(groups[3])
}
