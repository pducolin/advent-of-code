package main

import (
	"strings"
	"testing"

	"github.com/pducolin/advent-of-code/2022/common"
	"github.com/stretchr/testify/assert"
)

const data = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

var expectedData = []string{
	`#.######
#.>3.<.#
#<..<<.#
#>2.22.#
#>v..^<#
######.#`,
	`#.######
#.2>2..#
#.^22^<#
#.>2.^>#
#.>..<.#
######.#`,
	`#.######
#<^<22.#
#.2<.2.#
#><2>..#
#..><..#
######.#`,
}

func TestIterate(t *testing.T) {
	grid := NewGrid(data)
	for _, expectedData := range expectedData {
		grid = grid.Iterate()
		assert.Equal(t, strings.Split(expectedData, "\n"), grid.ToString())
	}
}

func TestIsValidPosition(t *testing.T) {
	grid := NewGrid(data)
	grid = grid.Iterate()
	assert.True(t, grid.isValidPlayerPosition(common.Point{X: 1, Y: 1}), "failed isValidPlayerPosition")
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "18", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "54", part2(data), "Failed testing part 2")
}
