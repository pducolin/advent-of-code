package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = "Hello World !"

func TestPart1(t *testing.T) {
	assert.Equal(t, part1(data), "Hello World !", "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, part2(data), "Hello World !", "Failed testing part 2")
}
