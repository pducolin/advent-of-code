package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

	fmt.Println(part1(inputData))
}

func part1(data string) string {
	sum := 0

	for _, line := range strings.Split(data, "\n") {
		sum += SnafuToInt(line)
	}

	return IntToSnafu(sum)
}

func IntToSnafu(number int) string {
	snafuChars := []string{}
	for number > 0 {
		mod := number % 5
		if mod <= 2 {
			snafuChars = append(snafuChars, strconv.Itoa(mod))
			number -= mod
		} else if mod == 3 {
			snafuChars = append(snafuChars, "=")
			number += 5 - mod
		} else {
			snafuChars = append(snafuChars, "-")
			number += 5 - mod
		}
		number /= 5
	}
	return invertString(strings.Join(snafuChars, ""))
}

func SnafuToInt(snafu string) int {
	num := 0

	for i, r := range invertString(snafu) {
		switch r {
		case '0':
			num += int(r-'0') * int(math.Pow(5, float64(i)))
		case '1':
			num += int(r-'0') * int(math.Pow(5, float64(i)))
		case '2':
			num += int(r-'0') * int(math.Pow(5, float64(i)))
		case '-':
			num -= int(math.Pow(5, float64(i)))
		case '=':
			num -= 2 * int(math.Pow(5, float64(i)))
		}
	}

	return num
}

func invertString(s string) string {
	ret := ""
	for i := len(s) - 1; i >= 0; i-- {
		ret += string(s[i])
	}
	return ret
}
