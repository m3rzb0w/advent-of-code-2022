package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/1/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

var biggerNum int

var num []int

func main() {
	data := strings.Split(input, "\n")
	totalMaxCalories := 0
	for _, v := range data {
		vnum, _ := strconv.Atoi(v)
		if vnum == 0 {
			if biggerNum > totalMaxCalories {
				totalMaxCalories = biggerNum
			}
			num = append(num, biggerNum)
			biggerNum = vnum
		}
		biggerNum += vnum
	}
	num = append(num, biggerNum)
	fmt.Println("Part 1 max calories => : ", totalMaxCalories)
	sort.Ints(num)
	//fmt.Println(num)
	totalCal := num[len(num)-1] + num[len(num)-2] + num[len(num)-3]
	fmt.Println("Part 2 max calories top 3 => : ", totalCal)
}
