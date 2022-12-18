package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestPart1(t *testing.T) {
	assert.Equal(t, "3068", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "1514285714288", part2(data), "Failed testing part 2")
}
