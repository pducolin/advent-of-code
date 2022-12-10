package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pducolin/advent-of-code/2022/2022/day_07/filesystem"
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

const CMD_BACK = "$ cd .."
const CMD_LIST = "$ ls"
const CMD_IN_REGEX = "\\$\\ cd\\ (\\w+)"
const LS_DIR_REGEX = "dir\\ (\\w+)"
const LS_FILE_REGEX = "(\\d+)\\ (.+)"

const MAX_SIZE = 100000

const TOTAL_SIZE = 70000000
const MIN_FREE_SIZE = 30000000

var (
	dirRule, _   = regexp.Compile(LS_DIR_REGEX)
	fileRule, _  = regexp.Compile(LS_FILE_REGEX)
	cmdInRule, _ = regexp.Compile(CMD_IN_REGEX)
)

func part1(data string) string {
	rootDir := parseFileSystem(data)

	res := 0
	dirToEvaluate := []*filesystem.Directory{rootDir}

	for len(dirToEvaluate) > 0 {
		currentDir := dirToEvaluate[0]
		if len(dirToEvaluate) > 1 {
			dirToEvaluate = dirToEvaluate[1:]
		} else {
			dirToEvaluate = []*filesystem.Directory{}
		}

		if currentDir.GetOrEvaluateSize() < MAX_SIZE {
			res += currentDir.GetOrEvaluateSize()
		}

		for _, subdir := range currentDir.Subdirectories {
			dirToEvaluate = append(dirToEvaluate, subdir)
		}
	}

	return fmt.Sprint(res)
}

func part2(data string) string {
	rootDir := parseFileSystem(data)

	dirToEvaluate := []*filesystem.Directory{rootDir}
	dirSizes := []int{}

	availableFreeSize := TOTAL_SIZE - rootDir.GetOrEvaluateSize()
	sizeToFreeUp := MIN_FREE_SIZE - availableFreeSize

	for len(dirToEvaluate) > 0 {
		currentDir := dirToEvaluate[0]
		if len(dirToEvaluate) > 1 {
			dirToEvaluate = dirToEvaluate[1:]
		} else {
			dirToEvaluate = []*filesystem.Directory{}
		}

		if currentDir.GetOrEvaluateSize() >= sizeToFreeUp {
			dirSizes = append(dirSizes, currentDir.GetOrEvaluateSize())
		}

		for _, subdir := range currentDir.Subdirectories {
			dirToEvaluate = append(dirToEvaluate, subdir)
		}
	}

	sort.Ints(dirSizes)

	return fmt.Sprint(dirSizes[0])
}

func parseFileSystem(data string) *filesystem.Directory {
	currentDir := filesystem.NewDirectory("/", nil)
	for _, line := range strings.Split(data, "\n") {
		if line == CMD_BACK {
			if currentDir.Parent == nil {
				panic(errors.New("back from root"))
			}
			currentDir = currentDir.Parent
			continue
		}

		if cmdInRule.MatchString(line) {
			targetDirName := cmdInRule.FindStringSubmatch(line)[1]
			currentDir = currentDir.Subdirectories[targetDirName]
			continue
		}

		if line == CMD_LIST {
			continue
		}

		if dirRule.MatchString(line) {
			subdirName := dirRule.FindStringSubmatch(line)[1]
			currentDir.AddSubdirectory(subdirName)
			continue
		}

		if fileRule.MatchString(line) {
			fileInfo := strings.Split(line, " ")
			fileSize, err := strconv.Atoi(fileInfo[0])
			if err != nil {
				panic(err)
			}
			currentDir.AddFile(filesystem.NewFile(fileInfo[1], fileSize))
			continue
		}
	}

	// pop back
	for {
		if currentDir.Parent == nil {
			break
		}
		currentDir = currentDir.Parent
	}

	// and return root
	return currentDir
}
