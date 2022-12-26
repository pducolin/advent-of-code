package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/pducolin/advent-of-code/2022/common"
	"github.com/stretchr/testify/assert"
)

//go:embed input.test.txt
var data string

//go:embed input.txt
var realData string

func TestPart1(t *testing.T) {
	assert.Equal(t, "6032", part1(data, 4), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "5031", part2(data, 4, false), "Failed testing part 2")
}

func TestGridNextPointAroundCube(t *testing.T) {
	data = strings.Split(realData, "\n\n")[0]
	grid := NewGrid(data, 50)

	// right edge face 1, move right
	point, direction := grid.NextPointAroundCube(common.Point{X: 99, Y: 0}, RIGHT)
	assert.Equal(t, common.Point{X: 100, Y: 0}, point, "Failed testing face 1 right")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// left edge face 1, move left
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 0}, LEFT)
	assert.Equal(t, common.Point{X: 0, Y: 149}, point, "Failed testing face 1 left")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// top edge face 1, move up
	point, direction = grid.NextPointAroundCube(common.Point{X: 53, Y: 0}, UP)
	assert.Equal(t, common.Point{X: 0, Y: 153}, point, "Failed testing face 1 up")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// bottom edge face 1, move down
	point, direction = grid.NextPointAroundCube(common.Point{X: 54, Y: 49}, DOWN)
	assert.Equal(t, common.Point{X: 54, Y: 50}, point, "Failed testing face 1 down")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// right edge face 2, move right
	point, direction = grid.NextPointAroundCube(common.Point{X: 149, Y: 0}, RIGHT)
	assert.Equal(t, common.Point{X: 99, Y: 149}, point, "Failed testing face 2 right")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// left edge face 2, move left
	point, direction = grid.NextPointAroundCube(common.Point{X: 100, Y: 4}, LEFT)
	assert.Equal(t, common.Point{X: 99, Y: 4}, point, "Failed testing 2L")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// top edge face 2, move up
	point, direction = grid.NextPointAroundCube(common.Point{X: 105, Y: 0}, UP)
	assert.Equal(t, common.Point{X: 5, Y: 199}, point, "Failed testing 2U")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// bottom edge face 2, move down
	point, direction = grid.NextPointAroundCube(common.Point{X: 149, Y: 49}, DOWN)
	assert.Equal(t, common.Point{X: 99, Y: 99}, point, "Failed testing 2D")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// right edge face 3, move right
	point, direction = grid.NextPointAroundCube(common.Point{X: 99, Y: 50}, RIGHT)
	assert.Equal(t, common.Point{X: 100, Y: 49}, point, "Failed testing 3R")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// left edge face 3, move left
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 50}, LEFT)
	assert.Equal(t, common.Point{X: 0, Y: 100}, point, "Failed testing 3L")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// top edge face 3, move up
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 50}, UP)
	assert.Equal(t, common.Point{X: 50, Y: 49}, point, "Failed testing 3U")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// bottom edge face 3, move down
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 99}, DOWN)
	assert.Equal(t, common.Point{X: 50, Y: 100}, point, "Failed testing 3D")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// right edge face 4, move right
	point, direction = grid.NextPointAroundCube(common.Point{X: 99, Y: 100}, RIGHT)
	assert.Equal(t, common.Point{X: 149, Y: 49}, point, "Failed testing 4R")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// left edge face 4, move left
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 100}, LEFT)
	assert.Equal(t, common.Point{X: 49, Y: 100}, point, "Failed testing 4L")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// top edge face 4, move up
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 100}, UP)
	assert.Equal(t, common.Point{X: 50, Y: 99}, point, "Failed testing 4U")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// bottom edge face 4, move down
	point, direction = grid.NextPointAroundCube(common.Point{X: 50, Y: 149}, DOWN)
	assert.Equal(t, common.Point{X: 49, Y: 150}, point, "Failed testing 4D")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// right edge face 5, move right
	point, direction = grid.NextPointAroundCube(common.Point{X: 49, Y: 100}, RIGHT)
	assert.Equal(t, common.Point{X: 50, Y: 100}, point, "Failed testing 5R")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// left edge face 5, move left
	point, direction = grid.NextPointAroundCube(common.Point{X: 0, Y: 100}, LEFT)
	assert.Equal(t, common.Point{X: 50, Y: 49}, point, "Failed testing 5L")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// top edge face 5, move up
	point, direction = grid.NextPointAroundCube(common.Point{X: 0, Y: 100}, UP)
	assert.Equal(t, common.Point{X: 50, Y: 50}, point, "Failed testing 5U")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// bottom edge face 5, move down
	point, direction = grid.NextPointAroundCube(common.Point{X: 0, Y: 149}, DOWN)
	assert.Equal(t, common.Point{X: 0, Y: 150}, point, "Failed testing 5D")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// right edge face 6, move right
	point, direction = grid.NextPointAroundCube(common.Point{X: 49, Y: 150}, RIGHT)
	assert.Equal(t, common.Point{X: 50, Y: 149}, point, "Failed testing 6R")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// left edge face 6, move left
	point, direction = grid.NextPointAroundCube(common.Point{X: 0, Y: 150}, LEFT)
	assert.Equal(t, common.Point{X: 50, Y: 0}, point, "Failed testing 6L")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// top edge face 6, move up
	point, direction = grid.NextPointAroundCube(common.Point{X: 0, Y: 150}, UP)
	assert.Equal(t, common.Point{X: 0, Y: 149}, point, "Failed testing 6U")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// bottom edge face 6, move down
	point, direction = grid.NextPointAroundCube(common.Point{X: 0, Y: 199}, DOWN)
	assert.Equal(t, common.Point{X: 100, Y: 0}, point, "Failed testing 6D")
	assert.Equal(t, DOWN, direction, "Failed testing grid")
}

func TestGridNextPointAroundCubeTest(t *testing.T) {
	data = strings.Split(data, "\n\n")[0]
	grid := NewGrid(data, 4)

	// right edge face 1, move right
	point, direction := grid.NextPointAroundCubeTest(common.Point{X: 11, Y: 0}, RIGHT)
	assert.Equal(t, common.Point{X: 15, Y: 11}, point, "Failed testing face 1 right")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// left edge face 1, move left
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 0}, LEFT)
	assert.Equal(t, common.Point{X: 4, Y: 4}, point, "Failed testing face 1 left")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// top edge face 1, move up
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 0}, UP)
	assert.Equal(t, common.Point{X: 3, Y: 4}, point, "Failed testing face 1 up")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// bottom edge face 1, move down
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 3}, DOWN)
	assert.Equal(t, common.Point{X: 8, Y: 4}, point, "Failed testing face 1 down")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// right edge face 2, move right
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 3, Y: 4}, RIGHT)
	assert.Equal(t, common.Point{X: 4, Y: 4}, point, "Failed testing face 2 right")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// left edge face 2, move left
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 0, Y: 4}, LEFT)
	assert.Equal(t, common.Point{X: 15, Y: 11}, point, "Failed testing 2L")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// top edge face 2, move up
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 0, Y: 4}, UP)
	assert.Equal(t, common.Point{X: 11, Y: 0}, point, "Failed testing 2U")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// bottom edge face 2, move down
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 0, Y: 7}, DOWN)
	assert.Equal(t, common.Point{X: 11, Y: 11}, point, "Failed testing 2D")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// right edge face 3, move right
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 7, Y: 4}, RIGHT)
	assert.Equal(t, common.Point{X: 8, Y: 4}, point, "Failed testing 3R")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// left edge face 3, move left
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 4, Y: 4}, LEFT)
	assert.Equal(t, common.Point{X: 3, Y: 4}, point, "Failed testing 3L")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// top edge face 3, move up
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 4, Y: 4}, UP)
	assert.Equal(t, common.Point{X: 8, Y: 0}, point, "Failed testing 3U")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// bottom edge face 3, move down
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 7, Y: 7}, DOWN)
	assert.Equal(t, common.Point{X: 8, Y: 8}, point, "Failed testing 3D")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// right edge face 4, move right
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 11, Y: 4}, RIGHT)
	assert.Equal(t, common.Point{X: 15, Y: 8}, point, "Failed testing 4R")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// left edge face 4, move left
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 4}, LEFT)
	assert.Equal(t, common.Point{X: 7, Y: 4}, point, "Failed testing 4L")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// top edge face 4, move up
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 4}, UP)
	assert.Equal(t, common.Point{X: 8, Y: 3}, point, "Failed testing 4U")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// bottom edge face 4, move down
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 7}, DOWN)
	assert.Equal(t, common.Point{X: 8, Y: 8}, point, "Failed testing 4D")
	assert.Equal(t, DOWN, direction, "Failed testing grid")

	// right edge face 5, move right
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 11, Y: 9}, RIGHT)
	assert.Equal(t, common.Point{X: 12, Y: 9}, point, "Failed testing 5R")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")

	// left edge face 5, move left
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 8, Y: 9}, LEFT)
	assert.Equal(t, common.Point{X: 6, Y: 7}, point, "Failed testing 5L")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// top edge face 5, move up
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 9, Y: 8}, UP)
	assert.Equal(t, common.Point{X: 9, Y: 7}, point, "Failed testing 5U")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// bottom edge face 5, move down
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 9, Y: 11}, DOWN)
	assert.Equal(t, common.Point{X: 2, Y: 7}, point, "Failed testing 5D")
	assert.Equal(t, UP, direction, "Failed testing grid")

	// right edge face 6, move right
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 15, Y: 9}, RIGHT)
	assert.Equal(t, common.Point{X: 11, Y: 2}, point, "Failed testing 6R")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// left edge face 6, move left
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 12, Y: 9}, LEFT)
	assert.Equal(t, common.Point{X: 11, Y: 9}, point, "Failed testing 6L")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// top edge face 6, move up
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 13, Y: 8}, UP)
	assert.Equal(t, common.Point{X: 11, Y: 6}, point, "Failed testing 6U")
	assert.Equal(t, LEFT, direction, "Failed testing grid")

	// bottom edge face 6, move down
	point, direction = grid.NextPointAroundCubeTest(common.Point{X: 13, Y: 11}, DOWN)
	assert.Equal(t, common.Point{X: 0, Y: 6}, point, "Failed testing 6D")
	assert.Equal(t, RIGHT, direction, "Failed testing grid")
}
