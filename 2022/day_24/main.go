package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	playerPosition := grid.StartPosition()

	ret, grid, err := FindShortestPath(grid, playerPosition, grid.EndPosition())
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(ret)
}

func part2(data string) string {
	grid := NewGrid(data)

	totTimeElapsed := 0

	steps := []Step{
		Step{
			from: grid.StartPosition(),
			to:   grid.EndPosition(),
		},
		Step{
			from: grid.EndPosition(),
			to:   grid.StartPosition(),
		},
		Step{
			from: grid.StartPosition(),
			to:   grid.EndPosition(),
		},
	}
	for _, step := range steps {
		ret, newGrid, err := FindShortestPath(grid, step.from, step.to)
		if err != nil {
			panic(err)
		}
		totTimeElapsed += ret
		grid = newGrid
	}
	return strconv.Itoa(totTimeElapsed)
}

type Step struct {
	from common.Point
	to   common.Point
}

type Grid struct {
	blizzards     map[common.Point][]Blizzard
	width, height int
	timeElapsed   int
}

type Blizzard struct {
	id        int
	direction Direction
	position  common.Point
}

func isBlizzard(r rune) bool {
	return r == '>' || r == 'v' || r == '<' || r == '^'
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var BLIZZARD_TO_DIRECTION = map[rune]Direction{
	'>': RIGHT,
	'^': UP,
	'<': LEFT,
	'v': DOWN,
}

var DIRECTION_TO_BLIZZARD = map[Direction]string{
	RIGHT: ">",
	UP:    "^",
	LEFT:  "<",
	DOWN:  "v",
}

func NewGrid(data string) Grid {
	grid := Grid{
		blizzards:   map[common.Point][]Blizzard{},
		timeElapsed: 0,
	}
	lines := strings.Split(data, "\n")

	grid.height = len(lines)
	grid.width = len(lines[0])

	blizzardId := 1
	for y, line := range lines {
		for x, r := range line {
			if isBlizzard(r) {
				position := common.Point{X: x, Y: y}
				if _, found := grid.blizzards[position]; !found {
					grid.blizzards[position] = []Blizzard{}
				}

				grid.blizzards[position] = append(grid.blizzards[position], Blizzard{
					id:        blizzardId,
					position:  position,
					direction: BLIZZARD_TO_DIRECTION[r],
				})

				blizzardId++
			}
		}
	}

	return grid
}

func (grid *Grid) StartPosition() common.Point {
	return common.Point{X: 1, Y: 0}
}

func (grid *Grid) EndPosition() common.Point {
	return common.Point{X: grid.width - 2, Y: grid.height - 1}
}

func (grid *Grid) ToString() []string {
	lines := []string{}
	for y := 0; y < grid.height; y++ {
		line := ""
		for x := 0; x < grid.width; x++ {
			position := common.Point{X: x, Y: y}
			if position == grid.StartPosition() || position == grid.EndPosition() {
				line += "."
				continue
			}
			if grid.IsWall(position) {
				line += "#"
				continue
			}
			if _, found := grid.blizzards[position]; !found {
				line += "."
				continue
			}

			blizzards := grid.blizzards[position]
			if len(blizzards) > 1 {
				line += strconv.Itoa(len(blizzards))
				continue
			}
			line += DIRECTION_TO_BLIZZARD[blizzards[0].direction]
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

func (grid *Grid) IsWall(position common.Point) bool {
	if position == grid.StartPosition() || position == grid.EndPosition() {
		return false
	}

	return (position.X <= 0 || position.X >= grid.width-1 ||
		position.Y <= 0 || position.Y >= grid.height-1)
}

func (grid *Grid) Iterate() Grid {
	newGrid := Grid{
		blizzards:   map[common.Point][]Blizzard{},
		width:       grid.width,
		height:      grid.height,
		timeElapsed: grid.timeElapsed + 1,
	}
	for _, blizzards := range grid.blizzards {
		for _, blizzard := range blizzards {
			movedBlizzard := grid.iterateBlizzard(blizzard)
			if _, found := newGrid.blizzards[movedBlizzard.position]; !found {
				newGrid.blizzards[movedBlizzard.position] = []Blizzard{}
			}
			newGrid.blizzards[movedBlizzard.position] = append(newGrid.blizzards[movedBlizzard.position], movedBlizzard)
		}
	}
	return newGrid
}

func (grid *Grid) iterateBlizzard(blizzard Blizzard) Blizzard {
	switch blizzard.direction {
	case UP:
		position := common.Point{
			X: blizzard.position.X,
			Y: blizzard.position.Y - 1,
		}
		if grid.IsWall(position) {
			position.Y = grid.height - 2
		}
		blizzard.position = position
		return blizzard
	case DOWN:
		position := common.Point{
			X: blizzard.position.X,
			Y: blizzard.position.Y + 1,
		}
		if grid.IsWall(position) {
			position.Y = 1
		}
		blizzard.position = position
		return blizzard
	case LEFT:
		position := common.Point{
			X: blizzard.position.X - 1,
			Y: blizzard.position.Y,
		}
		if grid.IsWall(position) {
			position.X = grid.width - 2
		}
		blizzard.position = position
		return blizzard
	case RIGHT:
		position := common.Point{
			X: blizzard.position.X + 1,
			Y: blizzard.position.Y,
		}
		if grid.IsWall(position) {
			position.X = 1
		}
		blizzard.position = position
		return blizzard
	}
	return blizzard
}

type PlayerState struct {
	player      common.Point
	timeElapsed int
}

func (grid *Grid) GetMoves(position common.Point) []common.Point {
	moves := []common.Point{}

	nextPosition := moveTo(position, DOWN)
	if grid.isValidPlayerPosition(nextPosition) {
		moves = append(moves, nextPosition)
	}
	nextPosition = moveTo(position, RIGHT)
	if grid.isValidPlayerPosition(nextPosition) {
		moves = append(moves, nextPosition)
	}
	// stay still
	if grid.isValidPlayerPosition(position) {
		moves = append(moves, position)
	}
	// UP
	nextPosition = moveTo(position, UP)
	if grid.isValidPlayerPosition(nextPosition) {
		moves = append(moves, nextPosition)
	}
	// LEFT
	nextPosition = moveTo(position, LEFT)
	if grid.isValidPlayerPosition(nextPosition) {
		moves = append(moves, nextPosition)
	}
	return moves
}

func FindShortestPath(initialGrid Grid, startingPoint common.Point, targetPoint common.Point) (shortestPath int, finalGrid Grid, err error) {
	visited := map[PlayerState]struct{}{}

	// Inspired by @ValerieMauduit
	// and her tip: "Blizzard move over time, they do not depend on our position"
	// At a given moment in time, we have one and only one grid
	// Time is the third dimension of the grid
	gridByTime := []Grid{
		initialGrid,
	}

	queue := common.NewQueue[PlayerState]()

	queue.Push(PlayerState{
		player:      startingPoint,
		timeElapsed: 0,
	})

	for !queue.IsEmpty() {
		currentState := queue.Pop()

		if currentState.player == targetPoint {
			return currentState.timeElapsed, gridByTime[currentState.timeElapsed], nil
		}

		if _, found := visited[currentState]; found {
			continue
		}

		nextTime := currentState.timeElapsed + 1
		// iterate grid to get next minute grid, if it does not exist
		if len(gridByTime)-1 < nextTime {
			gridByTime = append(gridByTime, gridByTime[currentState.timeElapsed].Iterate())
		}
		nextGrid := gridByTime[nextTime]

		for _, move := range nextGrid.GetMoves(currentState.player) {
			queue.Push(PlayerState{
				player:      move,
				timeElapsed: nextTime})
		}

		visited[currentState] = struct{}{}
	}

	return -1, initialGrid, fmt.Errorf("no path from %#v", startingPoint)
}

func (grid *Grid) isValidPlayerPosition(position common.Point) bool {
	if grid.IsWall(position) {
		return false
	}

	if _, found := grid.blizzards[position]; found {
		return false
	}

	return true
}

func moveTo(position common.Point, direction Direction) common.Point {
	switch direction {
	case UP:
		return common.Point{X: position.X, Y: position.Y - 1}
	case DOWN:
		return common.Point{X: position.X, Y: position.Y + 1}
	case LEFT:
		return common.Point{X: position.X - 1, Y: position.Y}
	case RIGHT:
		return common.Point{X: position.X + 1, Y: position.Y}
	}
	panic("invalid direction")
}
