package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestPart1(t *testing.T) {
	assert.Equal(t, "64", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "58", part2(data), "Failed testing part 2")
}
