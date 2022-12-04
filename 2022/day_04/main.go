package main

import (
	_ "embed"
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

type sectionAssignment struct {
	start int
	end   int
}

func part1(data string) string {
	res := 0
	for _, line := range strings.Split(data, "\n") {
		segments := strings.Split(line, ",")
		elfA := newSections(segments[0])
		elfB := newSections(segments[1])

		if elfA.contains(elfB) || elfB.contains(elfA) {
			res += 1
		}
	}
	return fmt.Sprint(res)
}

func part2(data string) string {
	res := 0
	for _, line := range strings.Split(data, "\n") {
		segments := strings.Split(line, ",")
		elfA := newSections(segments[0])
		elfB := newSections(segments[1])

		if elfA.overlap(elfB) {
			res += 1
		}
	}
	return fmt.Sprint(res)
}

func newSections(s string) sectionAssignment {
	assignments := []int{}
	for _, a := range strings.Split(s, "-") {
		n, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		assignments = append(assignments, n)
	}
	return sectionAssignment{
		start: assignments[0],
		end:   assignments[1],
	}
}

// contains
// ...startA.................endA....
// ...........startB...endB..........
func (s sectionAssignment) contains(other sectionAssignment) bool {
	if s.start <= other.start && s.end >= other.end {
		return true
	}
	return false
}

// disjoint
// ...startA...........endA........
// ............................startB........endB...
// OR
// ....................startA....endA........
// ..startB......endB........................
func (s sectionAssignment) disjoint(other sectionAssignment) bool {
	if s.start < other.start && s.end < other.start {
		return true
	}
	if s.start > other.end && s.end > other.end {
		return true
	}
	return false
}

func (s sectionAssignment) overlap(other sectionAssignment) bool {
	return !s.disjoint(other)
}
