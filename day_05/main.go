package main

import (
	"fmt"
	fetch "getdata"
	"regexp"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2022/day/5/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

func parseCrate(data []string) [][]string {
	allCrates := make([][]string, 9)
	cratesSplit := strings.Split(data[0], "\n")
	for _, v := range cratesSplit {
		tmpLine := strings.Split(v, "")
		for x, r := range tmpLine {
			IsLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
			if IsLetter(r) {
				allCrates[x/4] = append([]string{r}, allCrates[x/4]...)
			}

		}
	}
	return allCrates
}

func parseMove(data []string) [][]int {
	var allMoves [][]int
	moveSplit := strings.Split(data[1], "\n")
	for _, v := range moveSplit {
		moves := strings.Split(v, " ")
		num, _ := strconv.Atoi(moves[1])
		source, _ := strconv.Atoi(moves[3])
		target, _ := strconv.Atoi(moves[5])
		allMoves = append(allMoves, []int{num, source, target})
	}
	return allMoves
}

var ansPartOne string
var ansPartTwo string

func moveCrate(crates [][]string, num int, source int, target int) [][]string {
	source -= 1
	target -= 1
	var tmp = []string{}
	for i := 0; i < num; i++ {
		removedItem := crates[source][len(crates[source])-1:]
		fmt.Printf("[INFO]: moving item %s to stack : %d\n", removedItem, target+1)
		tmp = append(tmp, removedItem...)
		crates[source][len(crates[source])-1:][0] = ""
		crates[source] = crates[source][:len(crates[source])-1]
	}
	crates[target] = append(crates[target], tmp...)
	return crates
}

func moveCrateTwo(crates [][]string, num int, source int, target int) [][]string {
	source -= 1
	target -= 1
	var tmp = []string{}
	for i := 0; i < num; i++ {
		removedItem := crates[source][len(crates[source])-1:]
		fmt.Printf("[INFO]: moving item %s to stack : %d\n", removedItem, target+1)
		tmp = append([]string{removedItem[0]}, tmp...)
		crates[source][len(crates[source])-1:][0] = ""
		crates[source] = crates[source][:len(crates[source])-1]
	}
	crates[target] = append(crates[target], tmp...)
	return crates
}

func getAns(crates [][]string) string {
	var ans string
	for _, v := range crates {
		ans += v[len(v)-1]
	}
	return ans
}

func main() {
	fmt.Println(input)
	inputClean := strings.Trim(input, "\n")
	data := strings.Split(inputClean, "\n\n")
	moveData := parseMove(data)
	cratesPartOne := parseCrate(data)
	cratesPartTwo := parseCrate(data)
	for _, v := range moveData {
		num := v[0]
		source := v[1]
		target := v[2]
		cratesPartOne = moveCrate(cratesPartOne, num, source, target)
		cratesPartTwo = moveCrateTwo(cratesPartTwo, num, source, target)
	}
	fmt.Println(cratesPartOne)
	fmt.Println(cratesPartTwo)
	ansPartOne = getAns(cratesPartOne)
	ansPartTwo = getAns(cratesPartTwo)
	fmt.Println("[INFO]: ans part one => : ", ansPartOne)
	fmt.Println("[INFO]: ans part two => : ", ansPartTwo)
}
