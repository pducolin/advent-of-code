package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	assert.Equal(t, "21", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "8", part2(data), "Failed testing part 2")
}
