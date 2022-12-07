package main

import (
	"fmt"
	"strings"

	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/3/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

// partOne
func analyze(partOne string, partTwo string) []rune {
	container := []rune{}
	for _, v := range partOne {
		for _, c := range partTwo {
			if c == v {
				container = append(container, v)
			}
		}
	}
	return container
}

// partTwo
func analyzeTwo(tmpPartTwo []string) []rune {
	partOne := tmpPartTwo[0]
	partTwo := tmpPartTwo[1]
	partThree := tmpPartTwo[2]
	container := []rune{}
	for _, v := range partOne {
		for _, c := range partTwo {
			for _, x := range partThree {
				if c == v && v == x {
					container = append(container, v)
				}
			}

		}
	}
	return container
}

func runeToNumAlphabet(inputRune rune) int {
	intRune := int(inputRune)
	if intRune >= 97 {
		return intRune - 96
	} else {
		return intRune - 38
	}
}

//var input = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw\n"

func main() {
	inputClean := strings.Trim(input, "\n")
	data := strings.Split(inputClean, "\n")
	counterPartOne := 0
	counterPartTwo := 0
	loopCount := 0
	var tmpPartTwo []string
	for _, v := range data {
		firstComp := v[0 : len(v)/2]
		secondComp := v[len(v)/2:]
		//fmt.Println(firstComp, secondComp)
		//part Two
		if loopCount <= 3 {
			tmpPartTwo = append(tmpPartTwo, v)
			loopCount++
		}
		if loopCount == 3 {
			//fmt.Println(tmpPartTwo)
			sameValItemsPartTwo := analyzeTwo(tmpPartTwo)
			priorityValue := runeToNumAlphabet(sameValItemsPartTwo[0])
			counterPartTwo += priorityValue
			loopCount = 0
			tmpPartTwo = nil
		}
		//part One
		sameValItemsPartOne := analyze(firstComp, secondComp)
		priorityValue := runeToNumAlphabet(sameValItemsPartOne[0])
		counterPartOne += priorityValue
	}
	fmt.Println("Part one => ", counterPartOne)
	fmt.Println("Part two => ", counterPartTwo)
}
