package main

import (
	"fmt"
	fetch "getdata"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2022/day/11/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

type Monkey struct {
	id          string
	items       []string
	operation   string
	test        int
	targetTrue  int
	targetFalse int
	insp        int
}

func readOp(op string, item string) int {
	itemNum, _ := strconv.Atoi(item)
	opVal := idRegx.FindString(op)
	opNum, _ := strconv.Atoi(opVal)
	if strings.Contains(op, "+") {
		fmt.Println("+", opNum, itemNum)
		return opNum + itemNum
	}
	if strings.Contains(op, "*") && !strings.Contains(op, "old * old") {
		fmt.Println("*", opNum, itemNum)
		return opNum * itemNum
	}
	fmt.Println("old old?", op, itemNum)
	return itemNum * itemNum
}

var monkeyList []Monkey
var idRegx = regexp.MustCompile(`\d+`)
var itemRegx = regexp.MustCompile(`\d{2}`)

func main() {
	fmt.Println("toto")
	monkey := Monkey{"0", []string{"54", "89", "94"}, "old * 7", 17, 5, 3, 0}
	fmt.Println(monkey)
	inputTrim := strings.Trim(input, "\n")
	dataSplit := strings.Split(inputTrim, "\n\n")
	for _, v := range dataSplit {
		monkeyRawData := strings.Split(v, "\n")
		monkeyid := idRegx.FindString(monkeyRawData[0])
		monkeyitems := itemRegx.FindAllString(monkeyRawData[01], -1)
		monkeyOperation := strings.Split(monkeyRawData[2], "=")[1]
		monkeyOperation = strings.Trim(monkeyOperation, "")
		monkeyDivisible := idRegx.FindString(monkeyRawData[3])
		monkeyDivisibleNum, _ := strconv.Atoi(monkeyDivisible)
		targetTrue := idRegx.FindString(monkeyRawData[4])
		targetTrueNum, _ := strconv.Atoi(targetTrue)
		targetFalse := idRegx.FindString(monkeyRawData[5])
		targetFalseNum, _ := strconv.Atoi(targetFalse)
		fmt.Println(monkeyid, monkeyitems, monkeyOperation, monkeyDivisible, targetTrue, targetFalse)
		monkeyList = append(monkeyList, Monkey{monkeyid, monkeyitems, monkeyOperation, monkeyDivisibleNum, targetTrueNum, targetFalseNum, 0})
	}
	fmt.Println(monkeyList)
	n := 0
	for n < 10000 { //part 1 = 20 , part 2 = 10000
		n++
		for idx, monkey := range monkeyList {
			fmt.Println("MONKEY=>", monkey)
			for i := range monkey.items {
				fmt.Println(monkey.items[i])
				tmpItem := readOp(monkey.operation, monkey.items[i])
				//tmpItem /= 3 //part 1 only
				tmpItem = tmpItem % 9699690 //part 2 only + to do => modulo explicit
				tmpVal := strconv.Itoa(tmpItem)
				if tmpItem%monkey.test == 0 {
					fmt.Println("TMPVAL TRUE and Target=>", tmpVal, monkey.targetTrue)
					monkeyList[monkey.targetTrue].items = append(monkeyList[monkey.targetTrue].items, tmpVal)
				} else {
					fmt.Println("TMPVAL FALSE and Target=>", tmpVal, monkey.targetFalse)
					monkeyList[monkey.targetFalse].items = append(monkeyList[monkey.targetFalse].items, tmpVal)
				}
				fmt.Println("MONKEY INPS=>", monkeyList[idx])
				monkeyList[idx].insp++

			}
			fmt.Println(monkeyList[idx])
			monkeyList[idx].items = monkey.items[:0]
		}
	}
	fmt.Println(monkeyList)
	ans := []int{}
	for _, v := range monkeyList {
		ans = append(ans, v.insp)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	fmt.Println("Part ans => ", ans[0]*ans[1])
}
