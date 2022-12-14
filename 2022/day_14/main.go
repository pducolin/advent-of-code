package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
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
	caveMap := NewCaveMap(data)

	sandCounter := 0
	for {
		err := caveMap.DropSand()
		if err != nil {
			break
		}
		sandCounter++
	}

	return strconv.Itoa(sandCounter)
}

func part2(data string) string {
	caveMap := NewCaveMap(data)

	caveMap.YMax += 2
	caveMap.XMin = 0
	caveMap.XMax *= 2

	for x := caveMap.XMin; x <= caveMap.XMax; x++ {
		caveMap.Points[Point{X: x, Y: caveMap.YMax}] = '#'
	}

	sandCounter := 0
	for {
		err := caveMap.DropSand()
		if err != nil {
			break
		}
		sandCounter++
	}

	return strconv.Itoa(sandCounter)
}

type Point struct {
	X int
	Y int
}

type CaveMap struct {
	Points map[Point]rune
	XMin   int
	XMax   int
	YMax   int
}

func NewCaveMap(data string) (caveMap *CaveMap) {
	caveMap = &CaveMap{
		Points: map[Point]rune{},
		XMax:   0,
		XMin:   math.MaxInt,
		YMax:   0,
	}

	for _, line := range strings.Split(data, "\n") {
		rocks := strings.Split(line, " -> ")

		// add first rock
		rockPoint := parseRockPoint(rocks[0])
		caveMap.Points[rockPoint] = '#'
		caveMap.UpdateBoundaries(rockPoint)

		for _, rock := range rocks[1:] {
			nextRockPoint := parseRockPoint(rock)

			if nextRockPoint.X == rockPoint.X {
				// move vertically
				extremes := []int{nextRockPoint.Y, rockPoint.Y}
				sort.Ints(extremes)
				for y := extremes[0]; y <= extremes[1]; y++ {
					caveMap.Points[Point{X: nextRockPoint.X, Y: y}] = '#'
				}
			} else {
				// move horizontally
				extremes := []int{nextRockPoint.X, rockPoint.X}
				sort.Ints(extremes)
				for x := extremes[0]; x <= extremes[1]; x++ {
					caveMap.Points[Point{Y: nextRockPoint.Y, X: x}] = '#'
				}
			}

			rockPoint = nextRockPoint
			caveMap.UpdateBoundaries(rockPoint)
		}
	}
	return caveMap
}

func parseRockPoint(rock string) Point {
	rockPoint := strings.Split(rock, ",")
	x, err := strconv.Atoi(rockPoint[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(rockPoint[1])
	if err != nil {
		panic(err)
	}
	return Point{X: x, Y: y}
}

func (caveMap *CaveMap) UpdateBoundaries(point Point) {
	if point.X < caveMap.XMin {
		caveMap.XMin = point.X
	}
	if point.X > caveMap.XMax {
		caveMap.XMax = point.X
	}
	if point.Y > caveMap.YMax {
		caveMap.YMax = point.Y
	}
}

func (caveMap *CaveMap) Print() {
	for y := 0; y <= caveMap.YMax; y++ {
		s := ""
		for x := caveMap.XMin; x <= caveMap.XMax; x++ {
			point := Point{X: x, Y: y}
			if r, found := caveMap.Points[point]; found {
				s += string(r)
			} else {
				s += "."
			}
		}
		fmt.Println(s)
	}
}

func (caveMap *CaveMap) DropSand() (err error) {
	sandPoint := Point{X: 500, Y: 0}

	outOfBoundaryErr := errors.New("out of boundary")

	if _, found := caveMap.Points[sandPoint]; found {
		return errors.New("no more space")
	}

	for {
		// try move down
		downPoint := Point{X: sandPoint.X, Y: sandPoint.Y + 1}

		if downPoint.Y > caveMap.YMax {
			return outOfBoundaryErr
		}

		if _, found := caveMap.Points[downPoint]; !found {
			// we can move down, continue
			sandPoint = downPoint
			continue
		}
		// try move down left
		isOutOfBoundariesLeft := false
		downLeftPoint := Point{X: sandPoint.X - 1, Y: sandPoint.Y + 1}

		if downLeftPoint.X < caveMap.XMin || downLeftPoint.Y > caveMap.YMax {
			isOutOfBoundariesLeft = true
		} else {
			if _, found := caveMap.Points[downLeftPoint]; !found {
				// we can move down, continue
				sandPoint = downLeftPoint
				continue
			}
		}
		// try move down right
		downRightPoint := Point{X: sandPoint.X + 1, Y: sandPoint.Y + 1}
		if downRightPoint.X > caveMap.XMax || downRightPoint.Y > caveMap.YMax {
			return outOfBoundaryErr
		}
		if _, found := caveMap.Points[downRightPoint]; !found {
			// we can move down, continue
			sandPoint = downRightPoint
			continue
		}
		if isOutOfBoundariesLeft {
			return outOfBoundaryErr
		}

		// cannot move, set sand point and return
		caveMap.Points[sandPoint] = 'o'
		return nil
	}
}
