package main

import (
	"fmt"
	fetch "getdata"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2022/day/9/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

type Position struct {
	x, y int
}

func (tail *Position) tailMoves(head Position) {

	p := Position{head.x - tail.x, head.y - tail.y}
	switch p {
	case Position{2, -2}, Position{1, -2}, Position{2, -1}:
		tail.x++
		tail.y--
	case Position{-1, -2}, Position{-2, -2}, Position{-2, -1}:
		tail.x--
		tail.y--
	case Position{-2, 2}, Position{-1, 2}, Position{-2, 1}:
		tail.x--
		tail.y++
	case Position{1, 2}, Position{2, 2}, Position{2, 1}:
		tail.x++
		tail.y++
	case Position{0, -2}:
		tail.y--
	case Position{-2, 0}:
		tail.x--
	case Position{2, 0}:
		tail.x++
	case Position{0, 2}:
		tail.y++
	}
}

var head, tail = Position{0, 0}, Position{0, 0}
var headMoveList, tailMoveList = []Position{}, []Position{}
var visitedByTail = make(map[Position]bool)

//partTwo

var knots = make([]Position, 10)
var visitedByKnotNine = make(map[Position]bool)

func main() {
	inputTrim := strings.Trim(input, "\n")
	dataSplit := strings.Split(inputTrim, "\n")
	for _, v := range dataSplit {
		line := strings.Split(v, " ")
		direction := line[0]
		moveNum, _ := strconv.Atoi(line[1])
		fmt.Println(direction, moveNum)
		for moveNum > 0 {
			fmt.Println(moveNum)
			switch direction {
			case "R":
				head.x++
				knots[0].x++
			case "U":
				head.y++
				knots[0].y++
			case "L":
				head.x--
				knots[0].x--
			case "D":
				head.y--
				knots[0].y--
			}

			for i := range knots[1:] {
				knots[i+1].tailMoves(knots[i])
				fmt.Println("KNOTS STATUS =>", knots)

			}

			moveNum--
			headMoveList = append(headMoveList, head)
			tail.tailMoves(head)
			tailMoveList = append(tailMoveList, tail)
			visitedByTail[tail] = true
			visitedByKnotNine[knots[9]] = true

		}
		fmt.Println("visited by tail => :", len(visitedByTail))
		fmt.Println("Move by Head => : ", len(headMoveList))
		fmt.Println("Move by Knot9 => : ", len(visitedByKnotNine), visitedByKnotNine)

	}

}
