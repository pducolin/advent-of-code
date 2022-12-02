package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = "Hello World !"

func TestPart1(t *testing.T) {
	assert.Equal(t, "Hello World !", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "Hello World !", part2(data), "Failed testing part 2")
}
