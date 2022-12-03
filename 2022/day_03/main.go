package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	res := 0
	for _, rack := range strings.Split(data, "\n") {
		firstRack := map[byte]struct{}{}
		secondRack := map[byte]struct{}{}

		halfRackIndex := len(rack) / 2
		for i := 0; i < halfRackIndex; i++ {
			a := rack[i]
			b := rack[i+halfRackIndex]
			firstRack[a] = struct{}{}
			secondRack[b] = struct{}{}
		}

		for k := range intersection(firstRack, secondRack) {
			res += evaluateItem(k)
		}

	}
	return fmt.Sprint(res)
}

func part2(data string) string {
	res := 0
	lineIndex := 0
	lines := strings.Split(data, "\n")
	for lineIndex < len(lines) {
		groupRacks := []map[byte]struct{}{}
		for i := 0; i < 3; i++ {
			rack := parseRackLine(lines[lineIndex+i])
			groupRacks = append(groupRacks, rack)
		}
		// intersection
		commonItems := groupRacks[0]

		for i := 1; i < 3; i++ {
			commonItems = intersection(commonItems, groupRacks[i])
		}

		for k := range commonItems {
			res += evaluateItem(k)
		}

		lineIndex += 3
	}

	return fmt.Sprint(res)
}

func evaluateItem(item byte) int {
	if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	}

	return int(item-'a') + 1
}

func intersection(setA, setB map[byte]struct{}) (intersection map[byte]struct{}) {
	intersection = map[byte]struct{}{}
	for k := range setA {
		if _, ok := setB[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func parseRackLine(line string) (rack map[byte]struct{}) {
	rack = map[byte]struct{}{}
	for i := 0; i < len(line); i++ {
		a := line[i]
		rack[a] = struct{}{}
	}
	return rack
}
