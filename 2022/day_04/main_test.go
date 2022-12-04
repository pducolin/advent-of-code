package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1(t *testing.T) {
	assert.Equal(t, "2", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "4", part2(data), "Failed testing part 2")
}
