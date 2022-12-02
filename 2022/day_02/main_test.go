package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	assert.Equal(t, "15", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "12", part2(data), "Failed testing part 2")
}
