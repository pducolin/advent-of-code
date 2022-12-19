package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)

	if part == 1 {
		fmt.Println(part1(inputData))
	} else {
		fmt.Println(part2(inputData))
	}
}

func part1(data string) string {
	grid := parseFromData(data)
	res := 0
	for cube := range grid.cubes {
		res += grid.CountFreeSides(cube)
	}
	return strconv.Itoa(res)
}

func part2(data string) string {
	return "Hello World !"
}

type Cube struct {
	x int
	y int
	z int
}

type Grid struct {
	cubes map[Cube]struct{}
}

func parseFromData(data string) Grid {
	grid := Grid{
		cubes: map[Cube]struct{}{},
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
		grid.cubes[Cube{x: x, y: y, z: z}] = struct{}{}
	}
	return grid
}

func (g *Grid) CountFreeSides(cube Cube) int {
	count := 0
	for _, neighbour := range cube.GetNeighbours() {
		if _, found := g.cubes[neighbour]; !found {
			count++
		}
	}
	return count
}

func (c *Cube) GetNeighbours() []Cube {
	neighbours := []Cube{}

	neighbours = append(neighbours, Cube{
		x: c.x - 1,
		y: c.y,
		z: c.z,
	})
	neighbours = append(neighbours, Cube{
		x: c.x + 1,
		y: c.y,
		z: c.z,
	})
	neighbours = append(neighbours, Cube{
		x: c.x,
		y: c.y - 1,
		z: c.z,
	})
	neighbours = append(neighbours, Cube{
		x: c.x,
		y: c.y + 1,
		z: c.z,
	})

	neighbours = append(neighbours, Cube{
		x: c.x,
		y: c.y,
		z: c.z - 1,
	})
	neighbours = append(neighbours, Cube{
		x: c.x,
		y: c.y,
		z: c.z + 1,
	})

	return neighbours
}
