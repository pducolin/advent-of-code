package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"container/heap"
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
	heightmap, startingPoint, targetPoint := parseMap(data)

	res, err := findShortestPath(heightmap, startingPoint, targetPoint)

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(res)
}

func part2(data string) string {
	heightmap, _, targetPoint := parseMap(data)

	startingPoints := []Point{}
	for point, height := range heightmap {
		if height == 0 {
			startingPoints = append(startingPoints, point)
		}
	}

	pathDistances := []int{}

	for _, start := range startingPoints {
		distance, err := findShortestPath(heightmap, start, targetPoint)
		if err != nil {
			continue
		}
		pathDistances = append(pathDistances, distance)
	}

	sort.Ints(pathDistances)

	return strconv.Itoa(pathDistances[0])
}

type Point struct {
	row    int
	column int
}

func (point *Point) GetNeighbours(heightmap map[Point]int) (neighbours []Point) {
	neighbours = []Point{}
	currentHeight := heightmap[*point]

	// up
	neighbour := Point{row: point.row - 1, column: point.column}
	if neighbourHeight, found := heightmap[neighbour]; found && neighbourHeight-currentHeight <= 1 {
		neighbours = append(neighbours, neighbour)
	}

	// down
	neighbour = Point{row: point.row + 1, column: point.column}
	if neighbourHeight, found := heightmap[neighbour]; found && neighbourHeight-currentHeight <= 1 {
		neighbours = append(neighbours, neighbour)
	}

	// left
	neighbour = Point{row: point.row, column: point.column - 1}
	if neighbourHeight, found := heightmap[neighbour]; found && neighbourHeight-currentHeight <= 1 {
		neighbours = append(neighbours, neighbour)
	}

	// right
	neighbour = Point{row: point.row, column: point.column + 1}
	if neighbourHeight, found := heightmap[neighbour]; found && neighbourHeight-currentHeight <= 1 {
		neighbours = append(neighbours, neighbour)
	}

	return neighbours
}

func parseMap(data string) (heightmap map[Point]int, startingPoint, targetPosition Point) {
	heightmap = map[Point]int{}
	for rowIndex, line := range strings.Split(data, "\n") {
		for colIndex, value := range line {
			point := Point{row: rowIndex, column: colIndex}
			if value == 'S' {
				heightmap[point] = 0
				startingPoint = point
				continue
			}

			if value == 'E' {
				heightmap[point] = int('z' - 'a')
				targetPosition = point
				continue
			}

			heightmap[point] = int(value - 'a')
		}
	}
	return heightmap, startingPoint, targetPosition
}

func findShortestPath(heightmap map[Point]int, startingPoint, target Point) (shortestPath int, err error) {
	visited := map[Point]struct{}{}
	distanceFromStartingPoint := map[Point]int{}

	for point := range heightmap {
		distanceFromStartingPoint[point] = math.MaxInt
	}
	distanceFromStartingPoint[startingPoint] = 0

	priorityQueue := MapPointHeap{MapPoint{
		Point:                 startingPoint,
		CostFromStartingPoint: 0,
	}}

	// heap ensures items in queue are sorted by distance from starting point
	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		currentMapPoint := heap.Pop(&priorityQueue).(MapPoint)

		if _, found := visited[currentMapPoint.Point]; found {
			continue
		}

		if currentMapPoint.Point == target {
			// we are done, this is the shortest path to target
			return currentMapPoint.CostFromStartingPoint, nil
		}

		for _, neighbour := range currentMapPoint.Point.GetNeighbours(heightmap) {
			heap.Push(&priorityQueue, MapPoint{Point: neighbour, CostFromStartingPoint: currentMapPoint.CostFromStartingPoint + 1})
		}

		visited[currentMapPoint.Point] = struct{}{}
	}

	return -1, fmt.Errorf("no path from %#v to %#v", startingPoint, target)
}

// Source https://pkg.go.dev/container/heap

type MapPoint struct {
	Point                 Point
	CostFromStartingPoint int
}

// An MapPointHeap is a min-heap of map nodes.
type MapPointHeap []MapPoint

func (h MapPointHeap) Len() int { return len(h) }
func (h MapPointHeap) Less(i, j int) bool {
	return h[i].CostFromStartingPoint < h[j].CostFromStartingPoint
}
func (h MapPointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MapPointHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(MapPoint))
}

func (h *MapPointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
