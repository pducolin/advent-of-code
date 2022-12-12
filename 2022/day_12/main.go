package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"math"
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
	heightmap, startPosition, targetPosition := parseMap(data)

	currentPosition := Point{row: startPosition.row, column: startPosition.column}
	visited := map[string]struct{}{}
	steps, err := countMinSteps(heightmap, currentPosition, targetPosition, visited)

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(steps)
}

func part2(data string) string {
	return "Hello World !"
}

type Point struct {
	row    int
	column int
}

func (point Point) Equal(other Point) bool {
	return point.row == other.row && point.column == other.column
}

func (point Point) ToString() string {
	return fmt.Sprintf("%d,%d", point.row, point.column)
}

func (point Point) Up() Point {
	return Point{row: point.row - 1, column: point.column}
}

func (point Point) Down() Point {
	return Point{row: point.row + 1, column: point.column}
}

func (point Point) Left() Point {
	return Point{row: point.row, column: point.column - 1}
}

func (point Point) Right() Point {
	return Point{row: point.row, column: point.column + 1}
}

func (point Point) CanMoveTo(heightmap [][]int, nextPoint Point, visited map[string]struct{}) bool {
	if !nextPoint.IsWithinMap(heightmap) {
		return false
	}

	if _, found := visited[nextPoint.ToString()]; found {
		return false
	}

	return heightmap[nextPoint.row][nextPoint.column]-heightmap[point.row][point.column] <= 1
}

func (point Point) IsWithinMap(heightmap [][]int) bool {
	return point.row >= 0 && point.row < len(heightmap) &&
		point.column >= 0 && point.column < len(heightmap[0])
}

func parseMap(data string) (heightmap [][]int, currentPosition, targetPosition Point) {
	heightmap = [][]int{}
	for rowIndex, line := range strings.Split(data, "\n") {
		heightmap = append(heightmap, []int{})
		for colIndex, value := range line {
			if value == 'S' {
				currentPosition = Point{row: rowIndex, column: colIndex}
				heightmap[rowIndex] = append(heightmap[rowIndex], 0)
				continue
			}

			if value == 'E' {
				targetPosition = Point{row: rowIndex, column: colIndex}
				heightmap[rowIndex] = append(heightmap[rowIndex], int('z'-'a'))
				continue
			}

			heightmap[rowIndex] = append(heightmap[rowIndex], int(value-'a'))
		}
	}
	return heightmap, currentPosition, targetPosition
}

func countMinSteps(heightmap [][]int, from Point, to Point, visited map[string]struct{}) (minSteps int, err error) {
	if !from.IsWithinMap(heightmap) {
		return -1, errors.New("out of map")
	}

	visited[from.ToString()] = struct{}{}
	minSteps = math.MaxInt
	err = fmt.Errorf("no way to target from %s", from.ToString())
	// move up, down, left, right
	moves := []Point{from.Up(), from.Down(), from.Left(), from.Right()}

	for _, nextPoint := range moves {
		if from.CanMoveTo(heightmap, nextPoint, visited) {
			steps, stepErr := countMinSteps(heightmap, nextPoint, to, visited)
			if stepErr == nil && steps+1 < minSteps {
				minSteps = steps + 1
			}
		}
	}

	return minSteps, err
}
