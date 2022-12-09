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

func part1(data string) string {
	treeMap := parseTreeMap(data)
	res := countVisibleTrees(treeMap)
	return strconv.Itoa(res)
}

func part2(data string) string {
	treeMap := parseTreeMap(data)
	res := findMaxScore(treeMap)
	return strconv.Itoa(res)
}

func parseTreeMap(data string) (treeMap [][]int) {
	lines := strings.Split(data, "\n")
	totRows := len(lines)
	totColumns := len(lines[0])

	treeMap = make([][]int, totRows)
	for i := range treeMap {
		treeMap[i] = make([]int, totColumns)
	}

	for rowIndex, line := range lines {
		for columnIndex, char := range line {
			value := int(char - '0')
			treeMap[rowIndex][columnIndex] = value
		}
	}

	return treeMap
}

func countVisibleTrees(treeMap [][]int) (visibileTreeCount int) {
	visibleTrees := map[string]struct{}{}

	for rowIndex := range treeMap {
		for columnIndex := range treeMap[rowIndex] {

			valueIndex := fmt.Sprintf("%d,%d", rowIndex, columnIndex)
			// check is already visible
			if _, found := visibleTrees[valueIndex]; found {
				continue
			}

			if isVisible(treeMap, rowIndex, columnIndex) {
				visibleTrees[valueIndex] = struct{}{}
			}
		}
	}

	return len(visibleTrees)
}

func isVisible(treeMap [][]int, rowIndex, columnIndex int) bool {
	return visibleFromLeft(treeMap, rowIndex, columnIndex) || visibleFromRight(treeMap, rowIndex, columnIndex) || visibleFromTop(treeMap, rowIndex, columnIndex) || visibleFromBottom(treeMap, rowIndex, columnIndex)
}

func visibleFromLeft(treeMap [][]int, rowIndex, columnIndex int) bool {
	if columnIndex == 0 {
		return true
	}

	currentValue := treeMap[rowIndex][columnIndex]
	for i := 0; i < columnIndex; i++ {
		leftValue := treeMap[rowIndex][i]
		if leftValue >= currentValue {
			return false
		}
	}

	return true
}

func visibleFromRight(treeMap [][]int, rowIndex, columnIndex int) bool {
	totColumns := len(treeMap[0])
	if columnIndex == totColumns-1 {
		return true
	}

	currentValue := treeMap[rowIndex][columnIndex]
	for i := columnIndex + 1; i < totColumns; i++ {
		rightValue := treeMap[rowIndex][i]
		if rightValue >= currentValue {
			return false
		}
	}

	return true
}

func visibleFromTop(treeMap [][]int, rowIndex, columnIndex int) bool {
	if rowIndex == 0 {
		return true
	}

	currentValue := treeMap[rowIndex][columnIndex]
	for i := 0; i < rowIndex; i++ {
		topValue := treeMap[i][columnIndex]
		if topValue >= currentValue {
			return false
		}
	}

	return true
}

func visibleFromBottom(treeMap [][]int, rowIndex, columnIndex int) bool {
	totRows := len(treeMap)

	if rowIndex == totRows-1 {
		return true
	}

	currentValue := treeMap[rowIndex][columnIndex]
	for i := rowIndex + 1; i < totRows; i++ {
		bottomValue := treeMap[i][columnIndex]
		if bottomValue >= currentValue {
			return false
		}
	}

	return true
}

func findMaxScore(treeMap [][]int) (maxScore int) {
	maxScore = 0

	for rowIndex := range treeMap {
		for columnIndex := range treeMap[rowIndex] {
			score := evaluateScore(treeMap, rowIndex, columnIndex)
			if maxScore < score {
				maxScore = score
			}
		}
	}

	return maxScore
}

func evaluateScore(treeMap [][]int, rowIndex, columnIndex int) int {
	return countVisibleToBottom(treeMap, rowIndex, columnIndex) *
		countVisibleToTop(treeMap, rowIndex, columnIndex) *
		countVisibleToLeft(treeMap, rowIndex, columnIndex) *
		countVisibleToRight(treeMap, rowIndex, columnIndex)
}

func countVisibleToLeft(treeMap [][]int, rowIndex, columnIndex int) int {
	if columnIndex == 0 {
		return 0
	}

	currentValue := treeMap[rowIndex][columnIndex]
	res := 0
	for i := columnIndex - 1; i >= 0; i-- {
		res++
		leftValue := treeMap[rowIndex][i]
		if leftValue >= currentValue {
			break
		}
	}

	return res
}

func countVisibleToRight(treeMap [][]int, rowIndex, columnIndex int) int {
	totColumns := len(treeMap[0])
	if columnIndex == totColumns-1 {
		return 0
	}

	currentValue := treeMap[rowIndex][columnIndex]
	res := 0
	for i := columnIndex + 1; i < totColumns; i++ {
		res++
		rightValue := treeMap[rowIndex][i]
		if rightValue >= currentValue {
			break
		}
	}

	return res
}

func countVisibleToTop(treeMap [][]int, rowIndex, columnIndex int) int {
	if rowIndex == 0 {
		return 0
	}

	currentValue := treeMap[rowIndex][columnIndex]
	res := 0
	for i := rowIndex - 1; i >= 0; i-- {
		res++
		topValue := treeMap[i][columnIndex]
		if topValue >= currentValue {
			break
		}
	}

	return res
}

func countVisibleToBottom(treeMap [][]int, rowIndex, columnIndex int) int {
	totRows := len(treeMap)

	if rowIndex == totRows-1 {
		return 0
	}

	currentValue := treeMap[rowIndex][columnIndex]
	res := 0
	for i := rowIndex + 1; i < totRows; i++ {
		res++
		bottomValue := treeMap[i][columnIndex]
		if bottomValue >= currentValue {
			break
		}
	}

	return res
}
