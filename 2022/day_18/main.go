package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"

	"github.com/pducolin/advent-of-code/2022/day_18/lava"
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
	grid := lava.NewGrid(data)
	res := 0
	for cube := range grid.LavaCubes {
		res += grid.CountFreeSides(cube)
	}
	return strconv.Itoa(res)
}

func part2(data string) string {
	grid := lava.NewGrid(data)
	grid.FloodFill()
	res := 0
	for cube := range grid.LavaCubes {
		res += grid.CountReachableSides(cube)
	}
	return strconv.Itoa(res)
}
