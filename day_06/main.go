package main

import (
	"fmt"
	"strings"

	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/6/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

func duplicateCheck(arr []string) int {
	visited := make(map[string]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] {
			return i
		} else {
			visited[arr[i]] = true
		}
		fmt.Println(visited)
	}
	return -1
}

var countPartOne, countPartTwo int = 0, 0

func main() {
	blocPartOne, blocPartTwo := input[:4], input[:14]
	blocSplit, blocSplitPartTwo := strings.Split(blocPartOne, ""), strings.Split(blocPartTwo, "")
	blocVal, blocValPartTwo := duplicateCheck(blocSplit), duplicateCheck(blocSplitPartTwo)
	for blocVal != -1 {
		countPartOne += 1
		blocPartOne = input[countPartOne : 4+countPartOne]
		blocSplit = strings.Split(blocPartOne, "")
		blocVal = duplicateCheck(blocSplit)
	}
	for blocValPartTwo != -1 {
		countPartTwo += 1
		blocPartTwo = input[countPartTwo : 14+countPartTwo]
		blocSplitPartTwo = strings.Split(blocPartTwo, "")
		blocValPartTwo = duplicateCheck(blocSplitPartTwo)
	}
	fmt.Println("[INFO]: ans part one => : ", len(blocSplit)+countPartOne)
	fmt.Println("[INFO]: ans part two => : ", len(blocSplitPartTwo)+countPartTwo)
}
