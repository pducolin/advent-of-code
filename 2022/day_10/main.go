package main

import (
	_ "embed"
	"flag"
	"fmt"
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

// 20th, 60th, 100th, 140th, 180th, and 220th
const (
	firstCycleToCheck = 20
	cycleInterval     = 40
	lastCycleToCheck  = 220
)

func part1(data string) string {
	instructions := strings.Split(data, "\n")
	x := 1
	res := 0
	currentInstructionIndex := 0
	nextCycleToCheck := firstCycleToCheck
	toAdd := []int{}

	for cycleIndex := 1; cycleIndex <= lastCycleToCheck; cycleIndex++ {
		if cycleIndex == nextCycleToCheck {
			res += x * cycleIndex
			nextCycleToCheck += cycleInterval
			if nextCycleToCheck > lastCycleToCheck {
				break
			}
		}

		if len(toAdd) > 0 {
			x += toAdd[0]
			toAdd = []int{}
			currentInstructionIndex++
			continue
		}

		if instructions[currentInstructionIndex] == "noop" {
			currentInstructionIndex++
			continue
		}

		increment, err := strconv.Atoi(strings.Split(instructions[currentInstructionIndex], " ")[1])
		if err != nil {
			panic(err)
		}
		toAdd = append(toAdd, increment)
	}

	return strconv.Itoa(res)
}

const lineLength = 40

func part2(data string) string {
	instructions := strings.Split(data, "\n")
	spriteMiddleX := 1
	res := ""
	currentInstructionIndex := 0
	toAdd := []int{}
	currentXIndex := 0

	for currentInstructionIndex < len(instructions) {
		// draw pixel
		if currentXIndex >= spriteMiddleX-1 && currentXIndex <= spriteMiddleX+1 {
			res += "#"
		} else {
			res += "."
		}
		currentXIndex++
		if currentXIndex == lineLength {
			currentXIndex = 0
			res += "\n"
		}
		// ##..##..##..##..##..##..##..##..##..##...
		// ##..##..##..##..##..##..##..##..##..##..

		// move sprite
		if len(toAdd) > 0 {
			spriteMiddleX += toAdd[0]
			toAdd = []int{}
			currentInstructionIndex++
			continue
		}

		if instructions[currentInstructionIndex] == "noop" {
			currentInstructionIndex++
			continue
		}

		increment, err := strconv.Atoi(strings.Split(instructions[currentInstructionIndex], " ")[1])
		if err != nil {
			panic(err)
		}
		toAdd = append(toAdd, increment)
	}

	return res
}
