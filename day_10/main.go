package main

import (
	"fmt"
	fetch "getdata"
	"math"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2022/day/10/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

var cycles int
var ans int
var targetCycles = map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}

// part 2
// the CRT draws a single pixel during each cycle.
func drawPxl(cycles int, x int) {
	row, column := cycles/40, cycles%40
	fmt.Println("X AND CYCLE", x, cycles)
	if math.Abs(float64(x-cycles%40)) <= 1 {
		crt[row][column] = "█"
	} else {
		crt[row][column] = " "
	}
}

var crt = [6][40]string{}

func main() {
	x := 1
	inputTrim := strings.Trim(input, "\n")
	dataSplit := strings.Split(inputTrim, "\n")
	for _, v := range dataSplit {
		line := strings.Split(v, " ")
		instruction := line[0]

		//part2-------------------
		drawPxl(cycles, x)
		//------------------------

		if instruction == "noop" {
			cycles += 1
			if targetCycles[cycles] {
				ans += cycles * x
			}
		}

		if len(line) == 2 {
			num, _ := strconv.Atoi(line[1])
			cycles += 1

			//part2-------------------
			drawPxl(cycles, x)
			//------------------------

			if targetCycles[cycles] {
				ans += cycles * x

			}
			cycles += 1
			if targetCycles[cycles] {
				ans += cycles * x

			}
			x += num
		}

	}
	fmt.Println(ans)
	fmt.Println("cycles =>", cycles)
	ansTwo := ""
	for _, v := range crt {
		ansTwo += strings.Join(v[:], "") + "\n"
	}
	fmt.Println(ansTwo)
}
