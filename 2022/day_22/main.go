package main

import (
	_ "embed"
	"errors"
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
		fmt.Println(part1(inputData, 50))
	} else {
		useRealData := true
		fmt.Println(part2(inputData, 50, useRealData))
	}
}

func part1(data string, size int) string {
	parts := strings.Split(data, "\n\n")

	grid := NewGrid(parts[0], size)

	instructions := parts[1]
	point := common.Point{X: grid.minY[0], Y: 0}
	direction := RIGHT
	steps := 0
	for _, char := range strings.Split(instructions, "") {
		val, err := strconv.Atoi(char)
		if err == nil {
			steps = steps*10 + val
			continue
		}
		// move and turn direction
		point = grid.Move(point, direction, steps)
		steps = 0
		direction = Turn(char, direction)
	}
	point = grid.Move(point, direction, steps)

	password := (point.Y+1)*1000 + (point.X+1)*4 + int(direction)

	return strconv.Itoa(password)
}

func part2(data string, size int, useRealData bool) string {
	parts := strings.Split(data, "\n\n")

	grid := NewGrid(parts[0], size)

	instructions := parts[1]

	var point common.Point
	if useRealData {
		point = common.Point{X: grid.size, Y: 0}
	} else {
		point = common.Point{X: 2 * grid.size, Y: 0}
	}
	direction := RIGHT
	steps := 0
	for _, char := range strings.Split(instructions, "") {
		val, err := strconv.Atoi(char)
		if err == nil {
			steps = steps*10 + val
			continue
		}
		// move and turn direction
		if useRealData {
			point, direction = grid.Move3D(point, direction, steps)
		} else {
			point, direction = grid.Move3DTest(point, direction, steps)
		}
		steps = 0
		direction = Turn(char, direction)
	}
	if useRealData {
		point, direction = grid.Move3D(point, direction, steps)
	} else {
		point, direction = grid.Move3DTest(point, direction, steps)
	}

	password := (point.Y+1)*1000 + (point.X+1)*4 + int(direction)

	return strconv.Itoa(password)
}

type Direction int

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
)

type Grid struct {
	points                 map[common.Point]string
	minY, maxY, minX, maxX []int
	size                   int
}

func NewGrid(data string, size int) Grid {
	grid := Grid{
		points: map[common.Point]string{},
		minY:   []int{},
		maxY:   []int{},
		minX:   []int{},
		maxX:   []int{},
		size:   size,
	}

	lines := strings.Split(data, "\n")

	maxColumns := math.MinInt

	for _, line := range lines {
		if len(line) > maxColumns {
			maxColumns = len(line)
		}
	}

	for x := 0; x < maxColumns; x++ {
		grid.maxY = append(grid.maxY, math.MinInt)
		grid.minY = append(grid.minY, math.MaxInt)
	}

	for y, line := range lines {
		grid.minX = append(grid.minX, math.MaxInt)
		grid.maxX = append(grid.maxX, math.MinInt)
		for x, char := range strings.Split(line, "") {
			if char != "." && char != "#" {
				continue
			}
			grid.points[common.Point{X: x, Y: y}] = char
			if x < grid.minX[y] {
				grid.minX[y] = x
			}
			if x > grid.maxX[y] {
				grid.maxX[y] = x
			}
			if y < grid.minY[x] {
				grid.minY[x] = y
			}
			if y > grid.maxY[x] {
				grid.maxY[x] = y
			}
		}
	}

	return grid
}

func (grid *Grid) EnsureFacesTest() {
	// ensure face 1
	for x := 2 * grid.size; x < 3*grid.size; x++ {
		for y := 0; y < grid.size; y++ {
			point := common.Point{X: x, Y: y}
			if _, found := grid.points[point]; !found {
				panic(fmt.Errorf("missing point %#v in face 1", point))
			}
		}
	}
	// ensure face 2
	for x := 0; x < grid.size; x++ {
		for y := grid.size; y < 2*grid.size; y++ {
			point := common.Point{X: x, Y: y}
			if _, found := grid.points[point]; !found {
				panic(fmt.Errorf("missing point %#v in face 2", point))
			}
		}
	}
	// ensure face 3
	for x := grid.size; x < 2*grid.size; x++ {
		for y := grid.size; y < 2*grid.size; y++ {
			point := common.Point{X: x, Y: y}
			if _, found := grid.points[point]; !found {
				panic(fmt.Errorf("missing point %#v in face 3", point))
			}
		}
	}
	// ensure face 4
	for x := 2 * grid.size; x < 3*grid.size; x++ {
		for y := grid.size; y < 2*grid.size; y++ {
			point := common.Point{X: x, Y: y}
			if _, found := grid.points[point]; !found {
				panic(fmt.Errorf("missing point %#v in face 4", point))
			}
		}
	}
	// ensure face 5
	for x := 2 * grid.size; x < 3*grid.size; x++ {
		for y := grid.size; y < 2*grid.size; y++ {
			point := common.Point{X: x, Y: y}
			if _, found := grid.points[point]; !found {
				panic(fmt.Errorf("missing point %#v in face 4", point))
			}
		}
	}
}

// 2d
func (g *Grid) Move(from common.Point, direction Direction, steps int) common.Point {
	point := common.Point{
		X: from.X,
		Y: from.Y,
	}
	for i := 0; i < steps; i++ {
		newPoint, err := g.Step(point, direction)
		if err != nil {
			break
		}
		point = newPoint
	}
	return point
}

func (g *Grid) Step(from common.Point, direction Direction) (common.Point, error) {
	nextPoint := NextPoint(from, direction)

	nextValue, found := g.points[nextPoint]

	if found {
		if nextValue == "." {
			return nextPoint, nil
		}
		return nextPoint, errors.New("wall")
	}
	// no point to the right, let's wrap
	nextPoint = g.Wrap(from, direction)

	nextValue, found = g.points[nextPoint]

	if !found {
		panic(fmt.Errorf("No wrapped point to the right in grid from %#v", from))
	}

	if nextValue == "." {
		return nextPoint, nil
	}
	return nextPoint, errors.New("wall")
}

func NextPoint(from common.Point, direction Direction) common.Point {
	switch direction {
	case UP:
		return nextPointUp(from)
	case LEFT:
		return nextPointLeft(from)
	case DOWN:
		return nextPointDown(from)
	case RIGHT:
		return nextPointRight(from)
	}
	panic(fmt.Errorf("Unknown direction %#v", direction))
}

func nextPointLeft(from common.Point) common.Point {
	return common.Point{
		X: from.X - 1,
		Y: from.Y,
	}
}

func nextPointRight(from common.Point) common.Point {
	return common.Point{
		X: from.X + 1,
		Y: from.Y,
	}
}

func nextPointUp(from common.Point) common.Point {
	return common.Point{
		X: from.X,
		Y: from.Y - 1,
	}
}

func nextPointDown(from common.Point) common.Point {
	return common.Point{
		X: from.X,
		Y: from.Y + 1,
	}
}

func (g *Grid) Wrap(from common.Point, direction Direction) common.Point {
	switch direction {
	case UP:
		return g.wrapUp(from)
	case LEFT:
		return g.wrapLeft(from)
	case DOWN:
		return g.wrapDown(from)
	case RIGHT:
		return g.wrapRight(from)
	}
	panic(fmt.Errorf("Unknown direction %#v", direction))
}

func (g *Grid) wrapLeft(from common.Point) common.Point {
	return common.Point{
		X: g.maxX[from.Y],
		Y: from.Y,
	}
}

func (g *Grid) wrapRight(from common.Point) common.Point {
	return common.Point{
		X: g.minX[from.Y],
		Y: from.Y,
	}
}

func (g *Grid) wrapUp(from common.Point) common.Point {
	return common.Point{
		X: from.X,
		Y: g.maxY[from.X],
	}
}

func (g *Grid) wrapDown(from common.Point) common.Point {
	return common.Point{
		X: from.X,
		Y: g.minY[from.X],
	}
}

func Turn(letter string, direction Direction) Direction {
	if letter == "R" {
		return turnRight(direction)
	}
	return turnLeft(direction)
}

func turnRight(direction Direction) Direction {
	switch direction {
	case UP:
		return RIGHT
	case LEFT:
		return UP
	case DOWN:
		return LEFT
	case RIGHT:
		return DOWN
	}
	panic(fmt.Errorf("Unknown direction %#v", direction))
}

func turnLeft(direction Direction) Direction {
	switch direction {
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	}
	panic(fmt.Errorf("Unknown direction %#v", direction))
}

// =============================================================
// =========================== 3d ==============================
// =============================================================
func (g *Grid) Move3D(from common.Point, direction Direction, steps int) (common.Point, Direction) {
	point := common.Point{
		X: from.X,
		Y: from.Y,
	}
	for i := 0; i < steps; i++ {
		newPoint, newDirection, err := g.Step3D(point, direction)
		if err != nil {
			break
		}
		point = newPoint
		direction = newDirection
	}
	return point, direction
}

func (g *Grid) Step3D(from common.Point, direction Direction) (common.Point, Direction, error) {
	currentFace := g.GetFace(from)
	if currentFace == -1 {
		panic("out of cube")
	}

	nextPoint := NextPoint(from, direction)
	nextFace := g.GetFace(nextPoint)

	if nextFace == currentFace {
		nextValue, found := g.points[nextPoint]
		if !found {
			panic(fmt.Errorf("out of cube moving %#v from point %#v", direction, from))
		}
		if nextValue == "." {
			return nextPoint, direction, nil
		}
		return nextPoint, direction, errors.New("wall")
	}

	// new face
	nextPoint, nextDirection := g.NextPointAroundCube(from, direction)
	nextValue, found := g.points[nextPoint]

	if !found {
		panic(fmt.Errorf("out of cube in direction %#v from %#v, nextPoint %#v", direction, from, nextPoint))
	}

	if nextValue == "." {
		return nextPoint, nextDirection, nil
	}
	return nextPoint, nextDirection, errors.New("wall")
}

func (g *Grid) GetFace(point common.Point) int {
	// face 1  size <= x < size * 2, 0 <= y < size
	if g.size <= point.X && point.X < g.size*2 && 0 <= point.Y && point.Y < g.size {
		return 1
	}
	// face 2  2*size <= x < 3*size, 0 <= y < size
	if 2*g.size <= point.X && point.X < 3*g.size && 0 <= point.Y && point.Y < g.size {
		return 2
	}
	// face 3  size <= x < size * 2, size <= y < size * 2
	if g.size <= point.X && point.X < g.size*2 && g.size <= point.Y && point.Y < g.size*2 {
		return 3
	}
	// face 4  size  <= x < size * 2, size * 2 <= y < size * 3
	if g.size <= point.X && point.X < g.size*2 && g.size*2 <= point.Y && point.Y < g.size*3 {
		return 4
	}
	// face 5  0 <= x < size , size * 2 <= y < size * 3
	if 0 <= point.X && point.X < g.size && g.size*2 <= point.Y && point.Y < g.size*3 {
		return 5
	}
	// face 6  0 <= x < size, size * 3 <= y < size * 4
	if 0 <= point.X && point.X < g.size && g.size*3 <= point.Y && point.Y < g.size*4 {
		return 6
	}
	return -1
}

func (g *Grid) NextPointAroundCube(from common.Point, direction Direction) (common.Point, Direction) {
	switch direction {
	case LEFT:
		return g.nextPointAroundCubeLeft(from)
	case UP:
		return g.nextPointAroundCubeUp(from)
	case DOWN:
		return g.nextPointAroundCubeDown(from)
	case RIGHT:
		return g.nextPointAroundCubeRight(from)
	}
	panic(fmt.Errorf("Unknown direction %#v", direction))
}

func (g *Grid) nextPointAroundCubeLeft(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFace(from)

	switch currentFace {
	case 1:
		// new face 5
		// {X:50,Y:0}  --> {X:0,Y:149}
		// {X:50,Y:49} --> {X:0,Y:100}
		return common.Point{
			X: 0,
			Y: g.size*3 - from.Y - 1,
		}, RIGHT
	case 2:
		// new face 1
		return common.Point{
			X: from.X - 1,
			Y: from.Y,
		}, LEFT
	case 3:
		// new face 5
		// {X:50,Y:50}  --> {X:0,Y:100}
		// {X:50,Y:99} --> {X:49,Y:100}
		return common.Point{
			X: from.Y - g.size,
			Y: g.size * 2,
		}, DOWN
	case 4:
		// new face 5
		return common.Point{
			X: from.X - 1,
			Y: from.Y,
		}, LEFT
	case 5:
		// new face 1
		// {X:0,Y:100}  --> {X:50,Y:49}
		// {X:0,Y:149} --> {X:50,Y:0}
		return common.Point{
			X: g.size,
			Y: g.size*3 - from.Y - 1,
		}, RIGHT
	case 6:
		// new face 1
		// {X:0,Y:150}  --> {X:50,Y:0}
		// {X:0,Y:199} --> {X:99,Y:0}
		return common.Point{
			X: from.Y - 2*g.size,
			Y: 0,
		}, DOWN
	}

	panic("from out of cube")
}

func (g *Grid) nextPointAroundCubeRight(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFace(from)

	switch currentFace {
	case 1:
		// new face 2
		return common.Point{
			X: from.X + 1,
			Y: from.Y,
		}, RIGHT
	case 2:
		// new face 4
		// {X:149,Y:0}  --> {X:99,Y:149}
		// {X:149,Y:49} --> {X:99,Y:100}
		return common.Point{
			X: 2*g.size - 1,
			Y: 3*g.size - 1 - from.Y,
		}, LEFT
	case 3:
		// new face 2
		// {X:99,Y:50} --> {X:100,Y:49}
		// {X:99,Y:99} --> {X:149,Y:49}
		return common.Point{
			X: g.size + from.Y,
			Y: g.size - 1,
		}, UP
	case 4:
		// new face 2
		// {X:99,Y:149} --> {X:149,Y:0}
		// {X:99,Y:100} --> {X:149,Y:49}
		return common.Point{
			X: 3*g.size - 1,
			Y: 3*g.size - 1 - from.Y,
		}, LEFT
	case 5:
		// new face 4
		return common.Point{
			X: from.X + 1,
			Y: from.Y,
		}, RIGHT
	case 6:
		// new face 4
		// {X:49,Y:150} --> {X:50,Y:149}
		// {X:49,Y:199} --> {X:99,Y:149}
		return common.Point{
			X: from.Y - 2*g.size,
			Y: g.size*3 - 1,
		}, UP
	}

	panic("from out of cube")
}

func (g *Grid) nextPointAroundCubeUp(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFace(from)

	switch currentFace {
	case 1:
		// new face 1
		// {X:50,Y:0} --> {X:0,Y:150}
		// {X:99,Y:0} --> {X:0,Y:199}
		return common.Point{
			X: 0,
			Y: 2*g.size + from.X,
		}, RIGHT
	case 2:
		// new face 6
		// {X:100,Y:0} --> {X:0,Y:199}
		// {X:149,Y:0} --> {X:49,Y:199}
		return common.Point{
			X: from.X - 2*g.size,
			Y: 4*g.size - 1,
		}, UP
	case 3:
		// new face 1
		return common.Point{
			X: from.X,
			Y: from.Y - 1,
		}, UP
	case 4:
		// new face 3
		return common.Point{
			X: from.X,
			Y: from.Y - 1,
		}, UP
	case 5:
		// new face 3
		// {X:0,Y:100} --> {X:50,Y:50}
		// {X:49,Y:100} --> {X:50,Y:99}
		return common.Point{
			X: g.size,
			Y: g.size + from.X,
		}, RIGHT
	case 6:
		// new face 5
		return common.Point{
			X: from.X,
			Y: from.Y - 1,
		}, UP
	}

	panic("from out of cube")
}

func (g *Grid) nextPointAroundCubeDown(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFace(from)

	switch currentFace {
	case 1:
		// new face 3
		return common.Point{
			X: from.X,
			Y: from.Y + 1,
		}, DOWN
	case 2:
		// new face 3
		// {X:100,Y:49} --> {X:99,Y:50}
		// {X:149,Y:49} --> {X:99,Y:99}
		return common.Point{
			X: 2*g.size - 1,
			Y: from.X - g.size,
		}, LEFT
	case 3:
		// new face 4
		return common.Point{
			X: from.X,
			Y: from.Y + 1,
		}, DOWN
	case 4:
		// new face 6
		// {X:50,Y:149} --> {X:49,Y:150}
		// {X:99,Y:149} --> {X:49,Y:199}
		return common.Point{
			X: g.size - 1,
			Y: 2*g.size + from.X,
		}, LEFT
	case 5:
		// new face 6
		return common.Point{
			X: from.X,
			Y: from.Y + 1,
		}, DOWN
	case 6:
		// new face 2
		// {X:0,Y:199} --> {X:100,Y:0}
		// {X:49,Y:199} --> {X:149,Y:0}
		return common.Point{
			X: 2*g.size + from.X,
			Y: 0,
		}, DOWN
	}

	panic("from out of cube")
}

// test

func (g *Grid) Move3DTest(from common.Point, direction Direction, steps int) (common.Point, Direction) {
	point := common.Point{
		X: from.X,
		Y: from.Y,
	}
	for i := 0; i < steps; i++ {
		newPoint, newDirection, err := g.Step3DTest(point, direction)
		if err != nil {
			break
		}
		point = newPoint
		direction = newDirection
	}
	return point, direction
}

func (g *Grid) Step3DTest(from common.Point, direction Direction) (common.Point, Direction, error) {
	currentFace := g.GetFaceTest(from)
	if currentFace == -1 {
		panic("out of cube")
	}

	nextPoint := NextPoint(from, direction)
	nextFace := g.GetFaceTest(nextPoint)

	if nextFace == currentFace {
		nextValue, found := g.points[nextPoint]
		if !found {
			panic(fmt.Errorf("out of cube moving %#v from point %#v", direction, from))
		}
		if nextValue == "." {
			return nextPoint, direction, nil
		}
		return nextPoint, direction, errors.New("wall")
	}

	// new face
	nextPoint, nextDirection := g.NextPointAroundCubeTest(from, direction)
	nextValue, found := g.points[nextPoint]

	if !found {
		panic(fmt.Errorf("out of cube in direction %#v from %#v, nextPoint %#v", direction, from, nextPoint))
	}

	if nextValue == "." {
		return nextPoint, nextDirection, nil
	}
	return nextPoint, nextDirection, errors.New("wall")
}

func (g *Grid) GetFaceTest(point common.Point) int {
	// face 1  size * 2 <= x < size * 3, 0 <= y < size
	if g.size*2 <= point.X && point.X < g.size*3 && 0 <= point.Y && point.Y < g.size {
		return 1
	}
	// face 2  0 <= x < size, size <= y < size * 2
	if 0 <= point.X && point.X < g.size && g.size <= point.Y && point.Y < g.size*2 {
		return 2
	}
	// face 3  size <= x < size * 2, size <= y < size * 2
	if g.size <= point.X && point.X < g.size*2 && g.size <= point.Y && point.Y < g.size*2 {
		return 3
	}
	// face 4  size * 2 <= x < size * 3, size <= y < size * 2
	if g.size*2 <= point.X && point.X < g.size*3 && g.size <= point.Y && point.Y < g.size*2 {
		return 4
	}
	// face 5  size * 2 <= x < size * 3, size * 2 <= y < size * 3
	if g.size*2 <= point.X && point.X < g.size*3 && g.size*2 <= point.Y && point.Y < g.size*3 {
		return 5
	}
	// face 6  size * 3 <= x < size * 4, size * 2 <= y < size * 3
	if g.size*3 <= point.X && point.X < g.size*4 && g.size*2 <= point.Y && point.Y < g.size*3 {
		return 6
	}
	return -1
}

func (g *Grid) NextPointAroundCubeTest(from common.Point, direction Direction) (common.Point, Direction) {
	switch direction {
	case LEFT:
		return g.nextPointAroundCubeLeftTest(from)
	case UP:
		return g.nextPointAroundCubeUpTest(from)
	case DOWN:
		return g.nextPointAroundCubeDownTest(from)
	case RIGHT:
		return g.nextPointAroundCubeRightTest(from)
	}
	panic(fmt.Errorf("Unknown direction %#v", direction))
}

func (g *Grid) nextPointAroundCubeLeftTest(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFaceTest(from)

	switch currentFace {
	case 1:
		// new face 3
		return common.Point{
			X: from.Y + g.size,
			Y: g.size,
		}, DOWN
	case 2:
		// new face 6
		return common.Point{
			X: g.size*5 - from.Y - 1,
			Y: g.size*3 - 1,
		}, UP
	case 3:
		// new face 2
		return common.Point{
			X: from.X - 1,
			Y: from.Y,
		}, LEFT
	case 4:
		// new face 3
		return common.Point{
			X: from.X - 1,
			Y: from.Y,
		}, LEFT
	case 5:
		// new face 3
		return common.Point{
			X: g.size*4 - from.Y - 1,
			Y: g.size*2 - 1,
		}, UP
	case 6:
		// new face 5
		return common.Point{
			X: from.X - 1,
			Y: from.Y,
		}, LEFT
	}

	panic("from out of cube")
}

func (g *Grid) nextPointAroundCubeRightTest(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFaceTest(from)

	switch currentFace {
	case 1:
		// new face 6
		return common.Point{
			X: g.size*4 - 1,
			Y: 3*g.size - 1 - from.Y,
		}, LEFT
	case 2:
		// new face 3
		return common.Point{
			X: from.X + 1,
			Y: from.Y,
		}, RIGHT
	case 3:
		// new face 4
		return common.Point{
			X: from.X + 1,
			Y: from.Y,
		}, RIGHT
	case 4:
		// new face 6
		return common.Point{
			X: g.size*5 - from.Y - 1,
			Y: g.size * 2,
		}, DOWN
	case 5:
		// new face 6
		return common.Point{
			X: from.X + 1,
			Y: from.Y,
		}, RIGHT
	case 6:
		// new face 1
		return common.Point{
			X: g.size*3 - 1,
			Y: g.size*3 - from.Y - 1,
		}, LEFT
	}

	panic("from out of cube")
}

func (g *Grid) nextPointAroundCubeUpTest(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFaceTest(from)

	switch currentFace {
	case 1:
		// new face 2
		return common.Point{
			X: g.size*3 - from.X - 1,
			Y: g.size,
		}, DOWN
	case 2:
		// new face 1
		return common.Point{
			X: g.size*3 - from.X - 1,
			Y: 0,
		}, DOWN
	case 3:
		// new face 1
		return common.Point{
			X: g.size * 2,
			Y: from.X - g.size,
		}, RIGHT
	case 4:
		// new face 1
		return common.Point{
			X: from.X,
			Y: from.Y - 1,
		}, UP
	case 5:
		// new face 4
		return common.Point{
			X: from.X,
			Y: from.Y - 1,
		}, UP
	case 6:
		// new face 4
		return common.Point{
			X: g.size*3 - 1,
			Y: g.size*5 - from.X - 1,
		}, LEFT
	}

	panic("from out of cube")
}

func (g *Grid) nextPointAroundCubeDownTest(from common.Point) (common.Point, Direction) {
	currentFace := g.GetFaceTest(from)

	switch currentFace {
	case 1:
		// new face 4
		return common.Point{
			X: from.X,
			Y: from.Y + 1,
		}, DOWN
	case 2:
		// new face 5
		return common.Point{
			X: g.size*3 - from.X - 1,
			Y: g.size*3 - 1,
		}, UP
	case 3:
		// new face 5
		return common.Point{
			X: g.size * 2,
			Y: g.size*4 - from.X - 1,
		}, RIGHT
	case 4:
		// new face 5
		return common.Point{
			X: from.X,
			Y: from.Y + 1,
		}, DOWN
	case 5:
		// new face 2
		return common.Point{
			X: g.size*3 - from.X - 1,
			Y: g.size*2 - 1,
		}, UP
	case 6:
		// new face 2
		return common.Point{
			X: 0,
			Y: g.size*5 - from.X - 1,
		}, RIGHT
	}

	panic("from out of cube")
}
