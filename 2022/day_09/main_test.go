package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestPart1(t *testing.T) {
	assert.Equal(t, "13", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "1", part2(data), "Failed testing part 2")
}
