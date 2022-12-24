package main

import (
	_ "embed"
	"flag"
	"fmt"
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
		fmt.Println(part1(inputData))
	} else {
		fmt.Println(part2(inputData))
	}
}

func part1(data string) string {
	lines := strings.Split(data, "\n")
	resultsChannel := make(chan int, len(lines))

	worker := func(line string, result chan int) {
		bp := NewBlueprint(line)
		result <- bp.id * bp.FindMaxGeodes(24)
	}

	for _, line := range lines {
		go worker(line, resultsChannel)
	}

	totQualityLevel := 0
	count := 0
	for count < len(lines) {
		res := <-resultsChannel
		totQualityLevel += res
		count++
	}
	return strconv.Itoa(totQualityLevel)
}

func part2(data string) string {
	lines := strings.Split(data, "\n")
	if len(lines) > 3 {
		lines = lines[0:3]
	}
	resultsChannel := make(chan int, len(lines))

	worker := func(line string, result chan int) {
		bp := NewBlueprint(line)
		result <- bp.FindMaxGeodes(32)
	}

	for _, line := range lines {
		go worker(line, resultsChannel)
	}

	totQualityLevel := 1
	count := 0
	for count < len(lines) {
		res := <-resultsChannel
		totQualityLevel *= res
		count++
	}
	return strconv.Itoa(totQualityLevel)
}

type Material int

const (
	ore Material = iota
	clay
	obsidian
	geode
)

var MATERIALS = []Material{ore, clay, obsidian, geode}

type RobotCost struct {
	material Material
	cost     map[Material]int
}

func NewRobotCost(material Material) RobotCost {
	return RobotCost{
		material: material,
		cost:     map[Material]int{},
	}
}

type Blueprint struct {
	id                        int
	robotCostByTargetMaterial map[Material]RobotCost
}

const blueprintRegex = `Blueprint (\d+): Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\.`

func NewBlueprint(s string) Blueprint {
	bp := Blueprint{
		robotCostByTargetMaterial: map[Material]RobotCost{},
	}
	r := regexp.MustCompile(blueprintRegex)

	groups := r.FindStringSubmatch(s)

	id, err := strconv.Atoi(groups[1])
	if err != nil {
		panic(err)
	}
	bp.id = id

	robotCost := NewRobotCost(ore)
	robotCost.cost[ore], err = strconv.Atoi(groups[2])
	if err != nil {
		panic(err)
	}
	bp.robotCostByTargetMaterial[ore] = robotCost

	robotCost = NewRobotCost(clay)
	robotCost.cost[ore], err = strconv.Atoi(groups[3])
	if err != nil {
		panic(err)
	}
	bp.robotCostByTargetMaterial[clay] = robotCost

	robotCost = NewRobotCost(obsidian)
	robotCost.cost[ore], err = strconv.Atoi(groups[4])
	if err != nil {
		panic(err)
	}
	robotCost.cost[clay], err = strconv.Atoi(groups[5])
	if err != nil {
		panic(err)
	}
	bp.robotCostByTargetMaterial[obsidian] = robotCost

	robotCost = NewRobotCost(geode)
	robotCost.cost[ore], err = strconv.Atoi(groups[6])
	if err != nil {
		panic(err)
	}
	robotCost.cost[obsidian], err = strconv.Atoi(groups[7])
	if err != nil {
		panic(err)
	}
	bp.robotCostByTargetMaterial[geode] = robotCost

	return bp
}

func (bp *Blueprint) TotNeedByMaterial(material Material) int {
	res := 0
	for _, robotCost := range bp.robotCostByTargetMaterial {
		res += robotCost.cost[material]
	}
	return res
}

type RobotState struct {
	time               int
	robotsByMaterial   map[Material]int
	resourceByMaterial map[Material]int
}

func (rs *RobotState) NextState() RobotState {
	nextState := RobotState{
		time:               rs.time + 1,
		robotsByMaterial:   map[Material]int{},
		resourceByMaterial: map[Material]int{},
	}

	for k, v := range rs.robotsByMaterial {
		nextState.robotsByMaterial[k] = v
	}

	for k, v := range rs.resourceByMaterial {
		nextState.resourceByMaterial[k] = v + rs.robotsByMaterial[k]
	}

	return nextState
}

func (rs *RobotState) HashString() string {
	hash := strconv.Itoa(rs.time)
	for _, material := range MATERIALS {
		hash += fmt.Sprintf("%d:%d", material, rs.resourceByMaterial[material])
		hash += fmt.Sprintf("%d:%d", material, rs.robotsByMaterial[material])
	}
	return hash
}

func (rs *RobotState) Score() int {
	return ((rs.robotsByMaterial[geode]+rs.resourceByMaterial[geode])*10+
		(rs.robotsByMaterial[obsidian]+rs.resourceByMaterial[obsidian]))*10 +
		(rs.robotsByMaterial[clay]+rs.resourceByMaterial[clay])*10 +
		rs.robotsByMaterial[ore] + rs.resourceByMaterial[ore]
}

func (bp *Blueprint) CanBuild(status RobotState, targetMaterial Material) bool {
	if targetMaterial != geode && status.robotsByMaterial[targetMaterial] >= bp.TotNeedByMaterial(targetMaterial) {
		return false
	}
	robotCost := bp.robotCostByTargetMaterial[targetMaterial]
	for resource, cost := range robotCost.cost {
		if status.resourceByMaterial[resource] < cost {
			return false
		}
	}
	return true
}

func (bp *Blueprint) AllNextStates(state RobotState) []RobotState {
	possibleNextStates := []RobotState{}

	if bp.CanBuild(state, geode) {
		nextState := state.NextState()
		robotCost := bp.robotCostByTargetMaterial[geode]
		for material, cost := range robotCost.cost {
			nextState.resourceByMaterial[material] -= cost
		}
		nextState.robotsByMaterial[geode] += 1
		return []RobotState{nextState}
	}

	for _, targetMaterial := range MATERIALS {
		if !bp.CanBuild(state, targetMaterial) {
			continue
		}
		nextState := state.NextState()
		robotCost := bp.robotCostByTargetMaterial[targetMaterial]
		for material, cost := range robotCost.cost {
			nextState.resourceByMaterial[material] -= cost
		}
		nextState.robotsByMaterial[targetMaterial] += 1
		possibleNextStates = append(possibleNextStates, nextState)
	}
	// add no action status
	possibleNextStates = append(possibleNextStates, state.NextState())
	return possibleNextStates
}

const MAX_ITEMS = 10000000

func (bp *Blueprint) FindMaxGeodes(maxTime int) int {
	initState := RobotState{
		time:               0,
		resourceByMaterial: map[Material]int{},
		robotsByMaterial:   map[Material]int{},
	}
	for _, material := range MATERIALS {
		initState.resourceByMaterial[material] = 0
		initState.robotsByMaterial[material] = 0
	}
	initState.robotsByMaterial[ore] = 1

	processed := map[string]struct{}{initState.HashString(): {}}
	states := []RobotState{initState}
	for t := 0; t < maxTime; t++ {
		oldStates := states
		states = []RobotState{}
		for _, state := range oldStates {
			for _, nextState := range bp.AllNextStates(state) {
				if _, found := processed[nextState.HashString()]; found {
					continue
				}
				processed[nextState.HashString()] = struct{}{}
				states = append(states, nextState)
			}
		}
		sort.Slice(states, func(i, j int) bool {
			return states[i].Score() > states[j].Score()
		})
		if len(states) > MAX_ITEMS {
			states = states[0:MAX_ITEMS]
		}
	}
	sort.Slice(states, func(i, j int) bool {
		return states[i].resourceByMaterial[geode] > states[j].resourceByMaterial[geode]
	})

	return states[0].resourceByMaterial[geode]
}
