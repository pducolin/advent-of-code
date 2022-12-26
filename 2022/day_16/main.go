package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"sort"
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

// both parts inspired by @bhosale-ajay
func part1(data string) string {
	valvesByName := map[string]Valve{}

	for _, line := range strings.Split(data, "\n") {
		valve := ParseValve(line)
		valvesByName[valve.name] = valve
	}

	timeMap, interestingValves := BuiltTimeMap(valvesByName)

	ret := findBestFlow(valvesByName, timeMap, "AA", 30, interestingValves)

	return strconv.Itoa(ret)
}

func part2(data string) string {
	valvesByName := map[string]Valve{}

	for _, line := range strings.Split(data, "\n") {
		valve := ParseValve(line)
		valvesByName[valve.name] = valve
	}

	timeMap, interestingValves := BuiltTimeMap(valvesByName)

	bestFlowByPath := map[string]int{}
	buildBestFlowByPath(bestFlowByPath, valvesByName, interestingValves, timeMap, "AA", 26, map[string]struct{}{}, 0)
	// build elephant possible paths
	elephantInterestingValves := []string{}
	for valve := range interestingValves {
		if valve == "AA" {
			continue
		}
		elephantInterestingValves = append(elephantInterestingValves, valve)
	}
	sort.Strings(elephantInterestingValves)
	extendBestFlowByPath(bestFlowByPath, elephantInterestingValves)
	result := 0
	for humanPathKey := range bestFlowByPath {
		humanValves := pathKeyToValves(humanPathKey)
		elephantPathKey := ""
		for _, elephantValve := range elephantInterestingValves {
			if _, found := humanValves[elephantValve]; found {
				continue
			}
			elephantPathKey += elephantValve
		}
		flow := bestFlowByPath[humanPathKey] + bestFlowByPath[elephantPathKey]
		if result < flow {
			result = flow
		}
	}

	return strconv.Itoa(result)
}

const regexStr = `Valve ([A-Z]{2}) has flow rate=(\d+); tunnel(s)? lead(s)? to valve(s)? (([A-Z]{2},? ?)+)`

type Valve struct {
	name            string
	flowRate        int
	connectedValves []string
}

func ParseValve(line string) Valve {
	r := regexp.MustCompile(regexStr)

	elements := r.FindStringSubmatch(line)

	currentValve := elements[1]
	flowRate, err := strconv.Atoi(elements[2])
	if err != nil {
		panic(err)
	}
	connectedValves := strings.Split(elements[6], ", ")

	return Valve{
		name:            currentValve,
		flowRate:        flowRate,
		connectedValves: connectedValves,
	}
}

type DistanceMap map[string]int
type TimeMap map[string]DistanceMap

// build map of distances between all valves
func BuiltTimeMap(valvesByName map[string]Valve) (timeMap TimeMap, interestingValves map[string]struct{}) {
	timeMap = TimeMap{}
	interestingValves = map[string]struct{}{}
	for valveName, valve := range valvesByName {
		if valve.flowRate == 0 {
			continue
		}
		interestingValves[valveName] = struct{}{}
	}
	// add starting point
	interestingValves["AA"] = struct{}{}
	for valveName := range interestingValves {
		valveToProcess := common.NewQueue[string]()
		valveToProcess.Push(valveName)
		distanceMap := DistanceMap{valveName: 0}
		visited := map[string]struct{}{}
		visited[valveName] = struct{}{}
		for !valveToProcess.IsEmpty() {
			valve := valveToProcess.Pop()
			for _, connectedValve := range valvesByName[valve].connectedValves {
				if _, found := visited[connectedValve]; !found {
					visited[connectedValve] = struct{}{}
					distanceMap[connectedValve] = distanceMap[valve] + 1
					valveToProcess.Push(connectedValve)
				}
			}
		}
		timeMap[valveName] = distanceMap
	}

	return timeMap, interestingValves
}

func findBestFlow(valvesByName map[string]Valve, timeMap TimeMap, fromValve string, time int, otherValves map[string]struct{}) int {
	possibleValves := map[string]struct{}{}
	// filter out starting point
	for valve := range otherValves {
		if valve == fromValve {
			continue
		}
		possibleValves[valve] = struct{}{}
	}

	bestFlow := 0
	for nextValve := range possibleValves {
		timeLeftAfterVisitingNext := time - timeMap[fromValve][nextValve] - 1
		if timeLeftAfterVisitingNext > 0 {
			flow := valvesByName[nextValve].flowRate*timeLeftAfterVisitingNext +
				findBestFlow(valvesByName, timeMap, nextValve, timeLeftAfterVisitingNext, possibleValves)
			if flow > bestFlow {
				bestFlow = flow
			}
		}
	}
	return bestFlow
}

func buildBestFlowByPath(
	bestFlowByPath map[string]int,
	valvesByName map[string]Valve,
	interestingValves map[string]struct{},
	timeMap TimeMap,
	fromValve string,
	time int,
	visitedValves map[string]struct{},
	pathFlow int,
) int {
	possibleValves := []string{}
	for valve := range interestingValves {
		if valve == "AA" {
			continue
		}
		if _, found := visitedValves[valve]; found {
			continue
		}
		possibleValves = append(possibleValves, valve)
	}

	visitedValveNames := []string{}
	for valve := range visitedValves {
		visitedValveNames = append(visitedValveNames, valve)
	}
	sort.Strings(visitedValveNames)
	pathKey := strings.Join(visitedValveNames, "")
	if _, found := bestFlowByPath[pathKey]; !found {
		bestFlowByPath[pathKey] = 0
	}
	if bestFlowByPath[pathKey] < pathFlow {
		bestFlowByPath[pathKey] = pathFlow
	}
	bestFlow := 0
	for _, nextValve := range possibleValves {
		timeLeft := time - timeMap[fromValve][nextValve] - 1
		if timeLeft > 0 {
			flow := valvesByName[nextValve].flowRate * timeLeft
			newVisited := map[string]struct{}{}
			for _, visitedValve := range visitedValveNames {
				newVisited[visitedValve] = struct{}{}
			}
			newVisited[nextValve] = struct{}{}
			flow = buildBestFlowByPath(bestFlowByPath,
				valvesByName,
				interestingValves,
				timeMap,
				nextValve,
				timeLeft,
				newVisited,
				flow+pathFlow)
			if flow > bestFlow {
				bestFlow = flow
			}
		}
	}
	return bestFlow
}

// no idea what this does
func extendBestFlowByPath(bestFlowByPath map[string]int, valvesInPath []string) int {
	pathKey := strings.Join(valvesInPath, "")
	if _, found := bestFlowByPath[pathKey]; !found {
		bestFlow := 0
		for _, valve := range valvesInPath {
			remainingValves := []string{}
			for _, remainingValve := range valvesInPath {
				if valve == remainingValve {
					continue
				}
				remainingValves = append(remainingValves, remainingValve)
			}
			flow := extendBestFlowByPath(bestFlowByPath, remainingValves)
			if flow > bestFlow {
				bestFlow = flow
			}
		}
		bestFlowByPath[pathKey] = bestFlow
	}

	return bestFlowByPath[pathKey]
}

func pathKeyToValves(pathKey string) map[string]struct{} {
	valves := map[string]struct{}{}
	for i := 0; i < len(pathKey); i += 2 {
		valves[pathKey[i:i+2]] = struct{}{}
	}
	return valves
}
