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
	mixingFile := NewMixingFile(data)
	mixingFile.Mix()
	res := 0
	for _, item := range mixingFile.GetGrooveItems() {
		res += item.value
	}
	return strconv.Itoa(res)
}

func part2(data string) string {
	mixingFile := NewMixingFile(data)
	for _, item := range mixingFile.numbers {
		item.value *= 811589153
	}
	for i := 0; i < 10; i++ {
		mixingFile.Mix()
	}
	res := 0
	for _, item := range mixingFile.GetGrooveItems() {
		res += item.value
	}
	return strconv.Itoa(res)
}

type Item struct {
	value         int
	originalIndex int
	currentIndex  int
}

type MixingFile struct {
	numbers   []*Item
	originals []*Item
	zero      *Item
	length    int
}

func NewMixingFile(data string) MixingFile {
	mixingFile := MixingFile{
		numbers:   []*Item{},
		originals: []*Item{},
	}
	for i, line := range strings.Split(data, "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		item := &Item{value: num, originalIndex: i, currentIndex: i}
		mixingFile.numbers = append(mixingFile.numbers, item)
		mixingFile.originals = append(mixingFile.numbers, item)
		if num == 0 {
			mixingFile.zero = item
		}
	}
	mixingFile.length = len(mixingFile.numbers)
	return mixingFile
}

func (mf *MixingFile) Mix() {
	for i := 0; i < mf.length; i++ {
		// mix item at original index i
		// let's see where the item is now
		var currentItem = mf.originals[i]

		moves := currentItem.value % (mf.length - 1)

		if moves == 0 {
			continue
		}

		newIndex := currentItem.currentIndex + moves
		if newIndex <= 0 {
			newIndex += mf.length - 1
		}
		if newIndex > mf.length-1 {
			newIndex -= mf.length - 1
		}
		// update indexes
		if newIndex > currentItem.currentIndex {
			// element moved to the right
			// all elements between currentIndex + 1 and newIndex included move to the left
			for j := currentItem.currentIndex + 1; j < newIndex+1; j++ {
				mf.numbers[j].currentIndex -= 1
				if mf.numbers[j].currentIndex < 0 {
					mf.numbers[j].currentIndex += mf.length
				}
			}
		} else {
			// element moved to the left
			for j := newIndex; j < currentItem.currentIndex; j++ {
				mf.numbers[j].currentIndex += 1
				if mf.numbers[j].currentIndex > mf.length-1 {
					mf.numbers[j].currentIndex -= mf.length
				}
			}
		}

		// update numbers
		newNumbers := []*Item{}
		if newIndex > currentItem.currentIndex {
			if currentItem.currentIndex > 0 {
				newNumbers = append(newNumbers, mf.numbers[0:currentItem.currentIndex]...)
			}
			newNumbers = append(newNumbers, mf.numbers[currentItem.currentIndex+1:newIndex+1]...)
			newNumbers = append(newNumbers, currentItem)
			newNumbers = append(newNumbers, mf.numbers[newIndex+1:mf.length]...)

		} else {
			if newIndex > 0 {
				newNumbers = append(newNumbers, mf.numbers[0:newIndex]...)
			}
			newNumbers = append(newNumbers, currentItem)
			newNumbers = append(newNumbers, mf.numbers[newIndex:currentItem.currentIndex]...)
			newNumbers = append(newNumbers, mf.numbers[currentItem.currentIndex+1:mf.length]...)
		}
		mf.numbers = newNumbers

		currentItem.currentIndex = newIndex
	}
}

func (mf *MixingFile) GetGrooveItems() []*Item {
	grooves := []*Item{}
	zeroIndex := mf.zero.currentIndex
	grooves = append(grooves, mf.numbers[(zeroIndex+1000)%mf.length])
	grooves = append(grooves, mf.numbers[(zeroIndex+2000)%mf.length])
	grooves = append(grooves, mf.numbers[(zeroIndex+3000)%mf.length])
	return grooves
}
