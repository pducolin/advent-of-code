package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestPart1(t *testing.T) {
	assert.Equal(t, "24", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "93", part2(data), "Failed testing part 2")
}
