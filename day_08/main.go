package main

import (
	"fmt"
	fetch "getdata"
	"sort"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2022/day/8/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

func strToInt(arr []string) []int {
	var res []int
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, num)
	}
	return res
}

func maxLeft(forest [][]int, row int, val int) int {
	var largerNumber int
	nums := forest[row][:val]
	// fmt.Println("MAXVAL LEFT", nums)
	for _, element := range nums {
		if element > largerNumber {
			largerNumber = element
		}
	}
	// fmt.Println(largerNumber)
	return largerNumber
}

func maxRight(forest [][]int, row int, val int) int {
	var largerNumber int
	nums := forest[row][val+1 : len(forest[row])]
	// fmt.Println(nums)
	// fmt.Println("LEN FOREST VAL3 ", len(forest[row]), val)
	// // fmt.Println("MAX RIGHT LIST", nums)
	for _, element := range nums {
		if element > largerNumber {
			largerNumber = element
		}
	}
	// fmt.Println(largerNumber)
	return largerNumber
}

func maxUp(forest [][]int, row int, val int) int {
	n := row
	var nums []int
	var largerNumber int
	for n >= 1 {
		// fmt.Println(n)
		tree := forest[row-n][val]
		nums = append(nums, tree)
		n--
	}
	// fmt.Println("MAXUP NUM", nums, val)
	for _, element := range nums {
		if element > largerNumber {
			largerNumber = element
		}
	}
	// fmt.Println(largerNumber)
	return largerNumber
}

func maxDown(forest [][]int, row int, val int) int {
	n := 1
	var nums []int
	var largerNumber int
	for n < len(forest)-row {
		// fmt.Println(n)
		tree := forest[row+n][val]
		nums = append(nums, tree)
		n++
	}
	// fmt.Println(nums)
	for _, element := range nums {
		if element > largerNumber {
			largerNumber = element
		}
	}
	// fmt.Println(largerNumber)
	return largerNumber
}

func maxViewLeft(forest [][]int, row int, val int) int {
	var count int
	tree := forest[row][val]
	nums := forest[row][:val]
	// fmt.Println("MAX VIEW LEFT", nums)
	for i := range nums {
		element := nums[len(nums)-1-i]
		if element < tree {
			count++
		}
		if element >= tree {
			count++
			break
		}
	}
	return count
}

func maxViewRight(forest [][]int, row int, val int) int {
	var count int
	tree := forest[row][val]
	nums := forest[row][val+1 : len(forest[row])]
	// fmt.Println(nums)
	// fmt.Println("LEN FOREST VAL3 ", len(forest[row]), val)
	// // fmt.Println("MAX RIGHT LIST", nums)
	for _, element := range nums {
		if element < tree {
			count++
		}
		if element >= tree {
			count++
			break
		}
	}
	return count
}

func maxViewUp(forest [][]int, row int, val int) int {
	n := row
	tree := forest[row][val]
	var nums []int
	var count int
	for n >= 1 {
		// fmt.Println(n)
		tree := forest[row-n][val]
		nums = append(nums, tree)
		n--
	}
	fmt.Println("NUMS UP", nums)
	for i := range nums {
		element := nums[len(nums)-1-i]
		if element < tree {
			count++
		}
		if element >= tree {
			count++
			break
		}
	}
	return count
}

func maxViewDown(forest [][]int, row int, val int) int {
	n := 1
	tree := forest[row][val]
	var nums []int
	var count int
	for n < len(forest)-row {
		// fmt.Println(n)
		tree := forest[row+n][val]
		nums = append(nums, tree)
		n++
	}
	fmt.Println("NUMS DOWN", nums)
	for _, element := range nums {
		if element < tree {
			count++
		}
		if element >= tree {
			count++
			break
		}
	}
	return count
}

var ans int
var ansTwo []int

func main() {
	fmt.Println("hello")
	inputTrim := strings.Trim(input, "\n")
	dataSplit := strings.Split(inputTrim, "\n")
	fmt.Println(len(dataSplit))
	forest := make([][]int, len(dataSplit))
	fmt.Println(dataSplit)
	for i, v := range dataSplit {
		fmt.Println(v)
		line := strings.Split(v, "")
		fmt.Println(line)
		lineInt := strToInt(line)
		forest[i] = append(forest[i], lineInt...)
	}
	fmt.Println("FOREST DONE =>", forest)
	fmt.Println("Len forest=", len(forest))
	//row
	for i := range forest {
		// fmt.Println("currentIndex =>", i, forest[i])
		for j := range forest[i] {
			// fmt.Println(len(forest[i]), j)
			currentNum := forest[i][j]
			numMaxLeft := maxLeft(forest, i, j)
			numMaxRight := maxRight(forest, i, j)
			numMaxUp := maxUp(forest, i, j)
			numMaxDown := maxDown(forest, i, j)

			//part2
			numMaxViewLeft := maxViewLeft(forest, i, j)
			// fmt.Printf("num max view left => %d, for %d\n", numMaxViewLeft, currentNum)
			numMaxViewRight := maxViewRight(forest, i, j)
			// fmt.Printf("num max view right => %d, for %d\n", numMaxViewRight, currentNum)
			numMaxViewUp := maxViewUp(forest, i, j)
			// fmt.Printf("num max view up => %d, for %d\n", numMaxViewUp, currentNum)
			numMaxViewDown := maxViewDown(forest, i, j)
			// fmt.Printf("num max view down => %d, for %d\n", numMaxViewDown, currentNum)
			scenicScore := numMaxViewLeft * numMaxViewRight * numMaxViewUp * numMaxViewDown
			fmt.Printf("SCENIC SCORE  => %d, for %d\n", scenicScore, currentNum)
			fmt.Printf("CURRENT NUM => %d, MAXVIEWLEFT: %d, MAXVIEWRIGHT : %d, MAXVIEWUP : %d, MAXVIEWDOWN : %d\n", currentNum, numMaxViewLeft, numMaxViewRight, numMaxViewUp, numMaxViewDown)
			ansTwo = append(ansTwo, scenicScore)
			// fmt.Println(ansTwo)

			fmt.Printf("CURRENT NUM => %d, MAXLEFT : %d, MAXRIGHT : %d, MAXUP : %d, MAXDOWN : %d\n", currentNum, numMaxLeft, numMaxRight, numMaxUp, numMaxDown)

			if j == 0 || numMaxLeft < currentNum {
				// fmt.Println("resp left=>", currentNum)
				ans += 1
			} else if j == len(forest[i])-1 || numMaxRight < currentNum {
				// fmt.Println("resp right =>", currentNum)
				ans += 1
			} else if i == 0 || numMaxUp < currentNum {
				// fmt.Println("resp up=>", currentNum)
				ans += 1
			} else if i == len(forest)-1 || numMaxDown < currentNum {
				// fmt.Println("resp down=>", currentNum)
				ans += 1
			}
			// fmt.Print(currentNum)
		}

	}
	fmt.Println("ANS => ", ans)
	sort.Ints(ansTwo)
	fmt.Println(ansTwo)
	fmt.Println(ansTwo[len(ansTwo)-1])

}
