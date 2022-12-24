package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `1
2
-3
3
-2
0
4`

func TestPart1(t *testing.T) {
	assert.Equal(t, "3", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "1623178306", part2(data), "Failed testing part 2")
}
