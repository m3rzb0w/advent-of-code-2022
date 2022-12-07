package main

import (
	"fmt"
	"strconv"
	"strings"

	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/4/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

var ansPartOne int = 0
var ansPartTwo int = 0

func main() {
	inputClean := strings.Trim(input, "\n")
	data := strings.Split(inputClean, "\n")
	for _, v := range data {
		splitPairs := strings.Split(v, ",")
		firstPairSplit := strings.Split(splitPairs[0], "-")
		secondPairSplit := strings.Split(splitPairs[1], "-")
		a, _ := strconv.Atoi(firstPairSplit[0])
		b, _ := strconv.Atoi(firstPairSplit[1])
		c, _ := strconv.Atoi(secondPairSplit[0])
		d, _ := strconv.Atoi(secondPairSplit[1])
		// a---b
		// c---d
		//part one
		if a <= c && b >= d || a >= c && b <= d {
			ansPartOne += 1
		}
		//part two
		if !(b < c || d < a) {
			ansPartTwo += 1
		}
	}
	fmt.Println("Part one =>", ansPartOne)
	fmt.Println("Part two =>", ansPartTwo)
}
