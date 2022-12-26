package lava

import (
	"math"
	"strconv"
	"strings"

	"github.com/pducolin/advent-of-code/2022/common"
)

type Grid struct {
	LavaCubes                          map[Cube]struct{}
	waterCubes                         map[Cube]struct{}
	xMin, xMax, yMin, yMax, zMin, zMax int
}

func NewGrid(data string) Grid {
	grid := parseFromData(data)
	grid.findLimits()
	return grid
}

func parseFromData(data string) Grid {
	grid := Grid{
		LavaCubes: map[Cube]struct{}{},
	}
	for _, line := range strings.Split(data, "\n") {
		xyz := strings.Split(line, ",")
		x, err := strconv.Atoi(xyz[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(xyz[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(xyz[2])
		if err != nil {
			panic(err)
		}
		grid.LavaCubes[Cube{X: x, Y: y, Z: z}] = struct{}{}
	}
	return grid
}

func (g *Grid) findLimits() {
	g.xMin = math.MaxInt
	g.yMin = math.MaxInt
	g.zMin = math.MaxInt
	g.xMax = math.MinInt
	g.yMax = math.MinInt
	g.zMax = math.MinInt

	for cube := range g.LavaCubes {
		// min
		if cube.X < g.xMin {
			g.xMin = cube.X
		}
		if cube.Y < g.yMin {
			g.yMin = cube.Y
		}
		if cube.Z < g.zMin {
			g.zMin = cube.Z
		}
		// max
		if cube.X > g.xMax {
			g.xMax = cube.X
		}
		if cube.Y > g.yMax {
			g.yMax = cube.Y
		}
		if cube.Z > g.zMax {
			g.zMax = cube.Z
		}
	}

	g.xMin -= 1
	g.yMin -= 1
	g.zMin -= 1

	g.xMax += 1
	g.yMax += 1
	g.zMax += 1
}

func (g *Grid) CountFreeSides(cube Cube) int {
	count := 0
	for _, neighbour := range cube.GetNeighbours() {
		if _, found := g.LavaCubes[neighbour]; !found {
			count++
		}
	}
	return count
}

func (g *Grid) CountReachableSides(cube Cube) int {
	count := 0
	for _, neighbour := range cube.GetNeighbours() {
		if _, found := g.LavaCubes[neighbour]; found {
			continue
		}
		if _, found := g.waterCubes[neighbour]; !found {
			continue
		}
		count++
	}
	return count
}

// Flood-fill (node):
//  1. Set Q to the empty queue or stack.
//  2. Add node to the end of Q.
//  3. While Q is not empty:
//  4. Set n equal to the first element of Q.
//  5. Remove first element from Q.
//  6. If n is Inside:
//     Set the n
//     Add the node to the west of n to the end of Q.
//     Add the node to the east of n to the end of Q.
//     Add the node to the north of n to the end of Q.
//     Add the node to the south of n to the end of Q.
//     Add the node to the front of n to the end of Q.
//     Add the node to the bottom of n to the end of Q.
//  7. Continue looping until Q is exhausted.
//  8. Return.
//
// Source https://en.wikipedia.org/wiki/Flood_fill
func (g *Grid) FloodFill() {
	g.waterCubes = map[Cube]struct{}{}

	startPosition := Cube{0, 0, 0}
	visited := map[Cube]struct{}{startPosition: {}}
	toVisit := common.NewQueue[Cube]()
	toVisit.Push(startPosition)

	for !toVisit.IsEmpty() {
		cube := toVisit.Pop()

		g.waterCubes[cube] = struct{}{}

		neighbours := cube.GetNeighbours()
		for _, direction := range DIRECTIONS {
			neighbour := neighbours[direction]
			if _, found := visited[neighbour]; found {
				continue
			}

			if g.IsExternal(neighbour) {
				continue
			}

			if _, isLava := g.LavaCubes[neighbour]; isLava {
				continue
			}
			visited[neighbour] = struct{}{}
			toVisit.Push(neighbour)
		}
	}
}

func (g *Grid) IsExternal(cube Cube) bool {
	return cube.X < g.xMin ||
		cube.X > g.xMax ||
		cube.Y < g.yMin ||
		cube.Y > g.yMax ||
		cube.Z < g.zMin ||
		cube.Z > g.zMax
}
