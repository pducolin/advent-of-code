package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
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
	packetPairsStr := strings.Split(data, "\n\n")

	res := 0
	for i, pair := range packetPairsStr {
		pairItems := strings.Split(pair, "\n")
		var left, right any
		err := json.Unmarshal([]byte(pairItems[0]), &left)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(pairItems[1]), &right)
		if err != nil {
			panic(err)
		}
		// left := ParsePacket(pairItems[0])
		// right := ParsePacket(pairItems[1])
		if ComparePackets(left, right) < 0 {
			res += i + 1
		}
	}

	return strconv.Itoa(res)
}

func part2(data string) string {
	packetPairsStr := strings.Split(data, "\n\n")

	packetPairsStr = append(packetPairsStr, "[[2]]\n[[6]]")

	packets := []any{}
	var p any
	for _, pair := range packetPairsStr {
		pairItems := strings.Split(pair, "\n")
		for _, item := range pairItems {
			err := json.Unmarshal([]byte(item), &p)
			if err != nil {
				panic(err)
			}
			packets = append(packets, p)
		}
	}

	sort.Slice(packets, func(i, j int) bool {
		return ComparePackets(packets[i], packets[j]) < 0
	})

	indexes := []int{}
	var firstPacket, secondPacket any
	err := json.Unmarshal([]byte("[[2]]"), &firstPacket)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte("[[6]]"), &secondPacket)
	if err != nil {
		panic(err)
	}
	for i, p := range packets {
		if ComparePackets(p, firstPacket) == 0 || ComparePackets(p, secondPacket) == 0 {
			indexes = append(indexes, i+1)
			if len(indexes) == 2 {
				break
			}
		}
	}

	return strconv.Itoa(indexes[0] * indexes[1])
}

// ComparePackets
//
// < 0: right order
//
//	0: continue
//	> 0: wrong order
func ComparePackets(left, right any) int {
	var leftValues, rightValues []any
	isLeftNum, isRightNum := false, false
	switch val := left.(type) {
	case float64:
		leftValues, isLeftNum = []any{val}, true
	case []any:
		leftValues = val
	}

	switch val := right.(type) {
	case float64:
		rightValues, isRightNum = []any{val}, true
	case []any:
		rightValues = val
	}

	if isLeftNum && isRightNum {
		return int(leftValues[0].(float64) - rightValues[0].(float64))
	}

	for i := 0; i < len(leftValues) && i < len(rightValues); i++ {
		if res := ComparePackets(leftValues[i], rightValues[i]); res != 0 {
			return res
		}
	}

	return len(leftValues) - len(rightValues)
}
