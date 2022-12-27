package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pducolin/advent-of-code/2022/common"
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
	grid := NewGrid(data)
	for i := 0; i < 10; i++ {
		grid.Iterate()
	}
	return strconv.Itoa(grid.CountEmptyTiles())
}

func part2(data string) string {
	grid := NewGrid(data)
	previousState := grid.ToString()
	iterationCount := 0
	for {
		grid.Iterate()
		iterationCount++
		currentState := grid.ToString()
		if strings.Join(currentState, "\n") == strings.Join(previousState, "\n") {
			break
		}
		previousState = currentState
	}
	return strconv.Itoa(iterationCount)
}

type Grid struct {
	elves     map[common.Point]struct{}
	moveIndex int
}

type Direction int

const (
	N Direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

var (
	DIRECTIONS = []Direction{N, NE, E, SE, S, SW, W, NW}
	MOVES      = []Direction{N, S, W, E}
)

func NewGrid(data string) Grid {
	grid := Grid{
		elves: map[common.Point]struct{}{},
	}
	for y, line := range strings.Split(data, "\n") {
		for x, r := range line {
			if r == '#' {
				point := common.Point{X: x, Y: y}
				grid.elves[point] = struct{}{}
			}
		}
	}
	return grid
}

func (grid *Grid) Iterate() {
	proposedPointCountByPoint := map[common.Point]int{}
	movingElvesByOrigin := map[common.Point]common.Point{}
	stillElves := []common.Point{}

	// first half
	for elfPosition := range grid.elves {
		if !grid.HasNeighboursAround(elfPosition) {
			stillElves = append(stillElves, elfPosition)
			continue
		}

		canMove := false
		for i := 0; i < len(MOVES); i++ {
			moveIndex := (grid.moveIndex + i) % len(MOVES)
			if grid.canMove(elfPosition, MOVES[moveIndex]) {
				newElfPosition := GetNeighbour(elfPosition, MOVES[moveIndex])
				movingElvesByOrigin[elfPosition] = newElfPosition
				if _, found := proposedPointCountByPoint[newElfPosition]; !found {
					proposedPointCountByPoint[newElfPosition] = 0
				}
				proposedPointCountByPoint[newElfPosition]++
				canMove = true
				break
			}
		}
		if !canMove {
			stillElves = append(stillElves, elfPosition)
		}
	}

	// second half
	newElves := map[common.Point]struct{}{}
	for oldElfPosition, newElfPosition := range movingElvesByOrigin {
		if proposedPointCountByPoint[newElfPosition] > 1 {
			stillElves = append(stillElves, oldElfPosition)
			continue
		}
		newElves[newElfPosition] = struct{}{}
	}

	if len(newElves)+len(stillElves) != len(grid.elves) {
		panic("elves were lost")
	}

	for _, elf := range stillElves {
		if _, found := newElves[elf]; found {
			panic("illegal elf move")
		}
		newElves[elf] = struct{}{}
	}

	// update elves
	grid.elves = map[common.Point]struct{}{}
	grid.elves = newElves

	// last part, update move index
	grid.moveIndex = (grid.moveIndex + 1) % len(MOVES)
}

func (grid *Grid) canMove(point common.Point, direction Direction) bool {
	var directionsToCheck []Direction
	switch direction {
	case N:
		directionsToCheck = []Direction{NE, N, NW}
	case S:
		directionsToCheck = []Direction{SE, S, SW}
	case W:
		directionsToCheck = []Direction{W, SW, NW}
	case E:
		directionsToCheck = []Direction{E, SE, NE}
	default:
		panic(fmt.Errorf("invalid direction to check %#v", direction))
	}

	for _, d := range directionsToCheck {
		point := GetNeighbour(point, d)
		if _, found := grid.elves[point]; found {
			return false
		}
	}
	return true
}

func (grid *Grid) HasNeighboursAround(point common.Point) bool {
	countNeighbours := 0
	for _, direction := range DIRECTIONS {
		neighbour := GetNeighbour(point, direction)
		if _, found := grid.elves[neighbour]; found {
			countNeighbours++
		}
	}
	return countNeighbours > 0
}

func GetNeighbour(point common.Point, direction Direction) common.Point {
	switch direction {
	case N:
		return common.Point{X: point.X, Y: point.Y - 1}
	case NE:
		return common.Point{X: point.X + 1, Y: point.Y - 1}
	case E:
		return common.Point{X: point.X + 1, Y: point.Y}
	case SE:
		return common.Point{X: point.X + 1, Y: point.Y + 1}
	case S:
		return common.Point{X: point.X, Y: point.Y + 1}
	case SW:
		return common.Point{X: point.X - 1, Y: point.Y + 1}
	case W:
		return common.Point{X: point.X - 1, Y: point.Y}
	case NW:
		return common.Point{X: point.X - 1, Y: point.Y - 1}
	}
	panic(fmt.Errorf("unknown direction %#v", direction))
}

func (grid *Grid) getLimits() (minX, maxX, minY, maxY int) {
	minX = math.MaxInt
	maxX = math.MinInt
	minY = math.MaxInt
	maxY = math.MinInt

	for position := range grid.elves {
		if position.X < minX {
			minX = position.X
		}
		if position.X > maxX {
			maxX = position.X
		}
		if position.Y < minY {
			minY = position.Y
		}
		if position.Y > maxY {
			maxY = position.Y
		}
	}
	return minX, maxX, minY, maxY
}

func (grid *Grid) CountEmptyTiles() int {
	minX, maxX, minY, maxY := grid.getLimits()
	gridWidth := maxX - minX + 1
	gridHeight := maxY - minY + 1

	return gridWidth*gridHeight - len(grid.elves)
}

func (grid *Grid) ToString() []string {
	minX, maxX, minY, maxY := grid.getLimits()

	lines := []string{}
	for y := minY; y <= maxY; y++ {
		line := ""
		for x := minX; x <= maxX; x++ {
			if _, found := grid.elves[common.Point{X: x, Y: y}]; found {
				line += "#"
				continue
			}
			line += "."
		}
		lines = append(lines, line)
	}
	return lines
}

func (grid *Grid) Print() {
	for _, line := range grid.ToString() {
		fmt.Println(line)
	}
}
