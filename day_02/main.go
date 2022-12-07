package main

import (
	"fmt"
	"strings"

	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/2/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

var rock string = "A"
var paper string = "B"
var scissors string = "C"

var meRock string = "X"
var mePaper string = "Y"
var meScissors string = "Z"

var loss int = 0
var draw int = 3
var win int = 6

func main() {
	data := strings.Split(input, "\n")
	count := 0
	countTwo := 0
	for _, v := range data {
		round := strings.Split(v, " ")
		//round[0] = op
		//round[1] = me
		//part1 & part2
		//oprock
		if round[0] == rock && round[1] == meRock {
			count += draw + 1
			countTwo += loss + 3
		}
		if round[0] == rock && round[1] == mePaper {
			count += win + 2
			countTwo += draw + 1
		}
		if round[0] == rock && round[1] == meScissors {
			count += loss + 3
			countTwo += win + 2
		}
		//oppaper
		if round[0] == paper && round[1] == meRock {
			count += loss + 1
			countTwo += loss + 1
		}
		if round[0] == paper && round[1] == mePaper {
			count += draw + 2
			countTwo += draw + 2
		}
		if round[0] == paper && round[1] == meScissors {
			count += win + 3
			countTwo += win + 3
		}
		//opscissors
		if round[0] == scissors && round[1] == meRock {
			count += win + 1
			countTwo += loss + 2
		}
		if round[0] == scissors && round[1] == mePaper {
			count += loss + 2
			countTwo += draw + 3
		}
		if round[0] == scissors && round[1] == meScissors {
			count += draw + 3
			countTwo += win + 1
		}
	}
	fmt.Println("Count part one => : ", count)
	fmt.Println("Count part two => : ", countTwo)
}
