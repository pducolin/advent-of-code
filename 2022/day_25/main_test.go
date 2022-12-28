package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

const intToSnafuData = `1 1
2 2
3 1=
4 1-
5 10
6 11
7 12
8 2=
9 2-
10 20
15 1=0
20 1-0
2022 1=11-2
12345 1-0---0
314159265 1121-1110-1=0`

func TestIntToSnafu(t *testing.T) {
	lines := strings.Split(intToSnafuData, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		snafu := parts[1]
		assert.Equal(t, snafu, IntToSnafu(num), fmt.Sprintf("Failed converting %d to snafu", num))
	}
}

func TestSnafuToInt(t *testing.T) {
	lines := strings.Split(intToSnafuData, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		snafu := parts[1]
		assert.Equal(t, num, SnafuToInt(snafu), fmt.Sprintf("Failed converting %s to num", snafu))
	}
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "2=-1=0", part1(data), "Failed testing part 1")
}
