package main

import (
	_ "embed"
	"errors"
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
	jp := ParseJetPattern(data)
	chamber := NewChamber()
	chamber.Play(jp, 2022)
	return strconv.FormatInt(chamber.height+chamber.additionalHeight, 10)
}

// from https://github.com/RascalTwo/AdventOfCode/blob/master/2022/solutions/17/solve.ts
func part2(data string) string {
	jp := ParseJetPattern(data)
	chamber := NewChamber()
	chamber.Play(jp, 1000000000000)
	return strconv.FormatInt(chamber.height+chamber.additionalHeight, 10)
}

type Point struct {
	x int64
	y int64
}

type Tile struct {
	Points [][]rune
}

var TILES = []Tile{
	{
		Points: [][]rune{
			{'#', '#', '#', '#'},
		},
	},
	{
		Points: [][]rune{
			{'.', '#', '.'},
			{'#', '#', '#'},
			{'.', '#', '.'},
		},
	},
	{
		Points: [][]rune{
			{'.', '.', '#'},
			{'.', '.', '#'},
			{'#', '#', '#'},
		},
	},
	{
		Points: [][]rune{
			{'#'},
			{'#'},
			{'#'},
			{'#'},
		},
	},
	{
		Points: [][]rune{
			{'#', '#'},
			{'#', '#'},
		},
	},
}

var CHARS = []string{"🔴", "🟡", "🟢", "🔵", "🟣"}

const EMPTY_RUNE = "🔘"

type Rock struct {
	tile         Tile
	width        int
	height       int
	topLeftPoint Point
	char         string
}

func NewRock(index int, chamberHeight int64) *Rock {
	tile := TILES[index]
	rock := Rock{
		tile:   tile,
		width:  len(tile.Points[0]),
		height: len(tile.Points),
	}
	rock.topLeftPoint = Point{
		x: 2,
		y: chamberHeight + 3 + int64(rock.height),
	}
	rock.char = CHARS[index]
	return &rock
}

func (r *Rock) AllPoints() []Point {
	points := []Point{}
	for i, row := range r.tile.Points {
		for j, rune := range row {
			if rune == '#' {
				points = append(points, Point{x: r.topLeftPoint.x + int64(j), y: r.topLeftPoint.y - int64(i)})
			}
		}
	}
	return points
}

type JetPattern struct {
	pattern      string
	currentJet   byte
	currentIndex int
}

func ParseJetPattern(data string) JetPattern {
	return JetPattern{
		pattern:      data,
		currentIndex: 0,
		currentJet:   data[0],
	}
}

func (jp *JetPattern) Next() byte {
	jp.currentIndex = (jp.currentIndex + 1) % len(jp.pattern)
	jp.currentJet = jp.pattern[jp.currentIndex]
	return jp.currentJet
}

func (jp *JetPattern) Reset() {
	jp.currentIndex = 0
	jp.currentJet = jp.pattern[0]
}

type Move struct {
	rockChar string
	x        int64
}

type Chamber struct {
	rocksCount       int64
	linesBottomUp    map[Point]string
	height           int64
	currentRockIndex int
	currentRock      *Rock
	emptyRow         string
	additionalHeight int64
	maxLines         int64
	moves            []Move
}

func NewChamber() *Chamber {
	emptyRow := strings.Repeat(string(EMPTY_RUNE), 7)
	chamber := Chamber{
		linesBottomUp:    map[Point]string{},
		rocksCount:       0,
		height:           0,
		currentRockIndex: 0,
		currentRock:      NewRock(0, 0),
		emptyRow:         emptyRow,
		additionalHeight: 0,
		maxLines:         5000,
		moves:            []Move{},
	}

	// bottom row
	for x := 0; x < 7; x++ {
		chamber.linesBottomUp[Point{x: int64(x), y: 0}] = "🖤"
	}

	// empty rows
	for y := 1; y < 5; y++ {
		for x := 0; x < 7; x++ {
			chamber.linesBottomUp[Point{x: int64(x), y: int64(y)}] = EMPTY_RUNE
		}
	}

	return &chamber
}

func (c *Chamber) Print(numberOfLines int64) []string {
	ret := []string{}
	start := c.height
	bottom := start - numberOfLines
	if bottom < 0 {
		bottom = 0
	}
	for y := start; y >= bottom; y-- {
		line := ""
		for x := 0; x < 7; x++ {
			r := c.linesBottomUp[Point{x: int64(x), y: y}]
			line += r
		}
		ret = append(ret, fmt.Sprintf("|%s|", line))
	}

	for _, line := range ret {
		fmt.Println(line)
	}

	return ret
}

func (c *Chamber) NextRockIndex() int {
	c.currentRockIndex = (c.currentRockIndex + 1) % len(TILES)
	return c.currentRockIndex
}

func (c *Chamber) MoveLeft() {
	points := c.currentRock.AllPoints()
	for _, point := range points {
		newPoint := Point{
			x: point.x - 1,
			y: point.y,
		}
		if newPoint.x < 0 {
			return
		}
		if c.linesBottomUp[newPoint] != EMPTY_RUNE {
			return
		}
	}
	c.currentRock.topLeftPoint.x -= 1
}

func (c *Chamber) MoveRight() {
	points := c.currentRock.AllPoints()
	for _, point := range points {
		newPoint := Point{
			x: point.x + 1,
			y: point.y,
		}
		if newPoint.x >= 7 {
			return
		}
		if c.linesBottomUp[newPoint] != EMPTY_RUNE {
			return
		}
	}

	c.currentRock.topLeftPoint.x += 1
	return
}

func (c *Chamber) MoveDown() error {
	points := c.currentRock.AllPoints()
	for _, point := range points {
		newPoint := Point{
			x: point.x,
			y: point.y - 1,
		}
		if c.linesBottomUp[newPoint] != EMPTY_RUNE {
			return errors.New("reached bottom")
		}
	}
	c.currentRock.topLeftPoint.y -= 1
	return nil
}

type Pattern struct {
	rocksCount int64
	height     int64
}

func (chamber *Chamber) Play(jp JetPattern, totRocks int64) {
	patterns := map[string]Pattern{}
	for chamber.rocksCount < totRocks {
		if chamber.currentRock == nil {
			chamber.NextRockIndex()
			chamber.currentRock = NewRock(chamber.currentRockIndex, chamber.height)
		}

		// move with jet
		if jp.currentJet == '>' {
			// move to the right
			chamber.MoveRight()
		} else {
			// move to the left
			chamber.MoveLeft()
		}
		jp.Next()
		//  move down
		err := chamber.MoveDown()
		if err != nil {
			// add rock to lines
			allPoints := chamber.currentRock.AllPoints()
			for _, point := range allPoints {
				chamber.linesBottomUp[point] = chamber.currentRock.char
			}
			// add rock to list
			chamber.rocksCount++
			if chamber.currentRock.topLeftPoint.y > chamber.height {
				chamber.height = chamber.currentRock.topLeftPoint.y
			}
			chamber.moves = append(chamber.moves, Move{x: chamber.currentRock.topLeftPoint.x, rockChar: chamber.currentRock.char})
			chamber.currentRock = nil
			// add 8 empty rows
			firstEmptyY := chamber.height + 1
			for y := firstEmptyY; y < firstEmptyY+8; y++ {
				for x := 0; x < 7; x++ {
					chamber.linesBottomUp[Point{x: int64(x), y: int64(y)}] = EMPTY_RUNE
				}
			}

			patternKey := chamber.BuildPatternKey(jp)

			if _, found := patterns[patternKey]; found {
				previous := patterns[patternKey]
				rocksChanges := chamber.rocksCount - previous.rocksCount
				highestPointChanges := chamber.height - previous.height
				cycles := (totRocks-previous.rocksCount)/rocksChanges - 1
				chamber.additionalHeight += cycles * highestPointChanges
				chamber.rocksCount += cycles * rocksChanges
				continue
			}
			patterns[patternKey] = Pattern{rocksCount: chamber.rocksCount, height: chamber.height}
		}
	}
}

func (c *Chamber) BuildPatternKey(jp JetPattern) string {
	return fmt.Sprintf("%d|%d|%s",
		jp.currentIndex,
		c.currentRockIndex,
		strings.Join(c.Print(5), "|"))
}
