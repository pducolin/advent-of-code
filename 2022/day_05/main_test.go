package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1(t *testing.T) {
	assert.Equal(t, "CMZ", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "MCD", part2(data), "Failed testing part 2")
}
