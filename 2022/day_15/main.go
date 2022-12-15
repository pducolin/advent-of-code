package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"math"
	"regexp"
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
		fmt.Println(part1(inputData, 2000000))
	} else {
		fmt.Println(part2(inputData, 0, 4000000))
	}
}

func part1(data string, y int) string {
	beaconMap := NewBeaconMap(data, y)
	res := beaconMap.CountImpossibleBeaconAt(y)
	return strconv.Itoa(res)
}

// part2 solution inspired by https://github.com/camaron-ai/adventofcode-2022/blob/b55b74b1a0e3d64de8e5674952201d10e8216f86/day15/main.py
func part2(data string, min, max int) string {
	fmt.Println("Start creating map")
	input := parseData(data)
	fmt.Println("Look for beacon")
	res := FindBeacon(input, min, max)
	fmt.Println("Beacon found")
	return strconv.Itoa(res.x*4000000 + res.y)
}

type Point struct {
	x int
	y int
}

func (point *Point) manhattanDistance(other Point) int {
	return abs(point.x-other.x) + abs(point.y-other.y)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type BeaconMap struct {
	Points   map[Point]rune
	Segments []Segment
	xMin     int
	xMax     int
	yMin     int
	yMax     int
}

type SensorBeacon struct {
	sensor Point
	beacon Point
}

func parseData(data string) []SensorBeacon {
	ret := []SensorBeacon{}
	for _, line := range strings.Split(data, "\n") {
		sensor, beacon := parseLine(line)
		ret = append(ret, SensorBeacon{sensor: sensor, beacon: beacon})
	}
	return ret
}

func NewBeaconMap(data string, targetY int) (beaconMap *BeaconMap) {
	beaconMap = &BeaconMap{
		Points: map[Point]rune{},
		xMin:   math.MaxInt,
		xMax:   math.MinInt,
		yMin:   math.MaxInt,
		yMax:   math.MinInt,
	}

	for _, line := range strings.Split(data, "\n") {
		sensor, beacon := parseLine(line)
		beaconMap.Points[sensor] = 'S'
		beaconMap.Points[beacon] = 'B'
		// fill map with no beacon points
		distance := sensor.manhattanDistance(beacon)
		if targetY >= sensor.y-distance && targetY <= sensor.y+distance {
			for x := sensor.x - distance; x <= sensor.x+distance; x++ {
				point := Point{x: x, y: targetY}
				if point.manhattanDistance(sensor) > distance {
					continue
				}
				if _, found := beaconMap.Points[point]; found {
					continue
				}
				beaconMap.Points[point] = '#'
			}
		}
		// update xMin, xMax
		if beaconMap.xMin > sensor.x-distance {
			beaconMap.xMin = sensor.x - distance
		}
		if beaconMap.xMax < sensor.x+distance {
			beaconMap.xMax = sensor.x + distance
		}
		// update yMin, yMax
		if beaconMap.yMin > sensor.y-distance {
			beaconMap.yMin = sensor.y - distance
		}
		if beaconMap.yMax < sensor.y+distance {
			beaconMap.yMax = sensor.y + distance
		}
	}

	return beaconMap
}

type Segment struct {
	xMin int
	xMax int
}

func evaluateIntervalsAtY(input []SensorBeacon, min, max, y int) []Segment {
	intervals := []Segment{}
	for _, sensorBeacon := range input {
		sensor := sensorBeacon.sensor
		beacon := sensorBeacon.beacon
		distanceFromBeacon := sensor.manhattanDistance(beacon)
		distanceFromY := sensor.manhattanDistance(Point{sensor.x, y})
		if distanceFromY <= distanceFromBeacon {
			// interval of beacons
			delta := distanceFromBeacon - distanceFromY
			newInterval := Segment{
				xMin: sensorBeacon.sensor.x - delta,
				xMax: sensorBeacon.sensor.x + delta,
			}

			if newInterval.xMin <= min && newInterval.xMax >= max {
				return []Segment{newInterval}
			}

			intervals = append(intervals, newInterval)
		}
	}
	return intervals
}

func FindBeacon(input []SensorBeacon, min, max int) Point {
	for y := min; y <= max; y++ {
		intervals := evaluateIntervalsAtY(input, min, max, y)
		// sort intervals
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].xMin < intervals[j].xMin
		})

		interval := intervals[0]
		disjointIntervals := []Segment{}
		for _, otherInterval := range intervals {
			if interval.xMax >= otherInterval.xMin {
				if interval.xMax < otherInterval.xMax {
					interval.xMax = otherInterval.xMax
				}
				continue
			}
			disjointIntervals = append(disjointIntervals, interval)
			interval = otherInterval
		}
		disjointIntervals = append(disjointIntervals, interval)
		if len(disjointIntervals) > 1 {
			x := disjointIntervals[0].xMax + 1
			fmt.Printf("Found point at [%d,%d]", x, y)
			return Point{x: x, y: y}
		}
	}
	panic(errors.New("no spot found"))
}

func (beaconMap *BeaconMap) CountImpossibleBeaconAt(y int) int {
	counter := 0
	for x := beaconMap.xMin; x <= beaconMap.xMax; x++ {
		if r, found := beaconMap.Points[Point{x: x, y: y}]; found {
			if r == '#' {
				counter += 1
			}
		}
	}
	return counter
}

func (beaconMap *BeaconMap) FindBeacon(min, max int) Point {
	for x := min; x <= max; x++ {
		for y := min; y <= max; y++ {
			point := Point{x: x, y: y}
			if _, found := beaconMap.Points[point]; !found {
				return point
			}
		}
	}
	panic(errors.New("no spot found"))
}

const lineRegExpStr = `Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`

func parseLine(line string) (sensorPoint Point, beaconPoint Point) {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	r, err := regexp.Compile(lineRegExpStr)
	if err != nil {
		panic(err)
	}

	matches := r.FindStringSubmatch(line)

	nums := []int{}

	for _, numStr := range matches[1:] {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return Point{x: nums[0], y: nums[1]}, Point{x: nums[2], y: nums[3]}
}

func (beaconMap *BeaconMap) Print() {
	for y := beaconMap.yMin; y <= beaconMap.yMax; y++ {
		s := ""
		for x := beaconMap.xMin; x <= beaconMap.xMax; x++ {
			if r, found := beaconMap.Points[Point{x: x, y: y}]; found {
				s += string(r)
				continue
			}
			s += "."
		}
		fmt.Println(s)
	}
}
