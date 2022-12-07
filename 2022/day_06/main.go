package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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
	i := 4
	for i <= len(data) {
		frame := data[i-4 : i]
		if countUniqueLetters(frame) == len(frame) {
			break
		}
		i++
	}
	return strconv.Itoa(i)
}

func part2(data string) string {
	i := 14
	for i <= len(data) {
		frame := data[i-14 : i]
		if countUniqueLetters(frame) == len(frame) {
			break
		}
		i++
	}
	return strconv.Itoa(i)
}

func countUniqueLetters(frame string) int {
	uniqueLetters := map[rune]struct{}{}
	res := 0
	for _, letter := range frame {
		if _, found := uniqueLetters[letter]; found {
			continue
		}
		uniqueLetters[letter] = struct{}{}
		res++
	}
	return res
}
