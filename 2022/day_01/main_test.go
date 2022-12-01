package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	assert.Equal(t, part1(data), "24000", "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, part2(data), "45000", "Failed testing part 2")
}
