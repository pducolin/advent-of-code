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
	caloriesPerElf, err := parseSortedCaloriesPerElf(data)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(caloriesPerElf[len(caloriesPerElf)-1])
}

func part2(data string) string {
	caloriesPerElf, err := parseSortedCaloriesPerElf(data)
	if err != nil {
		panic(err)
	}
	res := 0
	for i := 1; i <= 3; i++ {
		res += caloriesPerElf[len(caloriesPerElf)-i]
	}
	return fmt.Sprint(res)
}

func parseSortedCaloriesPerElf(data string) (caloriesPerElf []int, err error) {
	// parse calories
	currentElfCalories := 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			caloriesPerElf = append(caloriesPerElf, currentElfCalories)
			currentElfCalories = 0
			continue
		}
		c, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		currentElfCalories += c
	}
	caloriesPerElf = append(caloriesPerElf, currentElfCalories)

	sort.Ints(caloriesPerElf)
	return caloriesPerElf, nil
}
