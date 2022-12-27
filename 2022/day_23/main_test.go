package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const smallerData = `.....
..##.
..#..
.....
..##.
.....`

const expectedFinalSmaller = `..#..
....#
#....
....#
.....
..#..`

const data = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func TestIterate(t *testing.T) {
	grid := NewGrid(smallerData)
	for i := 0; i < 3; i++ {
		grid.Iterate()
	}
	expectedLines := strings.Split(expectedFinalSmaller, "\n")
	lines := grid.ToString()
	assert.Equal(t, expectedLines, lines, "Failed iterate")
	// check it doesn't change anymore
	grid.Iterate()
	assert.Equal(t, lines, grid.ToString(), "Failed iterate")
}

const round1 = `.....#...
...#...#.
.#..#.#..
.....#..#
..#.#.##.
#..#.#...
#.#.#.##.
.........
..#..#...`

func TestPart1(t *testing.T) {
	grid := NewGrid(data)
	initialElves := len(grid.elves)
	assert.Equal(t, 22, initialElves, "Failed parsing grid")
	assert.Equal(t, initialElves, len(grid.elves), "Elves count changed after iteration")
	expectedLines := strings.Split(round1, "\n")
	grid.Iterate()
	assert.Equal(t, expectedLines, grid.ToString(), "Failed iteration 1")
	assert.Equal(t, "110", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "20", part2(data), "Failed testing part 2")
}
