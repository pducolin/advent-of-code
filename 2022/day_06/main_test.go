package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	data            string
	expectedResult1 int
	expectedResult2 int
}

var dataSlice = []testData{
	{data: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expectedResult1: 7, expectedResult2: 19},
	{data: "bvwbjplbgvbhsrlpgdmjqwftvncz", expectedResult1: 5, expectedResult2: 23},
	{data: "nppdvjthqldpwncqszvftbrmjlhg", expectedResult1: 6, expectedResult2: 23},
	{data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expectedResult1: 10, expectedResult2: 29},
	{data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expectedResult1: 11, expectedResult2: 26},
}

func TestPart1(t *testing.T) {
	for _, d := range dataSlice {
		assert.Equal(t, strconv.Itoa(d.expectedResult1), part1(d.data), "Failed testing part 1")
	}
}

func TestPart2(t *testing.T) {
	for _, d := range dataSlice {
		assert.Equal(t, strconv.Itoa(d.expectedResult2), part2(d.data), "Failed testing part 2")
	}
}
