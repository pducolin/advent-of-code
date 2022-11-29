package template

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed input.txt
var inputData string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		fmt.Println(inputData)
	} else {
		fmt.Println(part2(inputData))
	}
}

func part1(data string) string {
	return data
}

func part2(data string) string {
	return data
}
